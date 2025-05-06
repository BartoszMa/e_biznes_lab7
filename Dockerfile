FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY controllers/ ./controllers/
COPY models/ ./models/
COPY routes/ ./routes/
COPY service/ ./service/

RUN apk add --no-cache gcc musl-dev sqlite-dev && \
    CGO_ENABLED=1 GOOS=linux go build -o myapp .

FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=build /app/myapp .
COPY --from=build /app/products.db .

RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 4000
ENTRYPOINT ["/app/myapp"]
