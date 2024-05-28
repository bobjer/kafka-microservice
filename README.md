# Kafka Microservice

## Overview

This microservice takes messages from Kafka topic `input_topic`, does some processing, and sends results to another Kafka topic `output_topic`. It is made to handle many messages fast.

## Tasks Breakdown

### Requirements

1. Take messages from Kafka topic `input_topic`.
2. Process each message:
   - Check the timestamp is not older than 24 hours.
   - Add a status field to the JSON object based on the data length (e.g., `status: "valid"` if the length is more than 10 characters, else `status: "invalid"`).
3. Send processed messages to topic `output_topic`.
4. Handle errors for bad messages, messages that can't be processed, and issues with Kafka connection.
5. Service should handle restarts and keep processing from where it left.
6. Service should be scalable.

### Tasks

1. **Setup Project**
   - Create Go project.
   - Setup modules and dependencies.
   - Create Dockerfile.

2. **Kafka Consumer**
   - Initialize Kafka consumer group.
   - Consume messages from `input_topic`.
   - Handle errors and retries.
   - Gracefully handle shutdowns.

3. **Message Processing**
   - Parse messages to JSON.
   - Check timestamp.
   - Add status field based on `data` length.
   - Handle and log errors.

4. **Kafka Producer**
   - Initialize Kafka producer.
   - Send processed messages to `output_topic`.
   - Handle errors and retries.

5. **Integration and Error Handling**
   - Integrate consumer and producer.
   - Implement error handling and logging.
   - Ensure service can recover from failures.

6. **Scalability**
   - Design to scale horizontally.
   - Handle high message volume.
   - Add metrics and monitoring.

7. **Documentation and Testing**
   - Write README.
   - Add integration tests.
   - Create architecture diagram.

## Setup and Running

### System Requirements

- Docker
- Docker Compose

### Instructions

1. Clone the repo:

    ```sh
    git clone https://github.com/bobjer/kafka-microservice.git
    cd kafka-microservice
    ```

2. Build and run services with Docker Compose:

    ```sh
    docker-compose up --build
    ```

3. To see logs of kafka microservice:

    ```sh
    docker-compose logs -f kafka-microservice
    ```

4. To stop services:

    ```sh
    docker-compose down
    ```

### Env Variables

- `KAFKA_BROKER`: Kafka broker address (e.g., `kafka:9092`)
- `GROUP_ID`: Kafka consumer group ID (e.g., `my-group`)
- `INPUT_TOPIC`: Kafka input topic (e.g., `input_topic`)
- `OUTPUT_TOPIC`: Kafka output topic (e.g., `output_topic`)
- `NUM_CONSUMER_WORKERS`: Number of consumer workers (e.g., `3`)
- `NUM_PRODUCER_WORKERS`: Number of producer workers (e.g., `3`)

## Design Choices

### Project Structure

The project is structured into several packages to promote modularity, ease of maintenance, and extensibility:

- **consumer**: Contains the logic for consuming messages from Kafka. This package is responsible for initializing the Kafka consumer, handling message consumption, and managing consumer group sessions.
  
  - `consumer.go`: Contains the implementation of the Kafka consumer, including setup, message consumption, and graceful shutdown handling.

- **producer**: Contains the logic for producing messages to Kafka. This package handles the initialization of the Kafka producer and the logic for sending messages to the `output_topic`.

  - `producer.go`: Implements the Kafka producer, including message production and error handling.

- **message**: Contains the logic for processing messages. This includes validating the timestamp, adding a status field based on data length, and handling any errors during processing.
  
  - `message.go`: Implements message validation and processing logic.

- **handler**: A helper package that provides functions for handling messages. It abstracts the processing logic and makes it reusable and testable.

  - `handler.go`: Provides the `HandleMessage` function that applies processing logic to the messages.

- **logger**: Provides a centralized logging mechanism using the logrus library. This ensures consistent logging across the application and simplifies debugging.

  - `logger.go`: Initializes and configures the logger.

- **metrics**: Contains the logic for exposing application metrics. This is useful for monitoring the health and performance of the microservice.

  - `metrics.go`: Sets up Prometheus metrics for tracking processed messages.

### Why This Structure?

- **Separation of Concerns**: Each package has a single responsibility, making the codebase easier to understand and modify.
- **Reusability**: Common functionality, such as logging and message handling, is encapsulated in reusable packages.
- **Testability**: By separating the logic into distinct packages, it becomes easier to write unit tests for each component.
- **Scalability**: The design allows the service to scale horizontally. By adding more instances of the microservice, you can handle a higher volume of messages without changing the core logic.
- **Maintainability**: A modular structure makes it easier to locate and fix bugs, as well as add new features.

## Running Integration Tests

Integration tests running automaticaly when the service is startedand are set up to produce test messages to the `input_topic` and verify the service logs. The results of the integration tests can be viewed in the logs of the service.

To view the results of the integration tests, run the following command:
    ```sh
    docker-compose logs -f kafka-microservice
    ```

## Architect diargam

+---------------------------+
|                           |
|      Kafka Broker         |
|                           |
|   +-------------------+   |
|   | input_topic       |   |
|   +---------+---------+   |
|             |             |
|             v             |
|   +---------+---------+   |
|   | Kafka Consumer   |    |
|   | (kafka-microservice)  |
|   +---------+---------+   |
|             |             |
|             v             |
|   +---------+---------+   |
|   | Message Processor |   |
|   +---------+---------+   |
|             |             |
|             v             |
|   +---------+---------+   |
|   | Kafka Producer   |    |
|   | (kafka-microservice)  |
|   +---------+---------+   |
|             |             |
|             v             |
|   +---------+---------+   |
|   | output_topic      |   |
|   +-------------------+   |
|                           |
+---------------------------+
