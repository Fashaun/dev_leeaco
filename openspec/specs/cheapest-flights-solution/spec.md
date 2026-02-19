### Requirement: Find cheapest price with at most k stops

The function `findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int` SHALL return the cheapest price from `src` to `dst` using at most `k` intermediate stops. If no such route exists, it SHALL return `-1`.

Each element in `flights` is `[from, to, price]` representing a directed edge with cost.

#### Scenario: Path exists within stop limit

- **WHEN** `n=4, flights=[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src=0, dst=3, k=1`
- **THEN** the function returns `700`
- **AND** the path uses at most 1 intermediate stop (0 -> 1 -> 3)

#### Scenario: Cheaper path available through direct and one-stop routes

- **WHEN** `n=3, flights=[[0,1,100],[1,2,100],[0,2,500]], src=0, dst=2, k=1`
- **THEN** the function returns `200`
- **AND** the path 0 -> 1 -> 2 is chosen over the direct flight costing 500

#### Scenario: Stop limit forces more expensive direct route

- **WHEN** `n=3, flights=[[0,1,100],[1,2,100],[0,2,500]], src=0, dst=2, k=0`
- **THEN** the function returns `500`
- **AND** the cheaper path 0 -> 1 -> 2 is excluded because it requires 1 stop

#### Scenario: No route exists to destination

- **WHEN** there is no path from `src` to `dst` within `k` stops
- **THEN** the function returns `-1`

### Requirement: Handle constraint boundaries

The function SHALL handle all inputs within the given constraints: `2 <= n <= 100`, `0 <= flights.length <= n*(n-1)/2`, `1 <= price <= 10^4`, `0 <= src, dst, k < n`, `src != dst`.

#### Scenario: No flights exist

- **WHEN** `n=3, flights=[], src=0, dst=2, k=1`
- **THEN** the function returns `-1`

#### Scenario: Single direct flight

- **WHEN** `n=2, flights=[[0,1,100]], src=0, dst=1, k=0`
- **THEN** the function returns `100`
- **AND** no intermediate stops are needed

### Requirement: Implementation in Go

The solution SHALL be implemented in Go, following the existing repository conventions: package-per-problem directory under `go/graph/cheapest-flights/`, with `solution.go`, `solution_test.go`, `go.mod`, a problem description markdown file, and `CLAUDE.md`.

#### Scenario: Package and module structure

- **WHEN** the solution is created
- **THEN** it resides in package `cheapest_flights` under `go/graph/cheapest-flights/`
- **AND** `go.mod` declares the module as `cheapest_flights`

#### Scenario: Table-driven tests

- **WHEN** tests are run with `go test ./...`
- **THEN** all scenarios from the spec pass as table-driven test cases using `t.Run` with named subtests
