# go-msvc

(opinionated + in-progress) Microservice Template in Golang

## Structure

```text
.
├── build
├── cmd
│   └── server
├── internal
│   ├── config
│   ├── database
│   ├── handler
│   └── model
└── migration
```

| Directory              | Description                                                                               |
|------------------------|-------------------------------------------------------------------------------------------|
| build                  | docker compose files for unit testing, Dockerfile for container image distribution        |
| cmd                    | application entrypoint, there can be multiple applications in this directory              |
| internal               | internal package                                                                          |
| internal/config        | central configuration (environment variables, files, etc)                                 |
| internal/database      | database connection management, records access, etc                                       |
| internal/handler       | restful api implementation                                                                |
| internal/model         | modelling for database                                                                    |
| migration              | sql files for migration up/down                                                           |

## Development

- [ ] Lint: golangci-lint.
- [ ] Code Analysis: Sonarqube.
- [ ] Continuous integration using Github Actions.
- [ ] Database schema migration using golang-migrate.
- [ ] Implement observability using OpenTelemetry.

## Testing

- [x] Unit test for pure database functions.
- [ ] Unit test for pure api functions.
- [ ] Integration between api and database.
- [ ] Integration test among µservices.

## Deployment

- [ ] Deploy all µservices using Skaffold + Minikube for local dev environment.
- [ ] Auto deployment (CD) to GKE using Skaffold + ArgoCD.
