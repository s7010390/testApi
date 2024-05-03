FROM golang:1.21.6 as builder

WORKDIR /app/test-api
COPY . /app/test-api

WORKDIR /app/test-api/src

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/test-api .

FROM alpine:latest as runner

RUN apk update && apk add --no-cache tzdata

WORKDIR /test-api

COPY --from=builder /app/test-api/build/test-api /usr/bin/test-api
COPY --from=builder /app/test-api/src/config /config
EXPOSE 3000

CMD ["test-api", "serve"]
