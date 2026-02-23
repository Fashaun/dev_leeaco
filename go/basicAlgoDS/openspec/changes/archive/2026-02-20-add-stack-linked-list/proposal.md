## Why

專案目前缺少堆疊（Stack）資料結構範例。使用鏈結串列（Linked List）實作堆疊是最基礎的方式之一，適合展示 LIFO（Last In, First Out）特性以及鏈結串列的節點操作。

## What Changes

- 新增 `stack_linked_list/` 資料夾與同名 package，使用鏈結串列實作堆疊，提供基本操作（Push、Pop、Peek、Size、IsEmpty、Print）
- 提供 `Run()` function 展示各項操作
- 在 `main.go` 的 switch 中新增 `"stack_linked_list"` case

## Capabilities

### New Capabilities

- `stack-linked-list`: 基於鏈結串列的堆疊資料結構，包含 Push、Pop、Peek、Size、IsEmpty、Print 等基本操作

### Modified Capabilities

- `cli-entry`: 在 main.go switch 中新增 `"stack_linked_list"` case 以支援新的範例 package

## Impact

- 新增檔案：`stack_linked_list/stack_linked_list.go`
- 修改檔案：`main.go`（新增 import 與 switch case）
