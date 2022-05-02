FROM golang:alpine AS builder
RUN apk add git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
LABEL maintainer="Jevgenij Bogdasic <jevbogd@gmail.com>"

WORKDIR /app

#add code to container
COPY . .

#build app
RUN go build -o main .


# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .

FROM alpine

COPY --from=builder /dist/main /
COPY --from=builder /app/.env .

ENTRYPOINT ["/main"]
