FROM golang:1.15-alpine AS builder
WORKDIR /src/fplbot
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GO111MODULE=on go build -o /app ./cmd

FROM alpine
COPY --from=builder /app /fpl
ENTRYPOINT [ "./fpl" ]
