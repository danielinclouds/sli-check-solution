# -----------------------------
# Build container
# -----------------------------
FROM golang:1.12.7 AS builder
WORKDIR /go/src/github.com/danielinclouds/app/

# Download dependencies
# RUN go get -d -v golang.org/x/net/html  

COPY main.go .
# RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# -----------------------------
# Final container
# -----------------------------
FROM alpine:latest 

ARG PORT
ENV APP_PORT=${PORT}

RUN apk --no-cache add ca-certificates

USER 65534
WORKDIR /app/
COPY --from=builder /go/src/github.com/danielinclouds/app/app .
EXPOSE ${PORT}
CMD ["./app"]