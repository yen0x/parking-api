# parking backend

## Configuration

To be docker-proof and to quickly run the binary, the runtime
configuration is directly read in the environment var. See runtime/config.go
for the parameters.

Example with default parameters:

```
go build
./parking 
```

Example with a parameter set:

```
go build
ADDR=localhost:9000 ./parking
```


