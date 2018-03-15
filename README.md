# myra-shell

A Golang implementation to use myracloud api via console.

## Requirements
- Dependencies via [glide](https://glide.sh/).
- Building a packed binary requires [UPX](https://upx.github.io/).
- Make

## Getting started

- Install dependencies
```
glide install
```
- Run test
```
make test
```

- Build
```
# Generates a static compiled packed executable
make

# Alternative creates a static compiled nonpacked executable
go build -ldflags="-s -w" -o myra-shell
```

## Status
Currently not all myracloud api's or options can be accessed via myra-shell.

Work in progress.
