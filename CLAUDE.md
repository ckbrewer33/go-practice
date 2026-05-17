# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

Personal practice monorepo for rebuilding Go fluency by reimplementing core data structures from scratch. Each topic in `Data Structures – Practice Curriculum.md` has its own subdirectory and is an **independent Go module** (its own `go.mod`, `package main`). There is no root module and no workspace file — modules do not import each other.

## Working in a module

All commands run from inside the module directory, not the repo root:

```sh
cd <module>          # e.g. cd dynamic-array
go run .             # run main.go
go test ./...        # run all tests in the module
go test -run TestDynamicArray   # run a single test by name
go test -bench .     # run benchmarks (where present)
go build -o bin/ .   # build into module-local bin/ (gitignored)
```

Tests live alongside `main.go` in the same `package main` (e.g. `dynamic-array/dynamic-array-test.go`), so they can access unexported types directly.

## How the user works here

This is a learning exercise, not a product. A few things follow from that:

- **Stubs are intentional.** Methods often start as signatures with a comment describing the algorithm and a `fmt.Println("X not implemented yet")` body. Do not "helpfully" fill them in — the user is implementing them by hand. The comment above each method is the spec they're working from.
- **No autocomplete/AI for the first pass.** Per README, the user wants to teach the implementation out loud before getting help. When asked for help, prefer Socratic hints (invariants, edge cases, what the comment says to do) over writing the code for them.
- **Curriculum-driven.** When a question references a data structure, the authoritative description (invariants, methods to build, pitfalls, test scenarios, complexity targets) is in `Data Structures – Practice Curriculum.md` at the repo root. Read the relevant section before suggesting designs.
- **Per-operation Big-O matters.** The curriculum expects each method to be annotated with its time/space complexity; treat that as part of "done."
- **Directly reference official Go documentation** The user will ask questions about Go syntax and idioms. All answers should be referenced from the official Go documentation located at https://go.dev/doc/. Use the context7 MCP tool as necessary to refernece documentation.
