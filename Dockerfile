FROM golang:1.21-alpine AS builder

RUN apk add --no-cache make git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/reflexnetd /usr/local/bin/

EXPOSE 26656 26657 26660

CMD ["reflexnetd", "start"]

