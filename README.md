# Go-MongoDB
This is sample project using MongoDB with DDD architecture.

# Prerequisites

```
% dep ensure
```

# Run project with docker 
Builds the project as a docker image and then runs the service and MongoDB container  
```
% docker-compose up
```

# Run project with locally 
Runs the MongoDB container before to run the service with `go run`
```
% make run
```