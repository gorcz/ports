# Ports

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation and Usage](#installation-and-usage)
  - [Prerequisites](#prerequisites)
  - [Build and run](#build-and-run)
  - [Development environment](#development-environment)

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

## Development environment

Project is using `vendor` mode for keeping track of dependencies. If a need for new dependency arise please:

* update `go.mod` file
* run `go mod tidy` - `go.sum` should be updated
* run `go mod vendor`
* commit all files added/changed in `vendor` directory.

The same rule applies when it comes to mocks. If any interface is changed that affects mocks, please commit
generated `.go` files.
