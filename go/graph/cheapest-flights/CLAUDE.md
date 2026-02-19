# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based LeetCode solution for the Cheapest Flights Within K Stops problem (#787). Part of a larger LeetCode repo organized under `go/<category>/<problem>/`. This is a graph/constrained shortest-path problem using the Bellman-Ford algorithm.

## Build & Test Commands

```bash
# Initialize module (if go.mod doesn't exist yet)
go mod init cheapest_flights

# Run all tests
go test ./...

# Run a single test by name
go test -run TestFindCheapestPrice -v

# Build check (compile without producing binary)
go vet ./...
```

## Conventions

- Solution file: `solution.go` with package name `cheapest_flights`
- Test file: `solution_test.go` using the standard `testing` package with table-driven tests
- Problem description lives in a markdown file (`CheapestFlights.md`)
