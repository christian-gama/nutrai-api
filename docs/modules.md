# Modules
To an overview, a module in your application follows a specific structure. Each module usually consists of a variety of folders that categorize the different responsibilities and functionalities within the module.

Let's consider the folder structure:
```
└── module
    ├── api
    │   ├── http
    │   │   ├── controller
    │   │   ├── middleware
    │   │   └── routes
    │   ├── grpc
    │   └── ...
    ├── app
    │   └── command
    │   └── query
    │   └── service
    ├── domain
    │   ├── model
    │   ├── repo
    │   ├── value
    ├── infra
    ├── module
    └── init.go
```
- `api`: Contains all the API related code, such as HTTP controllers, routes, middleware, gRPC handlers or any other API related code.
- `app`: Contains all the application related code, such as use cases, services, and commands.
- `domain`: Contains all the domain related code, such as entities, value objects, and domain services.
- `infra`: Contains all the infrastructure related code, such as database repositories, message queues, and third-party services.
- `module`: Contains the metadata of the module, such as its name.
- `init.go`: Contains the initialization logic of the module.

After creating a module and initializing it in the `init.go` file, you must register it in the [bootstrap](../internal/bootstrap.go) file. This file is responsible for initializing all the modules in the application.