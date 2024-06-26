package handler

//goland:noinspection SpellCheckingInspection
import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"nordkapp42/graph"
	"nordkapp42/graph/directive"
	"nordkapp42/pkg/jwt"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
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
	// handler.AddTransport(transport.Websocket{
	// 	KeepAlivePingInterval: websocketKeepAlivePingInterval,
	// 	Upgrader: websocket.Upgrader{
	// 		HandshakeTimeout: time.Minute,
	// 		CheckOrigin: func(r *http.Request) bool {
	// 			return true // TODO: Under Construction
	// 		},
	// 		EnableCompression: true,
	// 	},
	// 	InitFunc: webSocketInit,
	// })
	handler.AddTransport(transport.SSE{})
	handler.AddTransport(transport.Options{})
	// handler.AddTransport(transport.GET{})
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
	// cfg.Directives.User = directive.User
	cfg.Complexity.Query.Rooms = graph.QueryRoomsComplexity
	cfg.Complexity.Subscription.Rooms = graph.SubscriptionRoomsComplexity
	cfg.Complexity.Room.Messages = graph.RoomMessagesComplexity
	return cfg
}

func webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	ctx = context.WithValue(ctx, "isSubscription", true)
	accessToken := jwt.TrimBearer(initPayload.Authorization())
	if accessToken == "" {
		return ctx, &initPayload, errors.New("the auth token is missing in the initialization payload")
	}
	payload, err := jwt.ParseAccessToken(accessToken)
	if err != nil {
		return ctx, &initPayload, err
	}
	ctx = context.WithValue(ctx, "JWTPayload", payload)
	// TODO: вынести dbConn - можно ли использовать общий для нескольких корневых резолверов?
	go func() {
		<-ctx.Done()
		fmt.Println("close context") // TODO: for debug only
	}()
	return ctx, &initPayload, nil
}
