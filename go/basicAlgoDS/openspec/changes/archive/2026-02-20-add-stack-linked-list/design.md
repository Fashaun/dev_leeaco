## Context

專案已有多個資料結構範例（圖、排序等），但缺少堆疊。使用鏈結串列實作堆疊是經典做法，以串列頭部作為堆疊頂部，Push/Pop 操作均為 O(1)。

## Goals / Non-Goals

**Goals:**

- 定義 `Node` struct 封裝節點（包含 `Val int` 和 `Next *Node`）
- 定義 `LinkedListStack` struct，以鏈結串列頭部作為堆疊頂部
- 提供 Push、Pop、Peek、Size、IsEmpty、Print 操作
- 提供 `Run()` 展示所有操作

**Non-Goals:**

- 不實作基於陣列/slice 的堆疊（可另開 change）
- 不實作泛型版本（使用 `int` 保持簡單）
- 不實作執行緒安全版本

## Decisions

### 1. 使用鏈結串列頭部作為堆疊頂部

Push 時在頭部插入新節點，Pop 時移除頭部節點。

**理由**: 頭部操作為 O(1)，不需要遍歷串列。若使用尾部則需要遍歷到倒數第二個節點才能 Pop，時間複雜度為 O(n)。

### 2. 使用 `Node` struct 而非直接使用 Go 標準庫的 `container/list`

自行定義 `Node` struct。

**理由**: 教學目的，展示鏈結串列節點的指標操作。使用標準庫會隱藏實作細節。

### 3. Pop 和 Peek 回傳 `(int, bool)`

回傳值與是否成功的 bool，當堆疊為空時回傳 `(0, false)`。

**理由**: 比 panic 更安全，讓呼叫者可以檢查操作是否成功。符合 Go 慣用的 comma-ok 模式。

**替代方案**: 空堆疊時 panic — 不安全，不適合教學範例。

## Risks / Trade-offs

- **每個節點額外佔用指標空間** → 可接受，這是鏈結串列的特性，適合教學展示與陣列版本的對比
