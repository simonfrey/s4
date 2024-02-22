# syntax=docker/dockerfile:1

FROM docker.io/golang:1.14

WORKDIR /app

COPY . .

ENTRYPOINT ["env", "--chdir=/app", "-S", "./build.sh"]