## Why

專案需要更多排序演算法範例。Bubble Sort 是最基礎的排序演算法之一，適合作為排序類別的第一個範例 package。

## What Changes

- 新增 `bubble_sort/` 資料夾與同名 package，實作 bubble sort 演算法並提供 `Run()` function 作為進入點
- 在 `main.go` 的 switch 中新增 `"bubble_sort"` case，呼叫 `bubble_sort.Run()`

## Capabilities

### New Capabilities

- `bubble-sort`: Bubble sort 排序演算法範例 package，包含排序實作與執行展示

### Modified Capabilities

- `cli-entry`: 在 main.go switch 中新增 `"bubble_sort"` case 以支援新的範例 package

## Impact

- 新增檔案：`bubble_sort/bubble_sort.go`
- 修改檔案：`main.go`（新增 import 與 switch case）
