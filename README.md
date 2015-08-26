# parking backend

## Dependency

All the dependencies are vendored into the `vendor/` directory, following the Go 1.5 experimental vendoring system.

I've used the `godep` tool to automatically create the content of the `vendor/` directory.
It is needed to update/install/remove dependencies into this directory.
To setup `godep`:

```
go get github.com/tools/godep
```

Godeps stores its own metadata into the `Godeps` directory.

## Build and run

In the directory (after having `go get godep`):

``
GO15VENDOREXPERIMENT=1 godep restore
GO15VENDOREXPERIMENT=1 go build
./parking
``

## Configuration

To be docker-proof and to quickly run the binary, the runtime
configuration is directly read in the environment var. See runtime/config.go
for the parameters.

Example with default parameters:

```
GO15VENDOREXPERIMENT=1 go build
./parking 
```

Example with a parameter set:

```
GO15VENDOREXPERIMENT=1 go build
ADDR=localhost:9000 ./parking
```

