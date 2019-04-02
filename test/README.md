# Testing Fetch

## Setup
If not already installed, you'll need Ginkgo and Gomega:
```sh
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
```

## Running all tests
Fetch's testing harness uses [Ginkgo](https://github.com/onsi/ginkgo) to test Fetch, simply move to this directory (/test) and run:
```sh
go test ./...
```
or (you will need to install ginkgo as an CLI executable):
```sh
ginkgo -r
```

## Running a single test
Similar to above, to run a single test instead run:
```sh
go test ./<path to module test suite> (i.e. go test ./core)
```
or
```sh
ginkgo ./<path to module test suite> (i.e. go test ./core)
```

## Running with verbose output
With ginkgo you can append ```-v``` to the command to get verbose output, this will describe every test that is running.

## Benchmarking
Benchmarks can be run using the following command
```sh
go test ./<PATH TO TEST> -bench=<REGEX OF BENCHMARK TO RUN>
```
For e.g.
```sh
go test ./core -bench=Queue
```

### To run all tests with benchmarks
```sh
go test ./... -bench=.
```

### To only run all benchmarks
```sh
go test ./... -run=XYZ -bench=.
```
(XYZ should be entered literally since XYZ doesn't match any tests, and therefore won't run any)

### Benchmarking for N seconds
```sh
go test ./... -bench=. -benchtime=<N>s
```
For e.g.
```sh
go test ./core -bench=Queue -benchtime=30s
```