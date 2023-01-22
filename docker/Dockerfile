FROM golang:1.19.5 as build

RUN apt-get install make git curl
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
COPY . /go/
WORKDIR /go/
RUN make test
RUN make lint
RUN make compile
RUN cp -R ./bin/ports /ports

FROM buildpack-deps:bullseye-scm
RUN apt-get update && \
    apt-get install ca-certificates bash util-linux curl
COPY --from=build /ports /opt/ports/

EXPOSE 8080

WORKDIR /opt/ports/
