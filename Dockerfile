FROM golang:1.23.1 AS build

RUN apt-get update && apt-get install -y git make --no-install-recommends

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/server

RUN chmod +x /app/server

FROM debian:bookworm-slim as dev

COPY --from=build /go/bin/migrate /usr/local/bin/
COPY --from=build /usr/bin/make /usr/local/bin/
COPY --from=build /app/server /app/server
COPY --from=build /app/Makefile /app/Makefile
COPY --from=build /app/db/migrations /app/db/migrations
COPY --from=build /app/db/scripts/migrate_up.sh /app/db/scripts/migrate_up.sh
COPY --from=build /app/db/scripts/migrate_down.sh /app/db/scripts/migrate_down.sh

WORKDIR /app

RUN chmod +x /app/server
RUN chmod +x /app/db/scripts/migrate_up.sh

EXPOSE 8080
