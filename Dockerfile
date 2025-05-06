FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk add --no-cache gcc musl-dev sqlite-dev && \
    CGO_ENABLED=1 GOOS=linux go build -o myapp .

FROM alpine:3.19

WORKDIR /app

COPY --from=build /app/myapp .
COPY --from=build /app/products.db .

EXPOSE 4000
ENTRYPOINT ["/app/myapp"]
