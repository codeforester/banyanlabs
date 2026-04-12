# Banyan Labs

> Infrastructure, tooling, and DevOps scaffolding for small to medium enterprises — built to scale from zero to ~600 engineers.

---

## What is Banyan Labs?

Banyan Labs is an open, opinionated infrastructure and DevOps framework designed to help engineering teams — from solo founders to mid-size companies — bootstrap, manage, and scale their technical stack with consistency and confidence.

It is not a SaaS product. It is a living repository of patterns, tools, scripts, and services that teams can adopt, fork, and extend.

---

## Vision

The goal is to give small and medium engineering teams the same infrastructure discipline that large companies have, without the overhead of building it all from scratch.

**Target scale:** Teams up to ~2,000 people, ~600 engineers.
**Not intended for:** Hyperscalers like Google, Apple, or Oracle.

---

## Roadmap

### Phase 1 — Foundation (Months 1–3)
- Set up project structure, tooling wrappers, and documentation
- Implement Bash bootstrapping layer
- Implement Python setup and CLI layer
- Define YAML manifest-based dependency management
- Build core Bash standard library (logging, error handling)
- Build Python virtual environment management
- Daily progress tracked via GitHub Issues

### Phase 2 — Collaboration
- Onboard ex-colleagues and contributors
- Expand multi-repo support
- Formalize contribution guidelines

### Phase 3 — Public Launch
- Announce to the broader DevOps and engineering community
- Gather feedback and iterate

### Phase 4 — Commercial Viability
- Explore sustainable monetization models
- Build partnerships with cloud providers and tooling vendors

---

## Tech Stack

| Layer | Technology |
|---|---|
| Bootstrapping | Bash |
| Scripting & Glue | Python |
| CLIs (performance-critical) | Go |
| Services (learning & production) | Go, Java |
| Infrastructure as Code | Terraform |
| Configuration formats | YAML, HCL (HashiCorp), JSON |
| Package/dependency manifest | Custom YAML manifest |
| Version control | Git (GitHub) |
| Issue tracking | GitHub Issues |

---

## Architecture Overview

### Two-Layer Setup System

Setup is intentionally split into two layers:

**Layer 1 — Bash Bootstrap**
- Ensures Homebrew is installed
- Installs Python via Homebrew
- Installs any system-level prerequisites (e.g. Xcode CLI tools)
- Hands off to the Python layer once the environment is ready

**Layer 2 — Python Setup**
- Reads a YAML manifest that defines all tools, packages, and services to install
- Manages installation order (respecting inter-package dependencies)
- Manages Python virtual environments per project
- Handles cloning of additional repositories required for a company's full stack

### Wrapper Pattern

Every CLI — whether Bash or Python — is invoked through a wrapper, not directly. This avoids boilerplate repetition and enforces consistency.

**Bash Wrapper**
- Sources the Bash standard library (logging, error handling, utilities)
- Sets up common environment variables
- Discovers and invokes the target Bash script

**Python Wrapper**
- Activates the correct Python virtual environment
- Discovers the target Python CLI by convention
- Passes arguments through cleanly

### CLI Structure Convention

Every Python CLI is a directory (package), not a standalone file. Each CLI directory contains:

```
my-cli/
  __init__.py
  main.py
  README.md
  tests/
  submodules/
```

### YAML Manifest

All dependencies — Python packages, system tools, CLIs, services — are declared in a single YAML manifest. There is no `requirements.txt` or `setup.py`. The Python setup layer reads the manifest and handles installation.

This keeps configuration unified, readable, and language-agnostic.

---

## Multi-Repo Design

Banyan Labs is designed to be the **onboarding and bootstrapping layer** for an entire company's technical ecosystem. It is not a monorepo.

- Banyan Labs itself is the entry point
- The YAML manifest can reference and clone additional repositories
- Each product team, service, or tooling area lives in its own repo
- Banyan Labs orchestrates the setup of the full environment in one step

---

## Repository Structure

```
banyan-labs/
  README.md
  docs/
    01-vision.md
    02-architecture.md
    03-setup-strategy.md
    04-multi-repo-design.md
    05-tech-stack.md
    06-cli-conventions.md
    07-yaml-manifest-spec.md
  scripts/
    wrapper.sh           # Bash wrapper
    wrapper.py           # Python wrapper
    bootstrap.sh         # Layer 1 setup
    lib/
      stdlib.sh          # Bash standard library
  setup/
    setup.py             # Layer 2 Python setup entrypoint
    manifest.yaml        # Dependency manifest
  services/
    url-shortener/       # Example Go service (learning project)
  infra/
    terraform/
  .github/
    ISSUE_TEMPLATE/
```

---

## Git Workflow

- All work is done on feature branches, never directly on `main`
- Branch naming convention: `<initials>-<MMDD>-<short-description>-<issue-number>`
  - Example: `pd-0411-bash-stdlib-12`
- Every branch is backed by a GitHub Issue
- PRs are opened against `main` and merged after self-review
- Branches are deleted after merge

---

## Dependency Philosophy

- No `requirements.txt`
- No `setup.py`
- All Python dependencies are declared in `manifest.yaml`
- The Python setup layer installs them in the declared order
- This project is not distributed as a pip-installable package — it is used directly from the repo

---

## Documentation

All architecture decisions, design rationale, and implementation notes are documented in `/docs`. Markdown files are numbered to indicate reading order. The main `README.md` links to all docs.

When sharing project context with other tools (e.g. OpenAI Codex), paste the relevant docs to provide architectural grounding.

---

## Status

**Current phase:** Phase 1 — active development started April 2025.

Progress is tracked daily via [GitHub Issues](../../issues).

---

## Contributing

Collaboration opens in Phase 2. If you are one of the initial collaborators, reach out directly. Public contributions will be welcomed in Phase 3 after the initial structure is stable.

---

*Banyan Labs — Infrastructure that grows with you.*
