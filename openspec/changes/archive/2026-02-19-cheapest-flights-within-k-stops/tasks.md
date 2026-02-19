## 1. Project Setup

- [x] 1.1 Create directory `go/graph/cheapest-flights/`
- [x] 1.2 Initialize `go.mod` with module `cheapest_flights`
- [x] 1.3 Create `CheapestFlights.md` with problem description

## 2. Core Implementation

- [x] 2.1 Implement `findCheapestPrice` in `solution.go` using Bellman-Ford with k+1 iterations and dist array copying per iteration

## 3. Tests

- [x] 3.1 Create `solution_test.go` with table-driven tests using `t.Run` covering all spec scenarios (path within limit, cheaper one-stop, stop limit forcing direct, no route, no flights, single direct flight)

## 4. Documentation

- [x] 4.1 Create `CLAUDE.md` with build/test commands following existing conventions

## 5. Verify

- [x] 5.1 Run `go test ./...` and confirm all tests pass
- [x] 5.2 Run `go vet ./...` and confirm no issues
