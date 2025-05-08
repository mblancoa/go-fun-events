# Go Acme Events

## Project structure
Go Acme Events is implemented by  a hexagonal architecture and this is its distribution:

```
go-acme-events/
|-- core/
|   |-- domain
|   |   |-- model.go
|   |-- service.go
|   |-- configuration.go
|-- ports/
|   |-- provider.go
|   |-- repository.go
|   
|-- go.mod
```
## Logic business description

When the method GetEvents is called, its retrieves the events from a repository where they are persisted. The repository must be fed frequently, so that, from time to time, the service calls to the provider to get updated events.  
The frequent which the provider is called is determined by the property timeToFeed. The provider is called when the time from the last time has passed, and the block of code responsible of the call is locked and the parallel threads skip it.

## How to run it
- 1- Configure the application.yml file correctly /cofig/application.yml)
- 2- Execute `make build` to build the package
- 3- Execute `make deploy`to run the application with docker-compose

## Other considerations
