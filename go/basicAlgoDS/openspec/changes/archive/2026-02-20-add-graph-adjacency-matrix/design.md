## Context

專案已有多個演算法範例 package（`go_algods_example`、`btree_traversal`、`bubble_sort`），各自提供 `Run()` function。現在要新增基於鄰接矩陣的圖資料結構範例，涵蓋圖的基本操作。

## Goals / Non-Goals

**Goals:**

- 實作 `GraphAdjMat` struct，使用二維 slice `[][]int` 作為鄰接矩陣，`[]int` 儲存頂點列表
- 提供基本操作：初始化、新增/刪除邊、新增/刪除頂點、印出圖
- 實作無向圖（矩陣對稱）
- 提供 `Run()` 展示所有操作

**Non-Goals:**

- 不實作有向圖或加權圖（可作為後續 change）
- 不實作圖的走訪演算法（BFS/DFS）
- 不實作基於鄰接表的圖（另一種表示法，可另開 change）

## Decisions

### 1. 使用 struct 封裝圖資料結構

定義 `GraphAdjMat` struct 包含 `vertices []int` 和 `adjMat [][]int`，透過 method 操作。

**理由**: 使用 struct + method 符合 Go 慣例，封裝性好，也方便展示 Go 的 receiver 語法。

**替代方案**: 使用純 function + 全域變數 — 不符合 Go 慣例且不易擴展。

### 2. 使用 `NewGraphAdjMat(vertices []int, edges [][]int)` 建構函式

提供 constructor function 接受初始頂點與邊的列表，回傳 `*GraphAdjMat`。

**理由**: 與 `btree_traversal.NewTreeNode` 模式一致，讓使用者一次建立完整的初始圖。

### 3. 鄰接矩陣使用 0/1 表示邊的存在

矩陣值為 `0`（無邊）或 `1`（有邊），不使用權重值。

**理由**: 無向無權圖是最基礎的圖模型，適合教學。

### 4. 頂點以 int 值表示，位置索引對應矩陣行列

`vertices[i]` 對應矩陣第 `i` 行/列。新增頂點時擴展矩陣，刪除頂點時縮減矩陣並移除對應行列。

**理由**: 簡單直觀，適合展示鄰接矩陣的結構特性。

## Risks / Trade-offs

- **刪除頂點需要搬移矩陣行列，時間複雜度 O(n²)** → 可接受，這正是鄰接矩陣的特性，適合教學展示
- **頂點值重複時可能造成混淆** → 在教學範例中使用不重複的值即可，不需額外檢查
