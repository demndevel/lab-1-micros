# 1
FROM golang:1.24.1-alpine AS build-stage

ENV GOCACHE=/root/.cache/go-build

WORKDIR /src

COPY ./src ./src
COPY go.mod ./

RUN go mod tidy

RUN --mount=type=cache,target=/root/.cache/go-build \
   go build -o output ./src/main.go

# 2
FROM alpine:3.21

WORKDIR /app

COPY --from=build-stage /src/output ./

EXPOSE 6080

CMD ["./output"]
