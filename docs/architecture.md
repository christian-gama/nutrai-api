# Architecture
Here is a brief description of the folder structure for the Nutrai API.

## nutrai-api
```
root
 ┣ cmd
 ┣ config
 ┣ internal
 ┣ migration
 ┣ pkg
 ┗ scripts
```

### cmd
This folder contains all the executable command line applications.
  
### config
Holds all application configuration files, including `env` for environment configurations.

### internal
This is the main package that holds the core business logic of the API.
  
### migration
Contains the sql migration files.

### pkg
Contains all the shared packages used in the application.

### scripts
Contains all the scripts used to setup the project, such as git hooks and docker scripts.

## The Architecture
The architecture is organized according to the principles of Domain-Driven Design (DDD) and Clean Architecture.
- **Domain Layer**: The core business logic is encapsulated in the domain layer, which includes entities, value objects, and services.
- **Application Layer**: This layer coordinates high-level activities such as application (use case) services and domain events.
- **Infrastructure Layer**: This layer provides generic technical capabilities that support the higher layers, for instance, message queues, databases, web services, mailing systems etc.
- **API Layer**: This is the layer where all HTTP related features are located, including controllers, routes, and middleware.
` files contains module-specific functionality or configurations.