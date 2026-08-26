FROM golang:1.23-bookworm AS build
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o /out/practice .
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /out/practice /app/practice
RUN mkdir -p /app/data
EXPOSE 8080
ENTRYPOINT ["/app/practice"]
