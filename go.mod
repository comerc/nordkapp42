module nordkapp42

go 1.21.7

replace github.com/uptrace/bun => ../bun

replace github.com/uptrace/bun/dbfixture => ../bun/dbfixture

replace github.com/uptrace/bun/extra/bundebug => ../bun/extra/bundebug

replace github.com/uptrace/bun/dialect/pgdialect => ../bun/dialect/pgdialect

replace github.com/uptrace/bun/driver/sqliteshim => ../bun/driver/sqliteshim

require (
	github.com/99designs/gqlgen v0.17.45
	github.com/lib/pq v1.10.9
	github.com/uptrace/bun v1.1.17
	github.com/uptrace/bun/dialect/pgdialect v1.1.17
	github.com/uptrace/bun/extra/bundebug v1.1.17
	github.com/vektah/gqlparser/v2 v2.5.11
	github.com/vikstrous/dataloadgen v0.0.6
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.2.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/urfave/cli/v2 v2.27.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
