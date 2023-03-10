# noty

## ๐งโ๐ฌ About

A simple notification application built using Redis and WebSocket.

<img align="right" width="50%" src="./assets/GOPHER_ADADEMY.png">

## ๐ Getting started

### ๐ Table of contents

1. [Documentation](#๐-documentation).
2. [Requirements](#๐งฐ-requirements).
3. [Local development](#๐ป-local-development).
    * [Setting up pre-commit hooks](#๐ช-setting-up-pre-commit-hooks).
    * [Setting up environment variables](#๐-setting-up-environment-variables).
    * [Running server locally](#๐โโ๏ธ-runing-server-locally).
4. [Testing](#๐งช-testing).
5. [Linting](#๐-linting).


### ๐ Documentation

You can manually view the Swagger specification in the [api/](api/) directory or view it at http://localhost:8080/swagger/index.html if the server is running locally.

### ๐งฐ Requirements

* [pre-commit v2.20.0](https://pre-commit.com/#installation)
* [golangci-lint v1.50.1](https://golangci-lint.run/usage/install/)
* [Docker v20.10.17](https://docs.docker.com/engine/install/)

### ๐ป Local development

For local development, you need to repeat the following steps. ๐

1. [Setting up pre-commit hooks](#๐ช-setting-up-pre-commit-hooks).
2. [Setting up environment variables](#๐-setting-up-environment-variables).
3. [Running server locally](#๐โโ๏ธ-runing-server-locally).

#### ๐ช Setting up pre-commit hooks

To install the pre-commit hooks, run the `pre-commit install` command and check the installation with the `pre-commit run` command.

#### ๐ Setting up environment variables

To set up environment variables, you need to create a **local.env** file with the environment variables from the [example.env](example.env) file.

#### ๐โโ๏ธ Runing server locally

Once you [set up environment variables](#-setting-up-environment-variables), you can use the `make local-run` command, which creates a docker image and runs the server in the docker container.

### ๐งช Testing

To run tests, use `make test` command.

### ๐ Linting

To run linters, use `make lint` command.
