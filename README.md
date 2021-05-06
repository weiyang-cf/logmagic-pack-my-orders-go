# Getting Started

## Prerequisites

Install Go. This project is working on:

```bash
go version go1.15.5
```

But any version later than this should work just fine.

For more information, visit the [official Go website](https://golang.org/).

## Run the test cases

Clone the repository.

From the `orderHandler` directory, run `go test -v`.

## Implementation

Please provide your own implementation for the `PackOrder` method in `orderHandler.go`. Please take note:

- You should not modify the `PackOrder` method signature.
- You should not import 3rd party packages that implement the critical business logic (importing utility packages are fine).
- It is probably a good idea to apply good software engineering principles and code cleanly.
- It is probably a good idea to include unit tests.
