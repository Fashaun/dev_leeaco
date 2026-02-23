## Context

專案 `basicAlgoDS` 是一個 Go 演算法與資料結構範例集合，目前尚無任何 Go 原始碼。需要從零建立 Go module、程式進入點與第一個範例 package。

## Goals / Non-Goals

**Goals:**

- 建立可執行的 Go 程式進入點 (`main.go`)
- 透過 CLI 第一個參數選擇要執行的範例 package
- 建立 `go_algods_example` package 作為範例樣板
- 結構簡單易懂，方便未來新增更多範例 package

**Non-Goals:**

- 不建立自動發現 package 的機制（保持顯式註冊）
- 不處理子命令、flag parsing 等複雜 CLI 功能
- 不為既有空目錄 `hashtable` 建立程式碼

## Decisions

### 1. 使用 `switch` 分派範例 package

使用 `switch` statement 根據 `os.Args[1]` 分派到對應 package 的 function。

**理由**: 相比 map + function reference，switch 更直觀、可讀性高，且適合範例數量不多的場景。每新增一個範例只需加一個 case。

**替代方案**: 使用 `map[string]func()` 做分派 — 更彈性但對於教學用途的專案來說過度設計。

### 2. 每個範例 package 提供一個 `Run()` function

統一以 exported `Run()` 作為每個範例 package 的進入點，由 main.go 呼叫。

**理由**: 統一介面讓 main.go 的分派邏輯保持一致，新增範例時只需 import package 並在 switch 加一個 case 呼叫 `Run()`。

### 3. Go module 名稱使用 `basicAlgoDS`

**理由**: 與專案目錄名稱一致，簡潔明瞭。

## Risks / Trade-offs

- **每次新增範例需手動更新 main.go** → 可接受，因為這是教學專案，顯式註冊比自動發現更清晰
- **未提供參數時程式會 panic（index out of range）** → 在 main.go 中檢查 `len(os.Args)` 並印出使用提示
