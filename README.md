# Go Fun Events

## Project structure
Go Fun Events is implemented by  a hexagonal architecture and this is its distribution:

```
go-fun-events/
|-- core/
|   |-- domain
|   |   |-- model.go
|   |-- service.go
|   |-- configuration.go
|
|-- ports/
|   |-- provider.go
|   |-- repository.go
|
|-- adapters/
|   |-- xxx-provider/
|   |   |-- model.go
|   |   |-- provider.go
|   |   |-- configuration.go
|   |
|   |-- mongodb-respository/
|       |-- model.go
|       |-- repository.go
|       |-- configuration.go
|
|-- go.mod
```
## Business logic description

When the method GetEvents is called, its retrieves the events from a repository where they are persisted. The repository must be fed frequently, so that, from time to time, the service calls to the provider to get updated events.  
The frequent which the provider is called is determined by the property timeToFeed. The provider is called when the time from the last time has passed, and the block of code responsible of the call is locked and the parallel threads skip it.

## How to run it
...
## Step by step
### 1- Repositories generation
This step must be executed just when code is change and a new generation is necessary
- 1- Installation
>`go install github.com/sunboyy/repogen@latest`

- 2- Generation
>`make code-generation`
### 2- Run tests
- 1- Mocks generation

Installation:`go install github.com/vektra/mockery/v2@v2.40.1`

Generation: `make clean mocks`
- 2- Test 

`make test`

### 3- Prepare environment, build and deploy the project
- 1- Configure the application.yml file correctly (conf/application.yml)
- 2- Execute `make build` to build the package
- 3- Execute `make deploy`to run the application with docker-compose

