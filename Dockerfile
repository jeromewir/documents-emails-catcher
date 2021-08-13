##
## Build
##

FROM golang:1.16-buster AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/invoices-fwder
RUN chmod +x /app/invoices-fwder

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/invoices-fwder /app/invoices-fwder

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/invoices-fwder"]