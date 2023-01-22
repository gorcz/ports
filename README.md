# Ports

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation and Usage](#installation-and-usage)
    - [Prerequisites](#prerequisites)
    - [Build and run](#build-and-run)
    - [Docker image](#docker-image)
    - [Development environment](#development-environment)
- [Technical test requirements](#technical-test-requirements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Installation and Usage

## Prerequisites

* Installed Go version > 1.19
* [GolangCI-Lint](https://github.com/golangci/golangci-lint) 1.50.1 (optional)
* Docker (optional)

You can use [asdf](https://asdf-vm.com/) to manage development environment

## Build and run

* `$ make build` - build binary executable
* `$ make precommit` - ensure code is ready to be committed
* `$ make run` - run port service locally on port 8080
* `$ make test` - run unit tests
* `$ make lint` - run various linters

## Docker image

* `$ make image` - build Docker image without publishing it to external registry
* `$ make docker-image-run` - run ports docker container locally

## Development environment

Project is using `vendor` mode for keeping track of dependencies. If a need for new dependency arise please:

* update `go.mod` file
* run `go mod tidy` - `go.sum` should be updated
* run `go mod vendor`
* commit all files added/changed in `vendor` directory.

The same rule applies when it comes to mocks. If any interface is changed that affects mocks, please commit
generated `.go` files.

# Technical test requirements

- Created service exposes HTTP API on `loclhost:8080/port` endpoint which support POST request. Payload of that request
  should contain ports, similar to `ports.json` provided as example.
- For storing ports, in-memory datastore has been implemented based on simple synchronized map. Used datastore
  abstraction allows to replace that implementation later with any DB technology. I did not focus on the database itself
  because it is a technical detail which can be decided later.
- Due to the unpredictable size of the input data, need to account for duplicate ports and restrictions for service
  memory limit, I used a custom stream parser for JSON. I used an iterator as an interface to this parser. Thus, request
  payload processing is carried out sequentially port by port, allocating only memory for the necessary buffers.
- When it comes to tests, I covered the JSON stream parser and happy path for the port service by unit tests. It would
  be nice to cover the port controller as well, but I didn't for lack of time.
- README.md explains how to run service and test it.
- A Docker image after compilation contains only binary file for service. Two phase image building is applied.
- Service handles certain signals correctly and shutdown gracefully.