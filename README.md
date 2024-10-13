# Go Kafka Project

This repository contains a Kafka-based messaging system implemented in Go. It includes a producer service and a worker service that consume messages from Kafka.

## Prerequisites

Before running the project, ensure that you have the following installed:

- Docker
- Docker Compose
- Go (version 1.18+)

## Getting Started

### 1. Set up Kafka and Zookeeper

The project uses Docker to manage Kafka and Zookeeper. To start Kafka and Zookeeper, navigate to the root of the repository and run:

```bash
docker-compose up
```

This will start all the necessary services for Kafka and Zookeeper.

### 2. Running the Producer
Once the Kafka environment is up and running, follow these steps to run the producer service:

Open a terminal and navigate to the producer directory:

```bash
cd producer
```

Run the producer.go file:

```bash
go run producer.go
```
The producer service will start an HTTP server at http://localhost:3000/api/v1/comments. You can send HTTP POST requests to this endpoint to produce Kafka messages. Example request:

```bash
curl -X POST http://localhost:3000/api/v1/comments -d '{"text": "Third text"}'
```

### 3. Running the Worker
In a separate terminal, follow these steps to run the worker service:

Navigate to the worker directory:

```bash
cd worker
```
Run the worker.go file:

```bash
go run worker.go
```
The worker service will start consuming messages from the Kafka topic and process them.

### Project Structure

* producer: Contains the code for the HTTP producer that accepts comments and pushes them to Kafka.
* worker: Contains the code for the worker that consumes messages from Kafka and processes them.
* docker-compose.yml: Configuration file for setting up Kafka and Zookeeper using Docker.

### How It Works
The producer service exposes an HTTP API (http://localhost:3000/api/v1/comments) that accepts POST requests with comment data.
Upon receiving a request, the producer sends the comment to a Kafka topic.
The worker service consumes messages from the Kafka topic and processes them asynchronously.