package handler

//goland:noinspection SpellCheckingInspection
import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"nordkapp42/graph"
	"nordkapp42/graph/directive"
	"nordkapp42/pkg/jwt"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const websocketKeepAlivePingInterval = 5 * time.Second
const maxUploadSize = 30 * 1024 * 1024
const queryCacheLRUSize = 1000
const automaticPersistedQueryCacheLRUSize = 100
const complexityLimit = 2000

func NewGraphQLHandler() *gqlhandler.Server {
	handler := gqlhandler.New(
		graph.NewExecutableSchema(
			newSchemaConfig(),
		),
	)

	// Transports
	handler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: websocketKeepAlivePingInterval,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				fmt.Println("********")
				return true // TODO: under construction
			},
		},
		InitFunc: webSocketInit,
	})
	handler.AddTransport(transport.Options{})
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.MultipartForm{
		MaxUploadSize: maxUploadSize,
		MaxMemory:     maxUploadSize / 10,
	})

	// Query cache
	handler.SetQueryCache(lru.New(queryCacheLRUSize))

	// Enabling introspection
	handler.Use(extension.Introspection{})

	// APQ
	handler.Use(extension.AutomaticPersistedQuery{Cache: lru.New(automaticPersistedQueryCacheLRUSize)})

	// Complexity
	handler.Use(extension.FixedComplexityLimit(complexityLimit))

	// Unhandled errors logger
	handler.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		log.Println(err)
		return gqlerror.Errorf("internal server error")
	})

	return handler
}

func newSchemaConfig() graph.Config {
	cfg := graph.Config{Resolvers: &graph.Resolver{}}
	cfg.Directives.Auth = directive.Auth
	cfg.Complexity.Query.Rooms = graph.QueryRoomsComplexity
	cfg.Complexity.Subscription.Rooms = graph.SubscriptionRoomsComplexity
	cfg.Complexity.Room.Messages = graph.RoomMessagesComplexity
	return cfg
}

func webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	raw := jwt.TrimBearer(initPayload.Authorization())
	if raw == "" {
		return ctx, &initPayload, errors.New("the auth token is missing in the initialization payload")
	}
	memberID, err := jwt.ValidateJWT(raw)
	if err != nil {
		return ctx, &initPayload, err
	}
	ctx = context.WithValue(ctx, "memberID", memberID)
	// TODO: https://github.com/99designs/gqlgen/issues/2474#issuecomment-1986030908
	// Add the token expiration as a deadline and append a close reason to the context values that will be send to the client before the websocket actually closes
	// (Also throw away the cancel function, which the linter does not like)
	// newCtx, _ := context.WithDeadline(transport.AppendCloseReason(ctx, "authentication token has expired"), time.Unix(token.ExpiresAt, 0))
	return ctx, &initPayload, nil
}
