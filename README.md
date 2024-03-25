# Pantori

This project aims to help users manage goods and expiration dates efficiently, preventing food waste. The system consists of a Golang API for goods registration, a NoSQL database (currently AWS DynamoDB in production) for data storage, and a Flutter frontend web application for user interaction (served via Nginx).

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project-locally)
- [Project Structure](#project-structure)
- [External Dependencies](#external-dependencies)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Frontend](#frontend)
- [Infrastructure](#infrastructure)
- [Versions and Future Features](#versions-and-future-features)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Make sure you have the following software installed on your machine:

- Docker and Docker Compose: [Install Docker](https://www.docker.com/get-started)
- Go: [Install Go](https://golang.org/doc/install)

### Installation

Clone the repository:

```bash
git clone https://github.com/almenarar/pantori-backend.git
cd pantori-backend
```

### Running the project locally

Change mode.database value in *config.json* to "sql".
Complete .env.example as follows:
- JWT_KEY: any value you want;
- UNSPLASH_KEY: if you have one, to enable custom images. [Click here](https://unsplash.com/documentation#creating-a-developer-account) to know more, it's free!

Use the provided Makefile to build and run the containers:

```bash
make run
```

This command will start the Golang API and a MySQL database. Default user is *foo* with pwd *bar*. This is a temporary way to manage users.

## Project Structure

### Hexagonal Architecture Overview

This project follows [Hexagonal Architecture](https://www.google.com/search?q=hexagonal+architecture), separating the core application logic from external concerns.
In Hexagonal Architecture, the application is divided into three main layers:

1. **Core Application Logic (Hexagon):** Contains the domain model, business rules, and application services. This layer is independent of external details.

2. **Ports:** Interfaces defining how the core application interacts with the external world. Examples include repositories, services, and event listeners.

3. **Adapters:** Implementations of the port interfaces, providing connections to the external world. Adapters convert data between the core and external systems.

### Project Structure

- **`internal/`**: All code that should NOT be import/used by another code.
  - **`auth/`**: Code for user management and token generation.
  - **`domains/`**: Keeps code grouped by business domains.
    - **`core/`**: Contains the main application logic, no imports allowed here.
      - `domain.go`: Defines the entity and API models.
      - `service.go`: Contains application logic.
      - `ports.go`: Interfaces for data access.

    - **`handlers/`**: Implementation of entry ports interfaces, converts user input in to domain object and pass to service.
      - `http.go`: HTTP routes, input validation.

    - **`infra/`**: Implementations of the port interfaces, connecting the core to external systems.
      - `dynamodb.go`: Implements database port to AWS DynamoDB.
      - `sql.go`: Implements database port to MySQL.

- **`cmd/`**: Application entry point.
  - **`middlewares`**: Logger configuration and API Authorization
  - **`routes`**: Gin start, connects the routes to server.
  - **`swagger`**: Swagger related files.
  - `main.go`: Starts swagger and logger, load parameters and init server.

- **`docker-compose.yml`**: Docker services and configurations.

- **`Makefile`**: Commands for building, testing and running the project.

- **`README.md`**: Project documentation.

- **`CONTRIBUTING.md`**: Guidelines for contributing to the project.

- **`CODE_OF_CONDUCT.md`**: Code of Conduct for community guidelines.

- **`LICENSE`**: GNU General Public License (GPL) for project licensing.

## External Dependencies

- [Zerolog](https://github.com/rs/zerolog);
- [Gin](https://github.com/gin-gonic/gin);
- [Swag](https://github.com/swaggo/swag);
- [JWT](https://github.com/golang-jwt/jwt);
- [AWS SDK for Go](https://github.com/aws/aws-sdk-go);
- [GORM](https://gorm.io/index.html);


## API Documentation

The API documentation is available through Swagger. After running the project, visit http://localhost:8080/swagger/index.html to explore the API endpoints.

### Updating Swagger

You can update the swagger documentation with the following command:

```bash
cd ./cmd/api && ~/go/bin/swag init --parseInternal -d .,../../internal -o swagger/
```

## Testing

This project only has unit tests for now. We understand that all code inside **core/** directories should be unit tested since it don't have dependencies and it's pure business logic. 
I consider integration tests to be too costly to make at the moment. 
Functional tests to production will be written in future releases.
You can run unit tests with:

```bash
make unit
```

You can find test files in the same directory of the code they are testing. Test files has a **_test.go** prefix. All tests are based in [Table Driven Tests](https://www.google.com/search?q=table+driven+tests).

## Frontend

The frontend application interacts with the API to provide a user-friendly interface for goods management. Check the dedicated frontend repository for more details [here](https://github.com/almenarar/pantori-frontend).

## Infrastructure

The MySQL docker image is only for local test and developement purporse. Currently using AWS DynamoDB in Production. Infrastructure-related scripts and configurations are stored in the infrastructure repository. Please check it for more details [here](https://github.com/almenarar/pantori-infra).

## Versions and Future Features

This section outlines the current version of the Goods Expiry Management System (Pantori) and provides insights into planned features for future releases.

### Current Version

The current version of the project is `v1.0.0`. You can check the [release page](https://github.com/almenarar/pantori-backend/releases) for more details about each release.

### Roadmap

#### Version `v1.1.0` (Next Release)

- [ ] **Feature:** Multiple workspaces, allowing diferent users to manage their goods separately.
- [ ] **Enhancement:** Removing user data from enviroment variables and storing it in database with proper CRUD.

#### Version `v1.2.0` (Future Release)

- [ ] **Feature:** Allow user to create and manage their own categories.
- [ ] **Enhancement:** Enable each good to have up to three categories.

#### Version `v1.3.0` (Future Release)

- [ ] **Feature:** Integrate a notification system for impending expiration dates.

#### Version `v2.0.0` (Future Version)

- [ ] **Enhancement:** Record the amount of food you still have.
- [ ] **Enhancement:** Mark a good as missing.
- [ ] **Feature:** Generate a shopping list from missing goods.


## Contributing

We welcome contributions! Please follow our [Contribution Guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [GNU General Public License (GPL)](LICENSE).

