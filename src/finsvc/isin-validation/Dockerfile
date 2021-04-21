FROM golang:latest as builder
LABEL maintainer="Anton Efimow shimmeg@yandex.ru"
WORKDIR /app
COPY ../.. .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

####### Start a new stage #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]