FROM golang:1.14 as builder
WORKDIR build

# populate with go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# copy go-files
RUN mkdir -p cmd
COPY cmd/main.go ./cmd
RUN mkdir -p internal
COPY internal ./internal

RUN CGO_ENABLED=0 go build \
      -mod=readonly \
      -a \
      -ldflags '-w -s -extldflags "-static"' \
      -o main \
      ./cmd

FROM alpine:latest
WORKDIR website

# copy static resources
RUN mkdir -p templates
COPY templates ./templates
RUN mkdir -p static
COPY static ./static
RUN mkdir -p src
COPY src ./src

# copy binary file
COPY --from=builder go/build/main .

EXPOSE 8080

CMD ["./main"]