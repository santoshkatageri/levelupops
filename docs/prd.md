
# 🚀 LevelUpOps – Gamified DevOps Learning Platform

> A CLI-first, cross-platform DevOps learning tool built in Go. Learn infrastructure hands-on with quests, earn XP, and sync progress via GitHub or the cloud to grow your portfolio.

---

## 🎯 Core Objectives

LevelUpOps empowers learners to:

- ✅ Gain **real-world DevOps skills** through interactive quests
- ✅ Use **local, container, or VM-based sandboxes**
- ✅ Track learning with **XP, streaks, and daily logs**
- ✅ Build a **public portfolio and resume**
- ✅ Stay private-first with opt-in sync and backups

---

## 🧩 Functional Requirements

### 🔧 CLI Core

- `init`, `browse`, `start`, `progress`, `cleanup` commands
- XP and streak tracking (file-based or embedded DB)
- YAML-driven quest metadata format
- Structured command validations and Logrus-based logging

### 🧱 Sandbox Engine

- File-based sandbox (Phase 1)
- Container sandbox using **Podman** or Docker (Phase 2)
- VM sandbox using **Lima**, **WSL2**, or **Multipass** (Phase 3)
- Abstract sandbox layer for pluggable runtime

### 📊 Logging & Telemetry

- Structured logs via Logrus
- Quest run logs and daily XP entries
- TTL-aware cleanup of stale containers/VMs
- Auto-backup to GitHub or S3
- GDPR-compliant telemetry (opt-in only)

### ☁️ XP Sync & Journal Backup

- Daily XP log and progress export to GitHub repo
- Optional S3 backup
- Journaling and leaderboard integration (planned)

### 🧙 TUI Layer

- Quest browser in terminal
- XP bars, streak history, visual feedback
- BubbleTea or Tview powered UI

### 🔌 GraphQL API + Web Dashboard

- gqlgen-based API for XP, quests, streaks
- User dashboards with sync toggle
- Progress overview and collaboration readiness

### 🌐 Web Portfolio & Resume Builder

- Markdown XP journals converted to blog
- Resume generator using Jinja2 or Overleaf export
- Hosted on GitHub Pages or levelupops.dev subdomain

---

## 🧪 Non-Functional Requirements

- OS support: macOS (Intel/Apple Silicon), Linux, Windows
- No root permissions required (uses rootless tools)
- Binary size < 20MB for CLI
- Local-first with optional sync
- Modular Go packages
- Privacy and GDPR aligned

---

## 🔐 Security & Privacy

- User consent mandatory for public sync
- No hardcoded credentials
- TLS/HTTPS for all remote transfers
- Environment variables or `.env` for secrets

---

## ⚙️ DevOps & CI/CD

- GitHub Actions for CLI + Infra deployment
- GoReleaser for cross-platform binary builds
- Packer + Terraform for OCI-based Dev Sandbox
- Justfile for local automation tasks
- Tailscale setup for private remote sync

---

## 📦 Packaging & Distribution

- CLI binaries (macOS, Linux, Windows)
- DevContainer + Justfile for reproducible dev
- One-liner install scripts for easy onboarding
- Support for `levelupops doctor` diagnostics

---

## 🛠️ Tech Stack

| Component            | Tech / Tool                      | Purpose                                   |
|---------------------|----------------------------------|-------------------------------------------|
| CLI Core             | Go + Cobra                       | Cross-platform CLI app                    |
| Quest Definitions    | YAML                             | Modular and readable quest structure      |
| Logging              | Logrus                           | Logs and audit trail                      |
| Local Sandbox        | File-based isolated folders      | Lightweight starting point                |
| Container Sandbox    | Podman (rootless) / Docker       | Real-world infra experience               |
| VM Sandbox           | Lima, WSL2, Multipass            | Full system simulation                    |
| TUI Layer            | BubbleTea / Tview                | XP visuals and streaks UI                 |
| Database             | SQLite / LiteFS / BoltDB         | XP tracking and journal persistence       |
| XP Sync              | GitHub + S3                      | Remote sync and portfolio backup          |
| DevOps Automation    | GitHub Actions + GoReleaser      | CI/CD pipelines                           |
| Infra Provisioning   | Terraform + Packer (OCI Free)    | Remote Dev Sandbox                        |
| Resume Builder       | Python + Jinja2 / Overleaf       | Portfolio generation                      |
| Observability        | TTL logs + cleanup CLI           | Maintain sandbox hygiene                  |

---

## 📅 Phase Milestones

| Phase | Description                              |
|-------|------------------------------------------|
| 1     | CLI + File-based Quests                  |
| 2     | Container Sandbox                        | 
| 3     | VM Sandbox                               |
| 4     | Terminal UI with BubbleTea               | 
| 5     | GraphQL API + Web Dashboard              | 
| 6     | Resume Builder + Public Portfolio Page   | 
---

## 🧠 Developer Motto

> Build Systems. Build Skills. Build Stability.

---

## 📬 Stay Connected

- Blog: [https://learnedops.com](https://learnedops.com)
- GitHub: [github.com/learnedops/levelupops](https://github.com/learnedops/levelupops)
- Email: [contact@learnedops.com](mailto:contact@learnedops.com)

---

## 🤝 Contributing

We welcome learners, teachers, and infra nerds alike. Check `CONTRIBUTING.md` and start building the future of DevOps education.

