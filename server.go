package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"nordkapp42/graph"
	"os"
	"os/signal"
	"syscall"
	"time"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/lib/pq"
)

const Addr = ":8888"
const ShutdownTimeout = time.Duration(10) * time.Second

func ForDB(ctx context.Context) *bun.DB {
	return ctx.Value("db").(*bun.DB)
}

func WithDB(db *bun.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func NewAppHandler(db *bun.DB) *http.ServeMux {
	srv := gqlhandler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/api"))
	mux.Handle("/api", WithDB(db)(srv))
	return mux
}

func main() {
	// Connect to the database using Bun and the PostgreSQL driver.
	sqldb, err := sql.Open("postgres", "postgres://postgres:postgrespassword@postgres:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal("open db error", err)
	}

	// Create a Bun DB instance using the PostgreSQL dialect.
	db := bun.NewDB(sqldb, pgdialect.New())

	// Optionally, add a logger to Bun for debugging purposes.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	defer db.Close()

	// {
	// 	db := getDB(ctx)
	// 	var num int
	// 	err := db.QueryRowContext(context.Background(), "SELECT 1").Scan(&num)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	server := http.Server{
		Addr:    Addr,
		Handler: NewAppHandler(db),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("app stopped due error", err)
		}
		log.Println("app stopped gracefully")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt

	log.Println("app interruption signal received")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("app shutdown failed", err)
	}

	// if err := env.Close(); err != nil {
	// 	log.Fatal("app environment closing failed", zap.Error(err))
	// }

}

// func NewGraphQLHandler() *gqlhandler.Server {
// 	handler := gqlhandler.New(
// 		runtime.NewExecutableSchema(
// 			newSchemaConfig(env),
// 		),
// 	)
// 	// Transports
// 	handler.AddTransport(transport.Websocket{
// 		KeepAlivePingInterval: websocketKeepAlivePingInterval,
// 	})
// 	handler.AddTransport(transport.Options{})
// 	handler.AddTransport(transport.POST{})
// 	handler.AddTransport(transport.MultipartForm{
// 		MaxUploadSize: maxUploadSize,
// 		MaxMemory:     maxUploadSize / 10,
// 	})
// 	// Query cache
// 	handler.SetQueryCache(lru.New(queryCacheLRUSize))
// 	// Enabling introspection
// 	handler.Use(extension.Introspection{})
// 	// APQ
// 	handler.Use(extension.AutomaticPersistedQuery{Cache: lru.New(automaticPersistedQueryCacheLRUSize)})
// 	// Complexity
// 	handler.Use(extension.FixedComplexityLimit(complexityLimit))
// 	// Unhandled errors logger
// 	handler.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
// 		env.Logger.Error("unhandled error", zap.String("error", fmt.Sprintf("%v", err)))
// 		return gqlerror.Errorf("internal server error")
// 	})
// 	return handler
// }

// func newSchemaConfig(env *app.Env) runtime.Config {
// 	cfg := runtime.Config{Resolvers: resolver.NewResolver(env)}
// 	cfg.Directives.InputUnion = directive.NewInputUnionDirective()
// 	cfg.Directives.SortRankInput = directive.NewSortRankInputDirective()
// 	cfg.Complexity.ArticleQuery.Find = resolver.ArticleQueryFindComplexity
// 	cfg.Complexity.ArticleFindList.TotalCount = resolver.ArticleFindListTotalCountComplexity
// 	return cfg
// }
