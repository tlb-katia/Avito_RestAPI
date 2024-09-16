FROM golang:1.22.3 AS build
WORKDIR /go/src
COPY src/generated-go-server/go ./go
COPY src/generated-go-server/main.go .
COPY src/generated-go-server/go.sum .
COPY src/generated-go-server/go.mod .
COPY .env .

ENV CGO_ENABLED=0

RUN go build -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
COPY --from=build go/src/.env .
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
