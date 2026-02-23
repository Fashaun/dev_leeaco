## Context

專案已有基於鄰接矩陣的圖範例 (`graph_adjacency_matrix`)，使用二維 slice 儲存邊。現在要新增基於鄰接表的圖，使用 `map` 儲存每個頂點的鄰居列表，與鄰接矩陣形成對照。

## Goals / Non-Goals

**Goals:**

- 實作 `GraphAdjList` struct，使用 `map[Vertex][]Vertex` 作為鄰接表
- 定義 `Vertex` struct 封裝頂點值
- 提供基本操作：初始化、新增/刪除邊、新增/刪除頂點、印出圖
- 提供 `Run()` 展示所有操作

**Non-Goals:**

- 不實作有向圖或加權圖
- 不實作圖走訪演算法（BFS/DFS）
- 不與 `graph_adjacency_matrix` 共用介面或抽象

## Decisions

### 1. 使用 `Vertex` struct 而非純 int

定義 `Vertex` struct 包含 `Val int`，作為 map 的 key。

**理由**: 使用 struct 更貼近 Hello Algo 教材的範例風格，且與 `btree_traversal` 的 `TreeNode` 模式一致。`Vertex` 作為指標可作為 map key，語義更清晰。

**替代方案**: 使用 `map[int][]int` — 更簡單但失去物件導向的教學意義。

### 2. 鄰接表使用 `map[*Vertex][]*Vertex`

使用指標作為 map key，每個頂點對應一個鄰居 slice。

**理由**: 指標作為 key 保證唯一性，即使兩個頂點的 Val 相同也不會衝突。slice 作為 value 方便動態新增/刪除鄰居。

### 3. 提供 `NewVertex(val int)` 建構函式

**理由**: 與專案中 `NewTreeNode`、`NewGraphAdjMat` 的命名慣例一致。

### 4. 使用 `NewGraphAdjList(edges [][]*Vertex)` 建構函式

接受邊的列表，每條邊為 `[]*Vertex{v1, v2}`，自動建立頂點和鄰接關係。

**理由**: 從邊列表建構圖是常見模式，且可同時註冊頂點和邊。

## Risks / Trade-offs

- **刪除頂點需遍歷所有鄰居列表移除引用，時間複雜度 O(V+E)** → 可接受，這是鄰接表的特性
- **使用指標作為 map key，印出時順序不固定** → 教學範例中可接受，重點在展示操作邏輯
