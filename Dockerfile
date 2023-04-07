FROM golang:alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine:3.17 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main .

EXPOSE 8000

ENTRYPOINT ["./main"]
