package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"nordkapp42/http/handler"
	"nordkapp42/loaders"
)

const Addr = ":8888"
const ShutdownTimeout = time.Duration(10) * time.Second

func WithDB(db *bun.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithDataLoader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "loaders", loaders.NewLoaders())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewAppHandler(db *bun.DB) *http.ServeMux {
	mux := http.NewServeMux()
	// mux.Handle("/", playground.Handler("GraphQL playground", "/api"))
	mux.Handle("/api", WithDB(db)(WithDataLoader(handler.NewGraphQLHandler())))
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
		log.Printf("connect to %s for GraphQL playground", Addr)
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
