# builder stage (gorm migrate用)
FROM golang:1.25.4 AS builder
WORKDIR /src/migrate
COPY migrate/go.mod migrate/go.sum ./
RUN go mod download
COPY migrate .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o /out/migrate

FROM mariadb:12.0
COPY --from=builder /out/migrate /usr/local/bin/migrate
COPY docker/mariadb/wait-and-migrate.sh /usr/local/bin/wait-and-migrate.sh
COPY docker/mariadb/entrypoint-with-migrate.sh /usr/local/bin/entrypoint-with-migrate.sh
RUN chmod +x /usr/local/bin/wait-and-migrate.sh /usr/local/bin/entrypoint-with-migrate.sh
ENTRYPOINT ["/usr/local/bin/entrypoint-with-migrate.sh"]
CMD ["mariadb"]