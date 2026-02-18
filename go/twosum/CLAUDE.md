# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based LeetCode solutions. Each problem lives in its own directory under `go/`. This directory contains the Two Sum problem (#1).

## Build & Test Commands

```bash
# Initialize module (if go.mod doesn't exist yet)
go mod init twosum

# Run all tests
go test ./...

# Run a single test by name
go test -run TestTwoSum -v

# Build check (compile without producing binary)
go vet ./...
```

## Conventions

- Solution files: `solution.go` with package `twosum`
- Test files: `solution_test.go` using the standard `testing` package
- Problem descriptions are in markdown files (e.g., `TwoSum.md`)
