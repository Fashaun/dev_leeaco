## Why

專案已有基於鄰接矩陣的圖範例，但缺少鄰接表（Adjacency List）表示法。鄰接表是另一種常見的圖表示方式，在稀疏圖中比鄰接矩陣更節省空間，適合作為圖類別的第二個範例 package，與鄰接矩陣形成對照。

## What Changes

- 新增 `graph_adjacency_list/` 資料夾與同名 package，實作基於鄰接表的無向圖，提供基本操作（初始化、新增/刪除邊、新增/刪除頂點、印出圖）
- 提供 `Run()` function 展示各項操作
- 在 `main.go` 的 switch 中新增 `"graph_adjacency_list"` case

## Capabilities

### New Capabilities

- `graph-adjacency-list`: 基於鄰接表的圖資料結構，包含初始化、新增/刪除邊、新增/刪除頂點、印出圖等基本操作

### Modified Capabilities

- `cli-entry`: 在 main.go switch 中新增 `"graph_adjacency_list"` case 以支援新的範例 package

## Impact

- 新增檔案：`graph_adjacency_list/graph_adjacency_list.go`
- 修改檔案：`main.go`（新增 import 與 switch case）
