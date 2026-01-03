# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A collection of cryptographic and number theory arithmetic algorithms implemented in Go. Each algorithm lives in its own isolated Go module.

## Build Commands

```bash
# Run all tests across all modules
make test

# Run linting across all modules (requires golangci-lint via aqua)
make lint

# Format code across all modules
make format

# Run tests in a specific module
cd montgomery && go test -v ./...

# Run a single test
cd rabin && go test -v -run TestIsPrime ./...

# Run benchmarks
cd montgomery && go test -bench=. ./...
```

## Architecture

**Multi-Module Structure**: Each algorithm is a separate Go module with its own `go.mod`, allowing independent versioning. There is no root go.mod, so `go` commands must be run inside each module directory (use `make` targets for cross-module operations):

- `montgomery/` - Montgomery multiplication (three implementations: Bitwise, CIOS, CIOSWords)
- `pollard/` - Pollard's rho algorithm for integer factorization using Floyd's cycle detection
- `rabin/` - Miller-Rabin probabilistic primality test

## Testing

- Table-driven tests with named cases
- Benchmark tests using Go 1.24+ `b.Loop()` pattern
- `t.Parallel()` for concurrent test execution where applicable
- Test parameters include large numbers (2048-bit) for cryptographic relevance
