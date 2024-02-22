# syntax=docker/dockerfile:1

FROM golang:1.14

WORKDIR /app

COPY . .

ENTRYPOINT ["env", "--chdir=/app", "-S", "./build.sh"]