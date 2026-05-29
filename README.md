# Banyan Labs

Banyan Labs is now intended to be a project repo that lives inside a shared
workspace managed by the sibling `base` repo.

## Relationship to Base

The shared developer bootstrap layer no longer lives here.

That foundational layer now belongs in `base`, including:

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

As the split continues, Banyan Labs should increasingly focus on project-level
concerns such as:

- project source code
- infrastructure definitions
- service code
- manifests and configuration specific to Banyan Labs
- project-specific documentation
- project-specific tests and automation

## Status

This repo is in transition.

The first migration pass moved the shared CLI/bootstrap artifacts out of
`banyanlabs` and into the sibling `base` repo. Additional project-specific
structure will be rebuilt here on top of that new workspace model.

## Development Workflow

Banyan Labs follows Base's issue-first GitHub workflow:

- track work in GitHub Issues before implementation starts
- prefer `basectl gh` for supported GitHub issue and pull request tasks
- use one worktree per pull request
- keep each PR scoped to one issue and link it with `Fixes #<issue>`

See [CONTRIBUTING.md](CONTRIBUTING.md) and
[docs/github-workflow.md](docs/github-workflow.md) for the full workflow.
