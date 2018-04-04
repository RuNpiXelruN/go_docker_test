# build stage
FROM golang:alpine AS build-env
MAINTAINER Justin Davidson <justindavidson23@gmail.com>
ADD . /src
RUN cd /src && go build -o goappbuild

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goappbuild /app/
ENTRYPOINT ./goappbuild
