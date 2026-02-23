## Why

專案目前缺少程式進入點 (`main.go`) 與 Go module 初始化。需要建立 main.go 作為統一入口，透過 CLI 第一個參數動態選擇要執行的範例 package，同時建立第一個範例 package `go_algods_example` 作為樣板。

## What Changes

- 初始化 Go module (`go.mod`)
- 新增 `main.go`，讀取 `os.Args[1]` 作為資料夾/package 名稱，以 switch/map 方式呼叫對應 package 的 function
- 新增 `go_algods_example/` 資料夾與同名 package，內含一個 exported function，印出 `"Hello go_algods_example"`

## Capabilities

### New Capabilities

- `cli-entry`: main.go 程式進入點，透過第一個 CLI 參數分派到對應的範例 package
- `example-package`: 第一個範例 package `go_algods_example`，提供可被 main.go 呼叫的 function

### Modified Capabilities

(None — this is a greenfield setup)

## Impact

- 新增檔案：`go.mod`、`main.go`、`go_algods_example/go_algods_example.go`
- 未來新增範例 package 時需在 main.go 的分派邏輯中加入對應 case
