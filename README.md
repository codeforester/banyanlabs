# Banyan Labs

Banyan Labs is a realistic infrastructure and platform engineering lab.

The purpose is to build deep DevOps, SRE, and platform engineering knowledge by
assembling real services and infrastructure tools into a meaningful environment.
Certifications can introduce AWS, GCP, Azure, Kubernetes, Terraform, CI/CD,
monitoring, and distributed systems, but Banyan Labs is about learning those
ingredients by building with them.

The long-term goal is to grow toward the complexity of a medium-sized
engineering organization: multiple services, multiple languages, local and
remote environments, CI/CD, infrastructure as code, observability, Kubernetes,
and eventually multi-cloud patterns based on real open source tools.

The first concrete service is a Go URL shortener with local SQLite storage,
authentication, user management, tests, and a minimal HTML UI from day one.

See [Banyan Labs Vision](docs/banyanlabs-vision.md) for the broader direction.

## Relationship to Base

The shared developer bootstrap layer no longer lives here.

That foundational layer now belongs in
[`base`](https://github.com/codeforester/base#readme), including:

- shell environment bootstrap
- shared Bash wrapper and command conventions
- common shell libraries
- workspace-level setup and test entrypoints

The intended local workspace shape is:

```text
work/
  base/
  banyanlabs/
  other-project/
```

In that model:

- `base` owns shared developer tooling and workspace orchestration
- `banyanlabs` owns Banyan Labs-specific code, manifests, infrastructure, and
  project behavior

## What This Repo Will Focus On

Banyan Labs focuses on project-level and product-level concerns such as:

- project source code
- infrastructure definitions
- service code
- manifests and configuration specific to Banyan Labs
- project-specific documentation
- project-specific tests and automation

## Status

This repo is being rebuilt around a local-first platform lab.

The shared CLI/bootstrap artifacts moved into the sibling `base` repo. Banyan
Labs now owns the services, docs, tests, and project-specific infrastructure
that will grow on top of that workspace model.

## Development Workflow

Banyan Labs follows Base's issue-first GitHub workflow:

- track work in GitHub Issues before implementation starts
- prefer `basectl gh` for supported GitHub issue and pull request tasks
- use one worktree per pull request
- keep each PR scoped to one issue and link it with `Fixes #<issue>`

See [CONTRIBUTING.md](CONTRIBUTING.md) and
[docs/github-workflow.md](docs/github-workflow.md) for the full workflow.
