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
- [ ] ~~Redis~~ (1387)
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
- [ ] GitHub Actions (54)
- [ ] golangci-lint
- [ ] uber-go/automaxprocs
- [ ] uber-go/goleak - умеет показывать какие горутины не померли
- [ ] Reindexer https://habr.com/ru/articles/346884/
- [ ] Jitsu https://habr.com/ru/companies/jitsu/articles/523464/
- [ ] Centrifugo
- [ ] Livekit.io

## Для DevOPS

- [ ] ArgoCD + Helm - Наверное чтобы было понимание как вообще скейлится нагрузка в проде, чтобы не писал монолит которому нужно добавлять CPU только. ​Да не забей, научишься писать helm chart, и выучи что такое deployment/HPA и какие типы service есть. Напиши Helm chart для деплоя приложения в Kubernetes. Тебе нужно отдавать метрики в формате prometheus/тебе нужно иметь endpoint для health check/иметь логи в формате json.
- [ ] Kubernetes Operations - это Operator SDK for Platform Engineer.
- [ ] С помощью интеграции GitLab и Google Kubernetes Engine (GKE), вы можете установить GitLab Runners на GKE одним кликом и сразу начать запускать свои конвейеры CI2.

## Применить либы

- [ ] gherkingen - для BDD
- [ ] testcontainers-go
- [ ] bufbuild/buf
- [ ] golang-migrate/migrate
- [ ] pressly/goose
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
- [ ] gorilla/mux | stdlib mux 1.22
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
- Простейшая реализация PUB/SUB 1-1 & 1-N
- Members: и отправляют и читают 
- Rooms: приватные (1-1) и общие (1-N)
- Все сообщения хранятся вечно, и могут быть получены в отложенном режиме
- GraphQL Subscribe

### Stage 2

- Поднять нагрузочное тестирование

### Stage 3

- Поднять k8s

### Stage 4

- Поднять Kafka

### Stage 5

- Микросервисы для "бутылочных горлышек" (а не для красоты)

### Stage 6

- EventSourcing+CQRS

### Stage 7

- Temporal

***

- https://www.youtube.com/watch?v=tv8muwgj-Y4
- https://www.youtube.com/watch?v=UP4w70VXKt4
- https://github.com/nodkz/conf-talks
- https://github.com/acelot/graphql-articles
- https://altairgraphql.dev/
- https://habr.com/ru/articles/510448/
- https://github.com/jaydenseric/graphql-multipart-request-spec
- https://github.com/vektah/dataloaden
- https://github.com/99designs/gqlgen
 

