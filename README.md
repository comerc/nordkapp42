# nordkapp42

![](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Nordkapp_znad_morza_Barentsa.jpg/1024px-Nordkapp_znad_morza_Barentsa.jpg)

Напоминания по задачам

## Цели

- Каждодневная практика ответов на вопросы, которые спрашивают на собесах (анти пет-проект)
- Скрытая цель - "From Zero To Hero", подробности потом

## Требования

- Реализовать функционал в рамках секретного проекта
- Event Modeling > BDD > Integration/Unit Tests (via gherkingen) > code for development via tests for external API of modules - `package module_test` > Unit Tests Coverage for internal functions in modules (via ChatGPT)
- standard + evrone/go-clean-template + SOLID + Dependency Injection
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

## Для DevOPS

- [ ] ArgoCD + Helm - Наверное чтобы было понимание как вообще скейлится нагрузка в проде, чтобы не писал монолит которому нужно добавлять CPU только. ​Да не забей, научишься писать helm chart, и выучи что такое deployment/HPA и какие типы service есть. Напиши Helm chart для деплоя приложения в Kubernetes. Тебе нужно отдавать метрики в формате prometheus/тебе нужно иметь endpoint для health check/иметь логи в формате json.
- [ ] Kubernetes Operations - это Operator SDK for Platform Engineer.

## Применить либы

- [ ] gherkingen - для BDD
- [ ] testcontainers-go
- [ ] bufbuild/buf
- [ ] golang-migrate/migrate
- [ ] flaggy | go-flags | pflag
- [ ] [Лучший regexp для Go](https://habr.com/ru/articles/756222/)
- [ ] цветные логи: https://github.com/GolangLessons/url-shortener/blob/c3987f66469a8d0769add18521adb9023520be95/internal/lib/logger/handlers/slogpretty/slogpretty.go
- [ ] vegeta, wrk - для стресс-тестов
- [ ] allegro/bigcache - когда нужен просто кеш (рекомендации лучших собаководов из Avito)
- [ ] go-playground/validator - правильный валидатор
- [ ] ilyakaznacheev/cleanenv - yaml & env в одном флаконе + godotenv для чтения .env
- [ ] jackc/pgx/v5/pgxpool / go-pg + pool - PG Pool
- [ ] Masterminds/squirrel - SQL Builder (by Avito)
- [ ] er := errgroup.Group{}; eg.SetLimit(limit) - ещё один примитив синхронизации (golang.org/x/sync/errgroup)
- [ ] [Compile-time Dependency Injection for Go](https://github.com/google/wire)
- [ ] [Fx is a dependency injection system for Go](https://github.com/uber-go/fx)
- [ ] https://github.com/uber-go/zap
- [ ] https://github.com/golangci/golangci-lint
- [ ] https://github.com/uber-go/config
- [ ] [Методы организации DI и жизненного цикла приложения в GO](https://habr.com/ru/companies/vivid_money/articles/531822/)
- [ ] github.com/yonahd/kor@latest - инструмент для обнаружения неиспользуемых ресурсов Kubernetes
- [ ] gorilla/mux
