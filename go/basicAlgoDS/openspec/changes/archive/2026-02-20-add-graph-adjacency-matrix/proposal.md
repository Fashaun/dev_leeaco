## Why

專案目前缺少圖（Graph）相關的資料結構範例。基於鄰接矩陣（Adjacency Matrix）的圖是最基礎的圖表示法之一，適合作為圖類別的第一個範例 package，涵蓋初始化、新增/刪除邊、新增/刪除頂點、印出圖等基本操作。

## What Changes

- 新增 `graph_adjacency_matrix/` 資料夾與同名 package，實作基於鄰接矩陣的無向圖，提供基本操作（初始化、新增/刪除邊、新增/刪除頂點、印出圖）
- 提供 `Run()` function 展示各項操作
- 在 `main.go` 的 switch 中新增 `"graph_adjacency_matrix"` case

## Capabilities

### New Capabilities

- `graph-adjacency-matrix`: 基於鄰接矩陣的圖資料結構，包含初始化、新增/刪除邊、新增/刪除頂點、印出圖等基本操作

### Modified Capabilities

- `cli-entry`: 在 main.go switch 中新增 `"graph_adjacency_matrix"` case 以支援新的範例 package

## Impact

- 新增檔案：`graph_adjacency_matrix/graph_adjacency_matrix.go`
- 修改檔案：`main.go`（新增 import 與 switch case）
