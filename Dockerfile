FROM golang:1.18 AS builder

WORKDIR /src

COPY . ./

RUN go build -o server ./cmd/server

FROM gcr.io/distroless/base-debian10

COPY --from=builder /src/server /server

CMD ["/server"]
