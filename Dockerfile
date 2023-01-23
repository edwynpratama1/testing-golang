FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN cd /build/serverLocalGo && go build

EXPOSE 8080

ENTRYPOINT [ "/build/serverLocalGo/main" ]