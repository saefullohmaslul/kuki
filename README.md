# K.U.K.I
Golang Playground with gRPC and Echo

Kuki is a playground repository that demonstrates how to use Golang to build both gRPC and RESTful APIs using the Echo framework. It provides examples and explanations to help you understand the concepts of building APIs with these technologies.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Kuki is a project aimed at showcasing the usage of Golang with gRPC and Echo framework. It provides practical examples and step-by-step explanations to help you build APIs effectively using these technologies.

## Features

- Demonstrates building gRPC services with handler functions.
- Illustrates building RESTful APIs using Echo.
- Provides clear examples for various API operations: Create, Read, Update, Delete (CRUD).
- Promotes understanding of how to integrate gRPC and RESTful APIs within the same project.
- Supports interactive learning for beginners in the world of Golang APIs.

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/saefullohmaslul/kuki.git
cd kuki
```

2. Install the required dependencies:

```bash
go install github.com/cosmtrek/air@v1.45.0

go mod tidy
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

## Usage

1. Running service using

```bash
make build
make run
```

2. To run in development mode (hot loader), using this command

```bash
make run-dev
```

3. To generate proto files, using this command

```bash
make protoc
```

4. You can now access the gRPC and REST API endpoints as described in the [Endpoints](#endpoints) section below.

## Endpoints

### gRPC Endpoints

- Create Todo: `POST http://localhost:50051/todos`
- Read Todo: `GET http://localhost:50051/todos/{id}`
- Update Todo: `PUT http://localhost:50051/todos`
- Delete Todo: `DELETE http://localhost:50051/todos/{id}`

### REST API Endpoints

- Create Todo: `POST http://localhost:1323/todos`
- Read Todo: `GET http://localhost:1323/todos/:id`
- Update Todo: `PUT http://localhost:1323/todos/:id`
- Delete Todo: `DELETE http://localhost:1323/todos/:id`

## Contributing

Contributions to Kuki are welcome! Feel free to open issues for feature requests, bug fixes, or improvements. Pull requests are also appreciated.

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/my-feature`.
3. Make your changes and commit them: `git commit -am 'feat(deps): Add some feature'`. We are using conventional commit for standardization commit messages.
4. Push the branch: `git push origin feature/my-feature`.
5. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
