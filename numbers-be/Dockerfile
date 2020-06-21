FROM golang:1.13 AS builder

LABEL maintainer="faldy.findraddy@gmail.com"

# Always set workdir into application root
WORKDIR /app

# Copy the source code into container for compiling
COPY . /app/


# Compile the binary
RUN go get -v -d .
RUN go build main.go

FROM gcr.io/distroless/base
COPY --from=builder /app/main /app
USER 1000
CMD ["/app"]
