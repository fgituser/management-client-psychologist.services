FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*



WORKDIR /app/operator

# copy source files and build the binary
COPY operator/. . 
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/operator/service .
COPY --from=build /app/operator/cmd/service/entrypoint.sh .
COPY --from=build /app/operator/configs/*.yaml ./config/
RUN ls -la
ENTRYPOINT ["./entrypoint.sh"]
