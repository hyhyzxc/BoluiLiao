# BoluiLiao

This project is structured as follows:

/server
│
├── cmd/                     # Main applications for this project.
│   └── server/              # Main starter of the web server.
│       └── main.go          # Entry point that starts the Gin server.
│
├── api/                     # Web API specific components: controllers, middleware, etc.
│   ├── handler/             # Request handlers (equivalent to controllers).
│   ├── middleware/          # API middleware components.
│   └── response/            # Response-related utilities or constructors.
│
├── config/                  # Configuration file parsers and config related things.
│   └── config.go
│
├── internal/                # Private application and library code.
│   ├── model/               # Domain models used by your application.
│   └── service/             # Business logic and service layer.
│
├── pkg/                     # Library code that's ok to use by external applications.
│   ├── db/                  # Database related utilities.
│   └── util/                # Utility functions that might be useful across the project.
│
├── deployments/             # Deployment configurations and scripts.
│   ├── Dockerfile
│   └── docker-compose.yml
│
├── tests/                   # Additional external test apps and test data.
│   └── api_test.go
│
├── go.mod                   # Go module dependencies.
└── go.sum                   # Sum file for module dependencies.


## Description of Directories

- **`cmd/`**: Contains the main applications for this project, including the starter of the web server.
- **`api/`**: Manages Web API-specific components like handlers, middleware, and response formatting.
- **`config/`**: Houses configuration file parsers and other configuration-related utilities.
- **`internal/`**: Holds private application code that defines business logic and application models.
- **`pkg/`**: Library code that can be used externally, including database utilities and various helper functions.
- **`deployments/`**: Contains deployment configurations and scripts including Docker setups.
- **`tests/`**: Provides additional external test applications and test data.

