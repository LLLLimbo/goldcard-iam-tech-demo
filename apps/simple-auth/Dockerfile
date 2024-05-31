FROM golang:1.22 AS builder

WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=0 go build -o /app


FROM scratch
WORKDIR /root
COPY --from=builder /app .
CMD ["/root/app"]