FROM golang:alpine AS builder
RUN apk add git

#adding needed env variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
LABEL maintainer="Jevgenij Bogdasic <jevbogd@gmail.com>"
#move to /build
WORKDIR /app

#add code to container
COPY . .

#build app
RUN go build -o main .


# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .

# Build a small image
FROM alpine

COPY --from=builder /dist/main /
COPY --from=builder /app/.env .


    #change database url variable to match your needs

# Command to run
ENTRYPOINT ["/main"]
