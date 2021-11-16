FROM golang:1.17-alpine AS build
ADD . /olaf
ENV CGO_ENABLED=0
WORKDIR /olaf
RUN go build -o olaf ./cmd/olaf

FROM alpine:latest
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

COPY --from=build /olaf/olaf /olaf/olaf
EXPOSE 9999

