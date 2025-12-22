FROM golang:alpine as builder
RUN apk --no-cache add tzdata
WORKDIR /go/src/simple_redis_stream_event_listener
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/simple_redis_stream_event_listener/simple_redis_stream_event_listener /app/simple_redis_stream_event_listener
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/simple_redis_stream_event_listener/.env /app/.env
COPY --from=builder /go/src/simple_redis_stream_event_listener/serviceAccountKey.json /app/serviceAccountKey.json

ENTRYPOINT ./simple_redis_stream_event_listener