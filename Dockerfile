FROM golang:1.23.2 AS builder
COPY . .
RUN go mod download
RUN make build

FROM httpd:latest
COPY --from=builder /go/ui /usr/local/apache2/htdocs/
RUN echo "AddType application/wasm .wasm" >> /usr/local/apache2/conf/httpd.conf
