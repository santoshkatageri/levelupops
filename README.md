# 🚀 LevelUpOps – Gamified DevOps Learning Platform

LevelUpOps is a CLI-first, cross-platform DevOps learning tool built in Go. Learn infrastructure hands-on with quests, earn XP, and sync progress via GitHub or the cloud to grow your portfolio.


## Current Status (July 2025)


### Implemented
- Project structure scaffolded as per architecture
- Go modules and dependencies set up (Cobra, Logrus, SQLite, Testify)
- CLI entrypoint (`main.go`) and root command (`levelupops`)
- Core CLI commands implemented and tested:
  - `init`: initializes config and XP store
  - `quests browse`: lists available quests
  - `start <id>`: starts a specific quest
  - `progress`: shows XP progress and journal (placeholder)
  - `cleanup`: removes stale sandboxes/resources (placeholder)
  - `doctor`: runs diagnostics (placeholder)
- Test-driven development (TDD) in place: tests for all core commands
- GitHub Actions CI workflow: runs all tests on push/PR


### Verified
- All core CLI commands pass their tests locally and in CI
- Example verifications:
  - `go run ./cmd/levelupops/main.go init` prints `LevelUpOps initialized!`
  - `go run ./cmd/levelupops/main.go quests browse` lists quest IDs
  - `go run ./cmd/levelupops/main.go start <id>` prints quest started or not found
  - `go run ./cmd/levelupops/main.go progress` prints placeholder XP/journal
  - `go run ./cmd/levelupops/main.go cleanup` prints placeholder cleanup message
  - `go run ./cmd/levelupops/main.go doctor` prints placeholder diagnostics

### Next Steps
- Add more CLI commands: `quests browse`, `start <id>`, etc.
- Implement quest parsing and sandbox abstraction
- Expand test coverage for new features

---

## Features (Planned)
- Real-world DevOps quests (file, container, and VM sandboxes)
- XP, streak, and progress tracking
- YAML-driven quest definitions
- Structured logging and journaling
- Optional sync/backup to GitHub or S3
- Privacy-first, local-first design

## Quick Start
```sh
# Build the CLI
cd levelupops
make build  # or: go build -o levelupops ./cmd/levelupops

# Initialize LevelUpOps
./levelupops init

# Browse available quests
./levelupops quests browse
```

## Project Structure
```
/cmd/levelupops/         # CLI entrypoint
/internal/quests/        # Quest parsing, XP logic
/internal/sandbox/       # Sandbox engine (file/container/vm)
/internal/progress/      # XP tracking, journal logging
/internal/logger/        # Logrus-based logging
/internal/telemetry/     # Opt-in analytics
/internal/config/        # Config parsing
/pkg/tui/                # Terminal UI (BubbleTea/Tview)
/pkg/graphql/            # GraphQL API (future)
/pkg/utils/              # Shared helpers
/quests/                 # Quest definitions
```

## Development
- Go 1.22+
- [Cobra](https://github.com/spf13/cobra) for CLI
- [Logrus](https://github.com/sirupsen/logrus) for logging
- [SQLite](https://github.com/mattn/go-sqlite3) for XP store
- [Testify](https://github.com/stretchr/testify) for testing

## Testing
```sh
go test ./...
```

## Continuous Integration
- GitHub Actions runs all tests on push/PR (see `.github/workflows/ci.yml`)


## License
- GNU AGPL v3.0 (see `LICENSE` file)

## Contributing
See `CONTRIBUTING.md` (coming soon). PRs and feedback welcome!
