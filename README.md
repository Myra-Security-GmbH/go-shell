# myra-shell

A Golang implementation to use Myra's API via console.

## Requirements
- make
- go

## Getting started

- Build
```
# Generates a static compiled packed executable
make

# Alternative creates a static compiled nonpacked executable
make build
```

- Run test
```
make test
```

- Run
```
./myra-shell
```
Running myra-shell for the first time, it will ask for your Myra API credentials.

## Status
Currently not all Myra API's or options can be accessed via myra-shell.
