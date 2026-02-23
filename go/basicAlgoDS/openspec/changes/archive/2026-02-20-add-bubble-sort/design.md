## Context

專案已有 `go_algods_example` 和 `btree_traversal` 兩個範例 package，各自提供 `Run()` function 由 `main.go` 透過 switch 分派。現在要新增第一個排序演算法範例 — bubble sort。

## Goals / Non-Goals

**Goals:**

- 實作標準的 bubble sort 演算法，對 `[]int` 進行原地排序
- 提供 `Run()` function 展示排序前後的結果
- 遵循既有 package 慣例，在 main.go 中註冊新 case

**Non-Goals:**

- 不實作泛型版本（保持範例簡單，使用 `[]int`）
- 不加入效能基準測試
- 不實作其他排序演算法（後續可另開 change）

## Decisions

### 1. 函式簽名使用 `BubbleSort(arr []int)`

提供 exported `BubbleSort` function 接受 `[]int` 並原地排序（in-place），不回傳新 slice。

**理由**: Bubble sort 本身就是 in-place 演算法，原地修改更貼近演算法本質。exported function 讓其他 package 也能使用。

**替代方案**: 回傳排序後的新 slice — 但會增加不必要的記憶體配置，且不符合 bubble sort 的教學意圖。

### 2. `Run()` 內建範例資料

`Run()` function 內自行定義一組範例陣列，印出排序前後的結果，不需外部輸入。

**理由**: 與既有 `go_algods_example` 模式一致，執行即可看到效果。

## Risks / Trade-offs

- **範例資料固定** → 可接受，教學用途不需動態輸入
