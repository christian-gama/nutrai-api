# Exception (Panic) Handling in Application

This application is designed to handle exceptions (panics) robustly, ensuring that no panic goes unnoticed, and the application remains highly available. This document provides a detailed explanation of how exception handling is structured and the various components involved.

## Middleware
The application uses a middleware to intercept and handle exceptions during request handling. If a panic occurs, this middleware catches it, logs it, and returns a 500 Internal Server Error response to the client. In a production environment, the client only receives a generic error message for security reasons. The middleware uses a handler to delegate the task of logging the exception details.

## Message Consumer
The message consumer is responsible for consuming exception messages from the messaging system. When it receives a message, it deserializes the JSON payload into an `Exception` object and saves it to the database using the `Exception` repository.

## Exception Repository
The `Exception` repository provides an interface for working with `Exception` objects in a database. It offers methods to save and delete exceptions.

## Exception Model
This is the domain model representing an exception. It includes fields for an ID, the exception message, the stack trace, and the time when the exception occurred.

## Persistence
The persistence layer provides a concrete implementation of the `Exception` repository for a SQL database using the GORM library. It defines how to save and delete exceptions from the database.

## Summary
In summary, when an exception occurs during request handling, the application's middleware catches it, logs it, and responds to the client with an error status. The exception details are saved to the application's messaging system by a command handler. A consumer then receives these details from the messaging system, deserializes them into an `Exception` object, and saves them to the database using an `Exception` repository. The database operations are implemented in a persistence layer using GORM. This way, every panic in the application is logged and saved, ensuring high availability and robust error tracking.