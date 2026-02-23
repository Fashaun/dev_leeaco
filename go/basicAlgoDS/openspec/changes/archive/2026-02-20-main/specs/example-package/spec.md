## ADDED Requirements

### Requirement: Example package provides Run function

The `go_algods_example` package SHALL export a `Run()` function that can be called by `main.go`.

#### Scenario: Run function is called

- **WHEN** `go_algods_example.Run()` is invoked
- **THEN** the function SHALL print `"Hello go_algods_example"` to stdout

### Requirement: Package naming convention

The package name SHALL match the directory name (`go_algods_example`), following the project convention described in README.md.

#### Scenario: Package is importable

- **WHEN** `main.go` imports `basicAlgoDS/go_algods_example`
- **THEN** the package SHALL compile and be accessible as `go_algods_example`
