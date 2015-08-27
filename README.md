# parking backend

## Dependency

All the dependencies are vendored into the `vendor/` directory, following the Go 1.5 experimental vendoring system.

I've used the `godep` tool to automatically create the content of the `vendor/` directory.
It is needed to update/install/remove dependencies into this directory.
To setup `godep`:

```
go get -u github.com/tools/godep
```

Godeps stores its own metadata into the `Godeps` directory.

## Build and run

First you must play the `schema.sql` file into your PostgreSQL.

In the directory (after having `go get godep`):

``
godep restore
go build
./parking

``

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

