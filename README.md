# nordkapp42

![](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Nordkapp_znad_morza_Barentsa.jpg/1024px-Nordkapp_znad_morza_Barentsa.jpg)

Напоминания по задачам для клона Trello

## Цели

- Каждодневная практика ответов на вопросы, которые спрашивают на собесах (анти пет-проект)
- Скрытая цель - "From Zero To Hero", подробности потом

## Требования

- Реализовать функционал в рамках секретного проекта
- Event Modeling > BDD > Integration/Unit Tests (via gherkingen / goconvey) > code for development via tests for external API of modules - `package module_test` / blackbox for refactoring > Unit Tests Coverage for internal functions in modules `package module` / whitebox for modifications (via ChatGPT)
- standard + [modules/layout](https://go.dev/doc/modules/layout) + evrone/go-clean-template + SOLID + Dependency Injection
- API на gRPC внутри, ванильный GraphQL наружу

## Применить инфраструктуру (вхождений на hh.ru)

- [ ] ~~RabbitMQ~~ (1417)
- [ ] Redis (1387)
- [ ] Kafka (2573)
- [ ] Grafana (1491)
- [ ] ~~Nagios~~ (65)
- [ ] ~~Zabbix~~ (1289)
- [ ] ~~ELK~~ (826)
- [ ] ~~FileBeat~~ (22)
- [ ] Prometheus (1138)
- [ ] ClickHouse (1149)
- [ ] Kubernetes (2446) или Minikube
- [ ] ~~MongoDB~~ (780)
- [ ] ~~MySQL~~ (1661)
- [ ] ~~Greenplum~~ (746)
- [ ] PostgreSQL (4585)
- [ ] ~~Airflow~~ (676)
- [ ] ~~Cadence~~ (69)
- [ ] Temporal (8)
- [ ] ~~Nginx~~ (2027)
- [ ] ~~Haproxy~~ (232)
- [ ] ~~Traefik~~ (24)
- [ ] ~~KrakenD~~ (5)
- [ ] Hasura (3), только как API-Gateway
- [ ] ~~GitLab~~ (2157) для CI/CD
- [ ] ~~Jenkins~~ (1276)
- [ ] ~~TeamCity~~ (345)
- [ ] ~~DeployHQ~~ (132)
- [ ] GitHub Actions (54): [Try & Makefile](https://github.com/nektos/act), [Marketplace](https://github.com/marketplace/), [Starter-Workflows](https://github.com/actions/starter-workflows), [QuickStart](https://docs.github.com/en/actions/quickstart).
- [ ] golangci-lint
- [ ] uber-go/automaxprocs
- [ ] uber-go/goleak - умеет показывать какие горутины не померли
- [ ] Reindexer https://habr.com/ru/articles/346884/
- [ ] Jitsu https://habr.com/ru/companies/jitsu/articles/523464/
- [ ] Centrifugo + [gobwas/ws](https://github.com/centrifugal/centrifuge/blob/master/_examples/ws_benchmarks/benchmark_gobwas/main.go)
- [ ] Livekit.io
- [ ] MINIO https://github.com/minio/minio-go
- [ ] Stitching or Federation - https://habr.com/ru/articles/728476/

## Для DevOPS

- [ ] ArgoCD + Helm - Наверное чтобы было понимание как вообще скейлится нагрузка в проде, чтобы не писал монолит которому нужно добавлять CPU только. ​Да не забей, научишься писать helm chart, и выучи что такое deployment/HPA и какие типы service есть. Напиши Helm chart для деплоя приложения в Kubernetes. Тебе нужно отдавать метрики в формате prometheus/тебе нужно иметь endpoint для health check/иметь логи в формате json.
- [ ] Kubernetes Operations - это Operator SDK for Platform Engineer.
- [ ] С помощью интеграции GitLab и Google Kubernetes Engine (GKE), вы можете установить GitLab Runners на GKE одним кликом и сразу начать запускать свои конвейеры CI2.

## Применить либы

- [ ] [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) - для мапинга данных GraphQL - [example 1](https://github.com/99designs/gqlgen/blob/master/_examples/todo/todo.go) - [example 2](https://gqlgen.com/reference/changesets/)
- [ ] gherkingen - для BDD
- [ ] testcontainers-go
- [ ] bufbuild/buf
- [ ] golang-migrate/migrate
- [ ] sqlize 
  - Generate sql migration from diff between existed sql and objects
  - Generate migration version - compatible with golang-migrate/migrate
- [ ] pressly/goose
- [ ] [механизм миграций из Hasura](https://hasura.io/docs/latest/migrations-metadata-seeds/manage-migrations/)
- [ ] flaggy | go-flags | pflag
- [ ] [Лучший regexp для Go](https://habr.com/ru/articles/756222/)
- [ ] цветные логи: https://github.com/GolangLessons/url-shortener/blob/c3987f66469a8d0769add18521adb9023520be95/internal/lib/logger/handlers/slogpretty/slogpretty.go
- [ ] codesenberg/bombardier, tsenart/vegeta, grafana/k6, wrk - для стресс-тестов
- [ ] allegro/bigcache - когда нужен просто кеш (рекомендации лучших собаководов из Avito)
- [ ] go-playground/validator - правильный валидатор
- [ ] ilyakaznacheev/cleanenv - yaml & env в одном флаконе + godotenv для чтения .env
- [ ] Netflix/go-env
- [ ] caarlos0/env
- [ ] spf13/viper
- [ ] jackc/pgx/v5/pgxpool / go-pg + pool - PG Pool
- [ ] Masterminds/squirrel - SQL Builder (by Avito)
- [ ] er := errgroup.Group{}; eg.SetLimit(limit) - ещё один примитив синхронизации (golang.org/x/sync/errgroup)
- [ ] [Compile-time Dependency Injection for Go](https://github.com/google/wire)
- [ ] [Fx is a dependency injection system for Go](https://github.com/uber-go/fx)
- [ ] github.com/uber-go/zap
- [ ] github.com/golangci/golangci-lint
- [ ] github.com/uber-go/config
- [ ] [Методы организации DI и жизненного цикла приложения в GO](https://habr.com/ru/companies/vivid_money/articles/531822/)
- [ ] github.com/yonahd/kor@latest - инструмент для обнаружения неиспользуемых ресурсов Kubernetes
- [ ] gorilla/mux | stdlib mux 1.22 (or fiber or atreugo)
- [ ] bytedance/sonic
- [ ] failsafe-go.dev
- [ ] goconvey - is awesome BDD in Go
- [ ] github.com/uber-go/mock
- [ ] github.com/segmentio/kafka-go [Kafka, go и параллельные очереди](https://habr.com/ru/articles/769950/)
- [ ] https://github.com/grpc-ecosystem/grpc-gateway
- [ ] matoous/go-nanoid
- [ ] https://github.com/bufbuild/buf (вместо protoc для gRPC)
- [ ] https://github.com/mailhog/MailHog для тестирования почты
- [ ] https://github.com/IBM/sarama
- [ ] [настройки CORS](https://habr.com/ru/articles/598359/)
- [ ] [Boilerplate for Go REST API with PostgreSQL database](https://github.com/comerc/go-api-boilerplate)
- [ ] [PostGraphile - Instant lightning-fast GraphQL API backed primarily by your PostgreSQL database.](https://github.com/graphile/crystal/tree/main/postgraphile/postgraphile)
- [ ] [dataloader на дженериках](https://github.com/vikstrous/dataloadgen) - [example](https://github.com/stasundr/gqlgen-pg-todo-example)
- [ ] [bun - SQL-first Golang ORM](https://github.com/uptrace/bun)
- [ ] https://github.com/cckwes/fiber-graphql-go
- [ ] https://github.com/stretchr/testify

## Реализация

[Event Modeling](https://draft.io/a77sr5g3fhhmq7dyykmu5pzhr7yzvdrrrt5nf3gmsmaw)

## How To Start

```bash
$ brew install golangci-lint
$ brew install mockery
$ brew install go-task
```

## How to build & run

```bash
$ docker-compose up -d --build
```

## Верхнеуровневый план

### Stage 1

- [x] Составить план - уже хороший план
- [x] LiveSharing
- Boilerplate (модульный монолит)
- Все сообщения хранятся вечно, и могут быть получены в отложенном режиме
- Простейшая реализация PUB/SUB 1-1 & 1-N
- Members: и отправляют и читают
- Rooms: 1-1 и 1-N (приватные и публичные)
- GraphQL Subscribe via Websocket & SSE (https://github.com/99designs/gqlgen/pull/2498)

### Stage 2

- Поднять нагрузочное тестирование

### Stage 3

- Поднять Redis, а не RabbitMQ (для обработки всплесков нагрузки)

### Stage 4

- Поднять k8s

### Stage 5

- Микросервисы для "бутылочных горлышек" (а не для красоты)
- Поднять GRPS (для синхронных вызовов между микросервисами)
- Поднять Redis, а не RabbitMQ (для асинхронной шины данных между микросервисами)
- Поднять API Gateway https://loft.sh/blog/nginx-vs-traefik-vs-haproxy-comparing-kubernetes-ingress-controllers/
- Как поженить с Hasura под Highload

### Stage 6

- Чистую архитектуру, пожалуйста: ​или луковую, или гексагональную
- EventSourcing+CQRS
- Читаем с реплики, пишем в мастер

### Stage 7

- Temporal + https://github.com/vikstrous/tempts
- Поднять Kafka (для реализации "машины времени")

### Stage 8

- MINIO as S3 (для картинок)

### Stage 9

- Push-уведомления?

## gqlgen

- [Представляем gqlgen: генератор серверов GraphQL для Go](https://habr.com/ru/articles/792700/)
- [GraphQL и GO: простой способ бросить REST](https://www.youtube.com/watch?v=tv8muwgj-Y4)
- [gqlgen: GraphQL для строгих](https://www.youtube.com/watch?v=UP4w70VXKt4)
- https://github.com/comerc/graphql-articles
- https://github.com/nodkz/conf-talks
- https://altairgraphql.dev/
- [Лексоранги — что это такое и как их использовать для эффективной сортировки списков](https://habr.com/ru/articles/510448/)
- https://github.com/jaydenseric/graphql-multipart-request-spec
- https://github.com/vektah/dataloaden
- https://github.com/99designs/gqlgen
- [React Apollo, Gqlgen – авторизация. Часть 1](https://habr.com/ru/articles/598359/)
- [React Apollo, Gqlgen – авторизация. Часть 2](https://habr.com/ru/articles/599723/)
- [GraphQL и Golang](https://habr.com/ru/companies/ruvds/articles/444346/)
- [Разработка REST-серверов на Go. Часть 7: GraphQL](https://habr.com/ru/companies/ruvds/articles/568472/)
- [Как использовать GraphQL Federation для инкрементальной миграции с монолита (Python) на микросервисы (Go)](https://habr.com/ru/companies/macloud/articles/559352/)
- [Развитие инфраструктуры Погоды. Доклад Яндекса (как мы эволюционировали из REST в GraphQL и почему это сделало жизнь разработчиков фронтенда и приложений мягкой и шелковистой)](https://habr.com/ru/companies/yandex/articles/565234/)
- GQLGen PG TODO Example: [Variant1](https://github.com/comerc/gqlgen-pg-todo), [Variant2](https://github.com/stasundr/gqlgen-pg-todo-example)

---

- [GraphQL Multipart Request](https://www.apollographql.com/docs/router/executing-operations/subscription-multipart-protocol/)
- [Multipart HTTP protocol for GraphQL subscriptions](https://github.com/jaydenseric/graphql-multipart-request-spec)
- [GraphQL over WebSocket Protocol](https://github.com/apollographql/subscriptions-transport-ws/blob/master/PROTOCOL.md)
- [GraphQL Cursor Connections Specification](https://relay.dev/graphql/connections.htm)
- [Apollo Federation](https://www.apollographql.com/docs/apollo-server/federation/introduction)

---

Горизонтальное масштабирование кода на GoLang с использованием Kubernetes (k8s) можно реализовать следующим образом:

1. **Создание приложения**: Сначала создайте базовое веб-приложение на Go³.

2. **Контейнеризация приложения**: Примените Docker к вашему приложению Go, создав Dockerfile и собрав образ Docker³.

3. **Развертывание образа Docker в Kubernetes**: Создайте развертывание Kubernetes, которое использует ваш образ Docker³.

4. **Создание службы Kubernetes**: Создайте службу Kubernetes, которая обеспечивает доступ к вашему приложению³.

5. **Автомасштабирование**: Используйте Horizontal Pod Autoscaler (HPA) в Kubernetes для автоматического масштабирования вашего приложения в зависимости от нагрузки¹. HPA автоматически увеличивает или уменьшает количество подов в вашем развертывании в зависимости от текущей нагрузки¹.

6. **Оптимизация нагрузки на сервер**: Реализуйте лучшие практики для снижения нагрузки на сервер и достижения оптимального использования ресурсов². Это может включать в себя мониторинг с помощью готовности и живучести зондов, использование автомасштабирования и оптимизацию использования ресурсов².

7. **Тестирование**: Проведите тестирование, чтобы убедиться, что ваше приложение корректно масштабируется¹.

Помните, что горизонтальное масштабирование - это сложный процесс, который требует тщательного планирования и тестирования. Убедитесь, что вы полностью понимаете все аспекты вашего приложения, прежде чем начинать масштабирование¹.

Source: Conversation with Bing, 2/11/2024
(1) Build and Deploy Containerized Applications with Golang on Kubernetes. https://www.coding-bootcamps.com/build-containerized-applications-with-golang-on-kubernetes/.
(2) How to Configure Kubernetes Horizontal Pod Autoscaler using Metrics .... https://www.digitalocean.com/community/tutorials/how-to-configure-kubernetes-horizontal-pod-autoscaler-using-metrics-server.
(3) Optimizing Golang Applications for Kubernetes: Best Practices ... - Earthly. https://earthly.dev/blog/optimize-golang-for-kubernetes/.
(4) scale package - k8s.io/client-go/scale - Go Packages. https://pkg.go.dev/k8s.io/client-go/scale.
(5) Create a Kubernetes Operator in Golang to automatically manage a simple .... https://developers.redhat.com/blog/2020/12/16/create-a-kubernetes-operator-in-golang-to-automatically-manage-a-simple-stateful-application.

---

Выбор между Kafka, NATS и RabbitMQ для обработки всплесков нагрузки также зависит от ваших конкретных потребностей и предпочтений. Вот некоторые ключевые особенности каждого из них:

**Kafka**:

- Kafka обеспечивает высокую пропускную способность для обработки входящих и исходящих потоков данных в реальном времени¹.
- Kafka обеспечивает устойчивость к сбоям благодаря репликации данных между узлами¹.
- Kafka поддерживает "по крайней мере один раз" и "ровно один раз" доставку сообщений¹.
- Kafka не поддерживает многопользовательский режим¹.

**NATS**:

- NATS известен своей высокой масштабируемостью и низкой задержкой. Он спроектирован для эффективной обработки большого количества подписчиков и сообщений¹.
- NATS поддерживает "по крайней мере один раз", и "ровно один раз" доставку сообщений в JetStream¹.
- NATS поддерживает истинный многопользовательский режим и децентрализованную безопасность через учетные записи и определение общих потоков и услуг¹.

**RabbitMQ**:

- RabbitMQ обеспечивает "по крайней мере один раз" и "ровно один раз" доставку сообщений¹.
- Масштабируемость RabbitMQ ограничена его дизайном. Он полагается на централизованный брокер, который может стать узким местом при обработке большого объема сообщений².
- RabbitMQ поддерживает многопользовательский режим с помощью виртуальных хостов; обмен данными не поддерживается¹.

Все три инструмента являются мощными и могут быть использованы для обработки всплесков нагрузки, но выбор между ними будет зависеть от ваших конкретных требований и предпочтений. Вы можете найти дополнительную информацию о сравнении Kafka, NATS и RabbitMQ в [этой статье](^1^) и [этой статье](^2^).

Source: Conversation with Bing, 2/11/2024
(1) Compare NATS - NATS Docs. https://docs.nats.io/compare-nats.
(2) RabbitMQ против Kafka: два разных подхода к обмену сообщениями. https://habr.com/ru/companies/itsumma/articles/416629/.
(3) Compare NATS - NATS Docs. https://docs.nats.io/compare-nats.
(4) RabbitMQ против Kafka: два разных подхода к обмену сообщениями. https://habr.com/ru/companies/itsumma/articles/416629/.
(5) Apache Kafka и RabbitMQ: в чем разница и что лучше изучать?. https://habr.com/ru/companies/slurm/articles/666326/.

---

Tarantool может быть использован в качестве горячего кэша перед PostgreSQL. Вот некоторые основные преимущества этого подхода:

1. **Высокая производительность**: Tarantool известен своей высокой производительностью и способностью обрабатывать большое количество запросов в секунду¹.

2. **Гибкость**: Tarantool позволяет вам хранить данные в памяти и на диске, что дает вам гибкость в управлении вашими данными¹.

3. **Простота использования**: Tarantool имеет простой и интуитивно понятный API, который делает его легким в использовании¹.

4. **Совместимость с PostgreSQL**: Существует коннектор Tarantool для PostgreSQL, который позволяет вам легко интегрировать эти две системы³.

Однако, стоит отметить, что использование Tarantool в качестве кэша перед PostgreSQL может потребовать дополнительной настройки и адаптации вашего приложения. Вам нужно будет настроить ваше приложение так, чтобы оно сначала проверяло данные в Tarantool, а затем, если данные не найдены, обращалось к PostgreSQL¹.

Вы можете найти дополнительную информацию о использовании Tarantool в качестве кэша в [этой статье](^1^) и [этой статье](^2^).

Source: Conversation with Bing, 2/11/2024
(1) Как использовать tarantool в качестве кэша существующей базы?. https://ru.stackoverflow.com/questions/700552/%d0%9a%d0%b0%d0%ba-%d0%b8%d1%81%d0%bf%d0%be%d0%bb%d1%8c%d0%b7%d0%be%d0%b2%d0%b0%d1%82%d1%8c-tarantool-%d0%b2-%d0%ba%d0%b0%d1%87%d0%b5%d1%81%d1%82%d0%b2%d0%b5-%d0%ba%d1%8d%d1%88%d0%b0-%d1%81%d1%83%d1%89%d0%b5%d1%81%d1%82%d0%b2%d1%83%d1%8e%d1%89%d0%b5%d0%b9-%d0%b1%d0%b0%d0%b7%d1%8b.
(2) GitHub - tarantool/pg: PostgreSQL connector for Tarantool. https://github.com/tarantool/pg.
(3) Как использовать tarantool в качестве кэша существующей базы?. https://ru.stackoverflow.com/questions/700552/%d0%9a%d0%b0%d0%ba-%d0%b8%d1%81%d0%bf%d0%be%d0%bb%d1%8c%d0%b7%d0%be%d0%b2%d0%b0%d1%82%d1%8c-tarantool-%d0%b2-%d0%ba%d0%b0%d1%87%d0%b5%d1%81%d1%82%d0%b2%d0%b5-%d0%ba%d1%8d%d1%88%d0%b0-%d1%81%d1%83%d1%89%d0%b5%d1%81%d1%82%d0%b2%d1%83%d1%8e%d1%89%d0%b5%d0%b9-%d0%b1%d0%b0%d0%b7%d1%8b.
(4) postgresql - Using Tarantool and Postgres together - Stack Overflow. https://stackoverflow.com/questions/55460604/using-tarantool-and-postgres-together.
(5) undefined. https://www.computerweekly.com/feature/Write-through-write-around-write-back-Cache-explained.
(6) undefined. https://github.com/bucardo/bucardo.
(7) undefined. https://www.tarantool.io/en/doc/2.2/reference/reference_rock/dbms/.

## Что нужно от брокера?

![](./assets/pub-sub-broker.png)

- Производительность
- Сохранение порядка сообщений - [ULID](https://danyloff.tech/2023/10/01/ulid-%D1%82%D0%BE%D1%82-%D1%81%D0%B0%D0%BC%D1%8B%D0%B8-%D1%84%D0%BE%D1%80%D0%BC%D0%B0%D1%82-%D0%B8%D0%B4%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%82%D0%BE%D1%80%D0%B0-%D0%BA%D0%BE%D1%82/)
- Масштабируемость брокера
- Миллионы топиков
- Кэш/стрим сообщений
- Возможность писать процедуры - большой бонус

[Масштабируем WebSocket-соединения на Go / Александр Емелин (Авито)](https://www.youtube.com/watch?v=eP_02zkwM5Y)

> на 100000 соединений RabbitMQ потреблял 70 CPU, а Redis - 0,3 CPU

- Sentinel является менеджером отказоустойчивости репликации для Redis.

- github.com/gogo/protobuf - ускоряет сериализацию protobuf в пять раз, т.к. вметсо reflect использует кодогенерацию (ffjson & easyjson - такая же история).

***

Go-Fiber

https://github.com/99designs/gqlgen/discussions/1802
https://github.com/99designs/gqlgen/issues/1281
https://github.com/99designs/gqlgen/issues/1664#issuecomment-1616620967

***

Нагрузочное тестирование Centrifugo

https://github.com/centrifugal/centrifugo/blob/master/misc/benchmarking/k6/readme.md
https://github.com/centrifugal/centrifuge/tree/master/_examples/ws_benchmarks/benchmark_gobwas
https://github.com/centrifugal/centrifuge/tree/master/_examples/ws_benchmarks/benchmark_gorilla

***

## Setup Hasura

```bash
$ cd hasura && docker-compose up -d
```

#### Connect Database

- Database Display Name: default
- Environment Variable: PG_DATABASE_URL

## How to save DB-Schema

```bash
$ cd hasura
$ rm -rf migrations
$ hasura migrate create "init" --from-server --database-name default
$ rm -rf metadata
$ hasura metadata export
```

## How to restore DB-Schema

```bash
$ cd hasura
$ hasura migrate apply
$ hasura metadata apply
```

## How to backup data

```bash
$ cd hasura
$ curl --location --request POST 'http://localhost:8080/v1alpha1/pg_dump' --header 'x-hasura-admin-secret: myadminsecretkey' --header 'Content-Type: application/json' --data-raw '{ "opts": ["-O", "-x", "--data-only", "--schema", "public", "--schema", "auth"], "clean_output": true}' -o data.sql
```

## How to restore data

```bash
$ cd hasura
$ cat data.sql | docker exec -i nordkapp42-postgres-1 psql -U postgres
```

## How To Start

```bash
$ npm install
$ npm run dev
```

## Install the Hasura CLI, and dive into the console

```bash
hasura init
cd hasura && hasura console
```

## How To Generate GraphQL Schema

via Apollo CLI (deprecated)

```bash
$ npm install -g apollo graphql
$ apollo client:download-schema --endpoint http://localhost:8080/v1/graphql --header "X-Hasura-Admin-Secret: myadminsecretkey"
```

via Apollo Rover

```bash
$ npm install -g @apollo/rover
$ rover graph introspect --header "X-Hasura-Admin-Secret: myadminsecretkey" http://localhost:8080/v1/graphql > schema.graphql
```

***

## Соглашения

- Таблицы именуются в единственном числе, для соответствия GraphQL-типов.
- Таблицы и колонки именуются через подчёркивание (в Python-стиле).
- Data Manager > Edit > GraphQL Customization > Naming Conversion > graphql-default
- Поля id, created_by, updated_by в начале каждой таблицы сущности.
- Enum-таблицы имеют одно поле с именем value.
- Поле id инкрементится через identity.
