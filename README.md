# Gocker - A Docker-like Container System in Go

Welcome to **Gocker**! This project is an educational initiative to learn Go by developing a Docker-like container system from scratch. By building this project, we aim to cover all the features and concepts of the Go programming language, as well as delve into production-grade tools and design patterns.

## Table of Contents

- [Gocker - A Docker-like Container System in Go](#gocker---a-docker-like-container-system-in-go)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Container](#creating-a-container)
    - [Starting a Container](#starting-a-container)
    - [Stopping a Container](#stopping-a-container)
    - [Removing a Container](#removing-a-container)
    - [Listing Containers](#listing-containers)
    - [Building an Image](#building-an-image)
    - [Listing Images](#listing-images)
    - [Removing Images](#removing-images)
    - [Development Steps](#development-steps)
    - [Architecture](#architecture)
    - [Contributing](#contributing)
    - [License](#license)

## Introduction

Gocker is a lightweight container system inspired by Docker. It aims to provide a simplified version of containerization functionalities, allowing users to create, manage, and run containers. This project is designed for learning purposes and is not intended for production use.

## Features

- **Basic Container Creation:** Create and manage containers.
- **Resource Isolation:** CPU and memory limits for containers.
- **Networking:** Basic networking capabilities for containers.
- **Image Management:** Build, list, and remove container images.
- **Container Lifecycle Management:** Start, stop, and remove containers.
- **Simple Orchestration:** Manage multiple containers and their interactions.

## Installation

To get started with Gocker, follow these steps:

1. **Clone the repository:**
    ```bash
    git clone https://github.com/keshan/gocker.git
    cd gocker
    ```

2. **Build the project:**
    ```bash
    go build -o gocker
    ```

3. **Run Gocker:**
    ```bash
    ./gocker
    ```

## Usage

### Creating a Container

```bash
./gocker create -name mycontainer -image myimage
```
### Starting a Container
```bash
./gocker start mycontainer
```
### Stopping a Container
```bash
./gocker stop mycontainer
```
### Removing a Container
```bash
./gocker rm mycontainer
```
### Listing Containers
```bash
./gocker ps
```
### Building an Image
```bash
./gocker build -t myimage .
```
### Listing Images
```bash
./gocker images
```
### Removing Images
```bash
./gocker rmi myimage
```
### Development Steps
- Basic Container Creation
- Resource Isolation
- Networking
- Image Management
- Container Lifecycle Management
- Simple Orchestration

### Architecture

Gocker is structured to mirror the architecture of Docker, albeit in a simplified form. The core components include:

- **CLI:** Command-line interface for interacting with Gocker.
- **Container:** Manages the lifecycle of containers.
- **Image:** Handles the creation and management of container images.
- **Network:** Provides basic networking functionalities for containers.
- **Volume:** Manages data volumes for persistent storage.
- **Runtime:** Core logic for running containers with resource isolation.
- **Logging:** Captures and manages logs from containers.

### Contributing

We welcome contributions from the community! To contribute to Gocker, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Create a pull request to the main repository.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
