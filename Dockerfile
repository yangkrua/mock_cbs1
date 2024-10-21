FROM golang:1.23-alpine


ENV SSL_CERT_DIR=/etc/ssl/certs

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /mock_rlcs


EXPOSE 9343

ENTRYPOINT [ "/mock_rlcs" ]