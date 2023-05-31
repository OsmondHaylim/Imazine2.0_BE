# Compile stage
FROM golang:1.20 AS build-env

ADD . /Imazine2.0_BE
WORKDIR /Imazine2.0_BE

RUN go build -o /server

# Final stage
FROM debian:buster

# certs
RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect api.imgbb.com:443 </dev/null 2>/dev/null | openssl x509 -outform PEM > ${cert_location}/api.imgbb.crt
RUN update-ca-certificates

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /
RUN mkdir -p /download_cache

CMD ["/server"]