FROM golang:1.26.4 AS BUILDER

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \ 
GOOS=linux go build -o meuprimeirocrudgo .

FROM golang:1.26.4-alpine AS RUNNER

RUN adduser -D Usuario

COPY --from=BUILDER /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R Usuario:Usuario /app
RUN chmod +x /app/meuprimeirocrudgo

EXPOSE 8080

USER Usuario

CMD ["/app/meuprimeirocrudgo"]
