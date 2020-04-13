FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*



WORKDIR /app/psychologist


# install migrate which will be used by entrypoint.sh to perform DB migration
ARG MIGRATE_VERSION=4.10.0
COPY psychologist/dist/migrate/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate

# copy source files and build the binary
COPY psychologist/. . 
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /usr/local/bin/migrate /usr/local/bin
COPY --from=build /app/psychologist/migrations ./migrations/
COPY --from=build /app/psychologist/service .
COPY --from=build /app/psychologist/cmd/service/entrypoint.sh .
COPY --from=build /app/psychologist/configs/*.yaml ./config/
RUN ls -la
ENTRYPOINT ["./entrypoint.sh"]
