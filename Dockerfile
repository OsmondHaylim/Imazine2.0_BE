# Compile stage
FROM golang:1.20 AS build-env

ADD . /Imazine2.0_BE
WORKDIR /Imazine2.0_BE

RUN go build -o /server

# Final stage
FROM debian:buster

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /
RUN mkdir -p /download_cache

CMD ["/server"]