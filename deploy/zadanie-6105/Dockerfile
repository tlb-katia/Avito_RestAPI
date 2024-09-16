FROM golang:1.22.3 AS build
WORKDIR /go/src
COPY src/generated-go-server/go ./go
COPY src/generated-go-server/main.go .
COPY src/generated-go-server/go.sum .
COPY src/generated-go-server/go.mod .

ENV CGO_ENABLED=0

RUN go build -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]


# FROM golang:1.22.3 AS build
# WORKDIR /go/src/app
# COPY src/generated-go-server/go.mod src/generated-go-server/go.sum ./
# RUN go mod download
# COPY src/generated-go-server/ ./

# ENV CGO_ENABLED=0

# RUN go build -o /openapi .

# FROM scratch AS runtime
# COPY --from=build /openapi /openapi
# EXPOSE 8080/tcp
# ENTRYPOINT ["/openapi"]
