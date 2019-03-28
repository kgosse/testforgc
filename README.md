# testforgc

Test Exercice

## Getting Started

I've done the test using golang.

### Prerequisites

Make sure you have the go binary install on your machine. Here is the download page: https://golang.org/dl/

If you have access to the go command in your terminal, everything is perfect. You can take the version mentionned or above.

```
$ go version
go version go1.11.5 darwin/amd64
```

### Installing

```
$ git clone git@github.com:kgosse/testforgc.git
$ cd testforgc
$ go build #it will install the dependencies and create an executable
$ ./testforgc --help
$ ./testforgc -f your-input-file.list
$ ./testforgc
```

## Running the (unit) tests

```
$ go test ./...
```