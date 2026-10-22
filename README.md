# Programming & Backend Fundamentals Using Golang

A comprehensive, first-principles-based series designed to build deep engineering competence in Go (Golang), taking you from hardware-level memory management to scalable backend architectures.

## 📂 Repository Structure (Monorepo)

This repository is structured as a Go workspace containing modular packages for each video/topic in the series:

- `1-programming-fundamentals-using-golang/` - Introduction to Go syntax, variables, and core structures.
- `2-memory-management-and-pointers/` - Hardware memory, RAM, Stack vs. Heap, Garbage Collector, and Pointers in Go.
- _(Upcoming modules will be added here...)_

## 🚀 Getting Started

### Prerequisites

- Make sure you have [Go](https://golang.org/) installed.

### Cloning and Running

Clone the repository and use Go Workspaces to run any module:

```bash
git clone [https://github.com/your-username/programming-fundamentals-using-golang.git](https://github.com/your-username/programming-fundamentals-using-golang.git)
cd programming-fundamentals-using-golang
```

To run a specific module (e.g., Memory Management & Pointers):

```bash
go run ./2-memory-management-and-pointers/main.go
```

## 🛠️ Workspace Configuration (`go.work`)

This repository uses Go Workspaces (`go.work`) to manage multiple independent modules cleanly under a single Monorepo structure.

```go
go 1.27.2

use (
    ./1-programming-fundamentals-using-golang
    ./2-memory-management-and-pointers
)
```

## 📺 Watch the Series

This code accompanies the video series on YouTube by **Ahmed Ouda**.

- [YouTube Channel](https://www.youtube.com/@ahmed-ouda)

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
