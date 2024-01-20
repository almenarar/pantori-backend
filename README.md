# Pantori

This project aims to help users manage goods and expiration dates efficiently, preventing food waste. The system consists of a Golang API for goods registration, a MySQL database for data storage, and a Flutter frontend application for user interaction.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Frontend](#frontend)
- [Infrastructure](#infrastructure)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Make sure you have the following software installed on your machine:

- Docker: [Install Docker](https://www.docker.com/get-started)
- Go: [Install Go](https://golang.org/doc/install)

### Installation

Clone the repository:

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### Running the Project

Use the provided Makefile to build and run the containers:

```bash
make run
```

This command will start the Golang API and MySQL database.

## Project Structure

### Hexagonal Architecture Overview

This project follows Hexagonal Architecture, separating the core application logic from external concerns.
In Hexagonal Architecture, the application is divided into three main layers:

1. **Core Application Logic (Hexagon):** Contains the domain model, business rules, and application services. This layer is independent of external details.

2. **Ports:** Interfaces defining how the core application interacts with the external world. Examples include repositories, services, and event listeners.

3. **Adapters:** Implementations of the port interfaces, providing connections to the external world. Adapters convert data between the core and external systems.

### Project Structure


- **`core/`**: Contains the core application logic.
  - `domain/`: Defines the domain model.
  - `service/`: Contains application services.
  - `repository/`: Interfaces for data access.

- **`ports/`**: Interfaces (ports) defining how the core interacts with the external world.
  - `http/`: HTTP-related interfaces.
  - `database/`: Database-related interfaces.

- **`adapters/`**: Implementations of the port interfaces, connecting the core to external systems.
  - `http/`: HTTP adapters.
  - `database/`: Database adapters.

- **`cmd/`**: Application entry point.
  - `main.go`: Application bootstrap.

- **`docker-compose.yml`**: Docker services and configurations.

- **`Makefile`**: Commands for building and running the project.

- **`README.md`**: Project documentation.

- **`CONTRIBUTING.md`**: Guidelines for contributing to the project.

- **`CODE_OF_CONDUCT.md`**: Code of Conduct for community guidelines.

- **`LICENSE`**: GNU General Public License (GPL) for project licensing.


## API Documentation

The API documentation is available through Swagger. After running the project, visit http://localhost:8080/swagger/index.html to explore the API endpoints.

## Frontend

The frontend application interacts with the API to provide a user-friendly interface for goods management. Check the dedicated frontend repository for more details.

## Infrastructure

The MySQL docker image here is only for local test and developement purporse. Production infrastructure-related scripts and configurations are stored in the infrastructure repository. Please check it for more details.

## Versions and Future Features

This section outlines the current version of the Goods Expiry Management System and provides insights into planned features for future releases.

### Current Version

The current version of the project is `v1.0.0`. You can check the [release page](https://github.com/your-username/your-repo/releases) for more details about each release.

### Roadmap

#### Version `v1.1.0` (Next Release)

- [ ] **Enhancement:** Implement user authentication for secure access.
- [ ] **Feature:** Add support for multiple user accounts.
- [ ] **Bug Fix:** Resolve reported issues related to data synchronization.

#### Version `v1.2.0` (Future Release)

- [ ] **Feature:** Integrate a notification system for impending expiration dates.
- [ ] **Enhancement:** Improve the user interface for better user experience.
- [ ] **Infrastructure:** Explore options for container orchestration (e.g., Kubernetes).


## Contributing

We welcome contributions! Please follow our [Contribution Guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [GNU General Public License (GPL)](LICENSE).

