FROM golang:1.19 AS BUILDER


WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
 GOOS=linux go build -o mvc-api .

FROM golang:1.19-alpine3.15 as runner

RUN adduser -D richard

COPY --from=BUILDER /app/mvc-api /app/mvc-api

RUN chown -R richard:richard /app
RUN chmod +x /app/mvc-api

EXPOSE 8080

USER richard

CMD ["./mvc-api"]