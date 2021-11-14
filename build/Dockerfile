FROM golang:1.17.3 AS build

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src/

RUN go build -a -ldflags="-w -s" -installsuffix cgo -o /app/server cmd/server/main.go

FROM alpine:3.14.3 as release

WORKDIR /app

COPY --from=build /app/ /app/
COPY --from=build /src/migration/ /app/migration/

CMD [ "./server" ]
