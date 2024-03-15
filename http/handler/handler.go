package handler

//goland:noinspection SpellCheckingInspection
import (
	"context"
	"log"
	"nordkapp42/graph"
	"time"

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
	handler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: websocketKeepAlivePingInterval,
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
	cfg.Complexity.Query.Rooms = graph.QueryRoomsComplexity
	cfg.Complexity.Subscription.Rooms = graph.SubscriptionRoomsComplexity
	cfg.Complexity.Room.Messages = graph.RoomMessagesComplexity
	return cfg
}
