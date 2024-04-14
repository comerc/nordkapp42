package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/rs/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"nordkapp42/graph"
	"nordkapp42/http/handler"
	"nordkapp42/pkg/jwt"
)

const Addr = ":8888"
const ShutdownTimeout = time.Duration(10) * time.Second

// TODO: Лучшей практикой при обработке контекстных ключей будет создание неэкспортируемого пользовательского типа: `type key string; const myCustomKey key = "key"; ctx := context.WithValue(context.Background(), myCustomKey, "val")`

func WithCORS(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	return c.Handler(next)
}

func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := jwt.TrimBearer(r.Header.Get("Authorization"))
		if accessToken == "" {
			log.Println(errors.New("the auth token is missing in the initialization payload"))
			next.ServeHTTP(w, r)
			return
		}
		payload, err := jwt.ParseAccessToken(accessToken)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "JWTPayload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
		ctx := context.WithValue(r.Context(), "loaders", graph.NewLoaders())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewAppHandler(db *bun.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", playground.AltairHandler("GraphQL playground", "/api"))
	var h http.Handler
	h = handler.NewGraphQLHandler()
	// h = WithAuth(h)
	h = WithDataLoader(h)
	h = WithDB(db)(h)
	// h = WithCORS(h)
	mux.Handle("/api", h)
	return mux
}

func main() {
	dsn := "postgres://postgres:postgrespassword@postgres:5432/postgres?sslmode=disable"
	// Connect to the database using Bun and the PostgreSQL driver.
	// case 1
	// sqldb, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatal("open db error", err)
	// }
	// case 2
	// sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// case 3
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	// config.PreferSimpleProtocol = true
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	sqldb := stdlib.OpenDB(*config)

	// Create a Bun DB instance using the PostgreSQL dialect.
	db := bun.NewDB(sqldb, pgdialect.New())

	// Optionally, add a logger to Bun for debugging purposes.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	defer db.Close()

	// *********
	// go listen(db)

	// fmt.Println(`Type a message and press enter.

	// This message should appear in any other chat instances connected to the same
	// database.

	// Type "exit" to quit.`)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	msg := scanner.Text()
	// 	if msg == "exit" {
	// 		os.Exit(0)
	// 	}

	// 	_, err = db.ExecContext(context.Background(), "select pg_notify('chat', ?)", msg)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "Error sending notification:", err)
	// 		os.Exit(1)
	// 	}
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error scanning from stdin:", err)
	// 	os.Exit(1)
	// }
	// *********

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
		log.Printf("app start on %s", Addr)
		// if err := server.ListenAndServe(
		if err := server.ListenAndServeTLS(
			"localhost.pem",
			"localhost-key.pem",
		); err != http.ErrServerClosed {
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

// func listen(db *bun.DB) {
// 	conn, err := db.Conn(context.Background())
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error connection:", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()
// 	// pgxConn
// 	var pgxConn *pgx.Conn
// 	err = conn.Raw(func(driverConn any) error {
// 		pgxConn = driverConn.(*stdlib.Conn).Conn()
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "pgxConn:", err)
// 		os.Exit(1)
// 	}

// 	go func() {
// 		for {
// 			time.Sleep(1 * time.Second)
// 			_, err = db.ExecContext(context.Background(), `select pg_notify('rooms:updated', ?)`, "123")
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, "Error db.ExecContext:", err)
// 				os.Exit(1)
// 			}
// 		}
// 	}()

// 	_, err = conn.ExecContext(context.Background(), `LISTEN "rooms:updated"`)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error listening to chat channel:", err)
// 		os.Exit(1)
// 	}

// 	for {
// 		notification, err := pgxConn.WaitForNotification(context.Background())
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "Error waiting for notification:", err)
// 			os.Exit(1)
// 		}

// 		fmt.Println("PID:", notification.PID, "Channel:", notification.Channel, "Payload:", notification.Payload)
// 	}
// }
