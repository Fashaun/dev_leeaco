## ADDED Requirements

### Requirement: CLI argument dispatches to example package

The system SHALL read `os.Args[1]` as the package/folder name and invoke the corresponding example package's `Run()` function.

#### Scenario: Valid package name provided

- **WHEN** user runs `go run main.go go_algods_example`
- **THEN** the system SHALL call `go_algods_example.Run()` and display the package's output

#### Scenario: Unknown package name provided

- **WHEN** user provides a package name that does not match any registered example
- **THEN** the system SHALL print an error message indicating the package is not found

### Requirement: Usage hint when no argument provided

The system SHALL check `len(os.Args)` and print a usage hint if no argument is provided, instead of panicking.

#### Scenario: No arguments provided

- **WHEN** user runs `go run main.go` without any arguments
- **THEN** the system SHALL print a usage message (e.g., `Usage: go run main.go <example_name>`) and exit gracefully
