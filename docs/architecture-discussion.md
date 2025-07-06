# 🏗️ LevelUpOps – Architecture Overview

This document outlines the core architectural decisions, abstractions, and structural layout for the LevelUpOps CLI tool and its ecosystem.

---

## 📁 Go Project Structure

```
/cmd/
  levelupops/         → main.go (entrypoint with CLI commands)

/internal/
  quests/             → quest parsing, XP logic, metadata
  sandbox/            → pluggable sandbox engine (file/container/vm)
  progress/           → XP tracking, journal logging
  logger/             → Logrus-based log wrapper
  telemetry/          → opt-in analytics + event logs
  config/             → YAML/JSON config parsing

/pkg/
  tui/                → BubbleTea/Tview UI components
  graphql/            → gqlgen GraphQL API (future)
  utils/              → shared CLI helpers and I/O

/scripts/
  bootstrap.sh        → local dev bootstrap
  justfile            → reproducible automation tasks

/.github/workflows/
  oci-deploy.yml      → GitHub Actions for OCI setup

/quests/
  intro-bash/
    quest.yaml        → Metadata and logic
    script.sh         → Executable script
```

---

## 📄 YAML Quest Schema

Each quest is defined in a `quest.yaml` file.

```yaml
id: intro-bash
title: "Bash Basics: Hello CLI"
description: |
  Learn basic bash scripting, redirection, and permissions.
xp: 30
difficulty: 1
tags: ["bash", "shell", "beginner"]
dependencies: []
commands:
  - run: "bash ./script.sh"
    expect: "Hello DevOps"
verify:
  - file_exists: "/tmp/devops.txt"
  - grep: "success" in "/tmp/devops.txt"
```

---

## 🗃 Database Decision

| Feature        | SQLite          | BoltDB       |
|----------------|------------------|--------------|
| SQL Support    | ✅ Yes           | ❌ No        |
| Sync Options   | ✅ With LiteFS   | ❌ No        |
| Maturity       | ✅ High          | ✅ Stable    |
| Use Case Fit   | ✅ Structured XP | ❌ Limited   |

**Choice**: SQLite with LiteFS support for backup and sync.

---

## 🧾 CLI Command Structure (Cobra)

| Command                  | Description                         |
|--------------------------|-------------------------------------|
| `init`                   | Initialize config and XP store      |
| `quests browse`          | List available quests               |
| `start <id>`             | Start a specific quest              |
| `progress`               | Show XP progress and journal        |
| `cleanup`                | Remove stale sandboxes/resources    |
| `doctor`                 | Run diagnostics                     |

Example flags:

```bash
levelupops start intro-bash --sandbox=file --verbose
levelupops cleanup --dry-run
```

---

## 🧱 Sandbox Engine Interface

```go
type Sandbox interface {
  Setup(quest Quest) error
  Execute() error
  Teardown() error
  Name() string
}
```

### Implementations

- `FileSandbox`       → Offline folder-based execution
- `ContainerSandbox`  → Podman or Docker (future)
- `VMSandbox`         → WSL2 / Multipass / Lima (future)

### Factory Method

```go
func NewSandbox(sandboxType string) Sandbox {
  switch sandboxType {
    case "file": return NewFileSandbox()
    case "container": return NewContainerSandbox()
    case "vm": return NewVMSandbox()
  }
}
```

---

## 🔭 Telemetry & Logging

- Structured logs using **Logrus**
- TTL-cleanup logs and journaling
- Optional telemetry (GDPR compliant)
- GitHub or S3 XP backup

---

## 🎯 Design Goals

- ✅ CLI-first for offline learners
- ✅ Modular and testable components
- ✅ Easy to extend sandbox types
- ✅ Local-first, privacy-respecting
- ✅ Future-ready for GraphQL and TUI
