FROM golang:1.22 as builder

RUN mkdir app
WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

ARG APP_VERSION
RUN go build -o bin/products_server -ldflags "-X main.Version=$APP_VERSION" ./cmd/server


FROM busybox

RUN mkdir app

COPY --from=builder /app/bin/* /app
COPY migrations /migrations

#copy static files if needed

CMD ["./app/products_server"]
