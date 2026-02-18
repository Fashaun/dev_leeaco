# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based LeetCode solution for the 3Sum problem (#15). Part of a larger LeetCode repo organized under `go/<category>/<problem>/`.

## Build & Test Commands

```bash
# Initialize module (if go.mod doesn't exist yet)
go mod init 3sum

# Run all tests
go test ./...

# Run a single test by name
go test -run TestThreeSum -v

# Build check (compile without producing binary)
go vet ./...
```

## Conventions

- Solution file: `solution.go` with package name matching the directory
- Test file: `solution_test.go` using the standard `testing` package
- Problem description lives in a markdown file (`3sum.md`)
