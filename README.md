# Golang API Template

This template serves a backend API using Onion Architecture to allow testability and maintainability by incorporating single-responsibility layers that are loosely coupled.

Following the interface segregation and dependency inversion principles of SOLID, all objects depend on small and specific interfaces rather than large concrete implementations. Using interfaces instead of concrete implementations allows us to perform unit testing of single objects without worrying about the implementations that they require and enables future flexibility when making large changes to the application.

## Folder Structure

	.
	├── build                   # Compiled files
	├── cmd                     # Entry point to application
	├── internal                # Private package only used by `./cmd/api.go`
	│   ├── api                 # HTTP presentation layer `entity_handler.go`
	│   │   └── config          # Configuration for entry point to serve api. Dependency injection and http server setup.
	│   ├── application         # Business models, logic, interfaces `entity_service.go`
	│   ├── domain              # Enterprise models and logic `entity.go`
	│   └── infrastructure      # Data access concerns (MySQL, Mongo, S3, Twilio, etc.)
	|   pkg			    # Shared logic (Client SDK, Logger)
	├── tests                   # Automated tests
    │   ├── api                 # Presentation logic unit tests (Validation)
	│   ├── application         # Business logic unit tests
	│   ├── domain              # Enterprise logic unit tests
	│   └── infrastructure      # Integration tests
	└── README.md

## Layers / Dependency Flow

As we see in the folder structure, the application is separated into a few different layers each with its own concern. To further organize our code, we follow one-way dependencies by only importing from an inward layer of the dependency diagram below. Layers in the center of the diagram, defined as the "core" of the application logic, are more abstract and applicable across the organization (domain) or application (app), then become less abstract towards the outside of the diagram with data access implementations (infra) and presentation logic (api).

![Dependency Diagram](https://jasontaylor.dev/wp-content/uploads/2020/01/Figure-01-2.png)

**API**: Responsible for defining how the server responds to API requests via handlers, configuring the API with any third-party integrations, and injecting all application dependencies into the application API handlers.
- Flow: API depends on \[Application, Infrastructure\]
  - Why? Since the API layer configures our API server, it relies on business logic from the Application layer and requires implementations (defined in Infrastructure layer) for the interfaces that the Application layer defines and uses for data access.

**Application**: Responsible for defining interfaces for its services and its data access dependencies (through an external source like S3 or internal source like our DBs). Also defines business logic models and business logic implementations for use in the API handlers containing validation, mapping, calculation, or other business logic.
- Flow: Application depends on \[Domain\]
  - Why? Since the Application layer is responsible for business logic and cannot rely on outer layers, it depends only on the Domain layer for enterprise business logic where applicable. Most of the time this will be in the form of defining Domain model return types for the interfaces that have implementations in the Infrastructure layer.

**Domain**: Responsible for defining domain models and associated enterprise logic if applicable.
- Flow: Domain depends on \[\]
  - Why? Since the Domain layer is the inner most layer, it depends on no layers. It defines the core domain models that get used in the data access layers (Infrastructure)

**Infrastructure**: Responsible for defining implementations of third-party integrations (S3, Twilio, etc.)
- Flow: Infrastructure depends on \[Application, Domain\]
	- Why? Since the Infrastructure layer defines implementations for external data access, it relies on the Application layer for interface definitions and the Domain layer for the necessary data types that are stored and retrieved from external data sources.

_Note on Infrastructure: Because golang interfaces are implicit, you won't see imports from the Application layer, but it still depends on it._
