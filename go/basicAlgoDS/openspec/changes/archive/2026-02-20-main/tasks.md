## 1. Project Setup

- [x] 1.1 Initialize Go module with `go mod init basicAlgoDS`

## 2. Example Package

- [x] 2.1 Create `go_algods_example/` directory
- [x] 2.2 Create `go_algods_example/go_algods_example.go` with package `go_algods_example` and exported `Run()` function that prints `"Hello go_algods_example"`

## 3. CLI Entry Point

- [x] 3.1 Create `main.go` with `package main` and `func main()`
- [x] 3.2 Add argument check: if `len(os.Args) < 2`, print usage message and return
- [x] 3.3 Add `switch os.Args[1]` with case `"go_algods_example"` calling `go_algods_example.Run()`
- [x] 3.4 Add `default` case that prints unknown package error message

## 4. Verification

- [x] 4.1 Run `go run main.go go_algods_example` and verify it prints `"Hello go_algods_example"`
- [x] 4.2 Run `go run main.go` with no args and verify usage message is printed
- [x] 4.3 Run `go run main.go unknown` and verify error message is printed
