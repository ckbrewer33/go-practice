# Go Practice

Personal practice monorepo for rebuilding Go fluency by implementing core data structures from scratch.

Curriculum: [Data Structures in Go – Practice Curriculum.md](./Data%20Structures%20in%20Go%20%E2%80%93%20Practice%20Curriculum.md)

## Layout

Each subdirectory is an independent Go module covering one topic from the curriculum:

| Directory | Topic |
|---|---|
| `dynamic-array/` | 1. Dynamic Array (ArrayList) |
| `singly-linked-list/` | 2. Singly Linked List |
| `doubly-linked-list/` | 3. Doubly Linked List |
| `stack-queue/` | 4. Stack and Queue |
| `binary-search-tree/` | 5. Binary Search Tree |
| `heap/` | 6. Heap (Priority Queue) |
| `hash-map/` | 7. Hash Map |
| `sorting/` | 8. Sorting Algorithms |
| `graph/` | 9. Graphs |
| `lru-cache/` | 10. LRU Cache (capstone) |

Each starts as a stub `main.go` with a hello-world. Build it out from there.

## Running

```sh
cd <module>
go run .
go test ./...
```

If you want a binary, build into `bin/` (gitignored):

```sh
go build -o bin/ .
```

## Workflow reminders

- Write the interface first, then implement incrementally
- Table-driven tests; benchmarks where relevant
- Document Big-O on each operation
- No autocomplete/AI for first pass — teach it out loud
