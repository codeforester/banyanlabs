# Skills

This file documents repeatable AI-assisted development workflows for Banyan
Labs. General workflow context lives in `CONTRIBUTING.md` and
`docs/github-workflow.md`.

## GitHub Issue And PR Workflow

Use this workflow when creating GitHub issues, branches, worktrees, or pull
requests for Banyan Labs.

- Prefer `basectl gh` for supported GitHub workflows so Banyan Labs follows the
  same issue-first train as Base.
- Fall back to the GitHub connector, raw `gh`, or `git` when `basectl gh` does
  not support the needed operation or local tooling is unavailable.
- Assign Codex-created issues to `codeforester`.
- Use GitHub default-style labels:
  - `bug` for defects.
  - `enhancement` for features, refactors, and most maintenance.
  - `documentation` for docs-only changes.
  - `ci` for GitHub Actions, tests, and release automation.
  - `security` for hardening, dependency pinning, and static analysis.
- Do not create new `type:*` labels.
- Name branches as `<category>/<issue>-<YYYYMMDD>-<slug>`, for example
  `documentation/48-20260608-agent-workflow-guidance`.
- Do all pull request implementation work in a dedicated worktree under
  `~/work/banyanlabs-worktrees/<slug>`.
- Before creating a worktree, check whether the current checkout is already a
  linked worktree for the issue. Do not create nested or duplicate worktrees.
- Keep the PR worktree available while review feedback is pending. After merge,
  sync `main`, remove the worktree, and delete local and remote branches.
- Link PRs to issues with `Fixes #<issue>` or `Closes #<issue>`.

## Debug and verify Banyan Labs behavior

Use this workflow when investigating failed service behavior, broken API smoke
tests, local setup drift, CI failures, observability gaps, deployment issues,
or infrastructure surprises.

- Read the full error output first, including stack traces, command output,
  logs, request/response bodies, and paths.
- Reproduce the symptom from a clean command line before fixing it. If the
  issue is not reproducible, gather more evidence instead of guessing.
- Check recent changes with `git status`, `git diff`, and relevant issue or PR
  context.
- Trace the bad value or failed state to its source. For cross-layer failures,
  inspect each boundary separately: service handler, domain logic, storage,
  config, runtime dependency, API contract, CI job, and Base-managed project
  command.
- Form one hypothesis, make the smallest change that tests it, and rerun the
  focused verification. Do not stack unrelated fixes.
- Before claiming completion, run the command that proves the claim in the
  current checkout or worktree and read the output. Report the command and the
  result in the PR and final summary.

## Add a Banyan Labs service

Use this workflow when adding or changing a Go, Java, or Python-backed service.

- Keep service code, tests, fixtures, API contracts, and service-local docs
  close to the owning service.
- Start behavior changes and bug fixes with a failing test, fixture, API
  smoke test, or documented reproduction whenever practical.
- Keep shared workstation setup in Base. Put Banyan Labs project runtime
  declarations in `base_manifest.yaml`, `Brewfile`, `.mise.toml`, and service
  docs.
- Add or update CI coverage for the service language.
- Preserve shared operational contracts for config, health, logs, metrics,
  traces, tests, and deployment, while using language-native libraries.

## Add or change project tooling

Use this workflow when changing `base_manifest.yaml`, `.mise.toml`, `Brewfile`,
test tooling, or project-level developer dependencies.

- Keep Base as the workspace control plane and Banyan Labs as the project owner.
- Prefer one working local path before adding alternate toolchains.
- Document how the tool is checked, installed, and used.
- Update `CONTRIBUTING.md` or service docs when developer workflow changes.
- Validate with `tests/validate.sh`, `git diff --check`, and relevant project
  checks.

## Change GitHub workflow behavior

Use this workflow when changing issues, PRs, labels, CODEOWNERS, worktrees,
GitHub Actions, or repo-local agent guidance.

- Keep workflow rules aligned with `docs/github-workflow.md`.
- Do not add `type:*` labels.
- Keep PRs issue-backed and scoped.
- Update `.github/pull_request_template.md` when PR expectations change.
- Validate `skills.md` changes against `.github/workflows/skills.yml`.

## Release-facing changes

Use this workflow when changing versioning, packaging, deployment, or release
notes.

- Update `CHANGELOG.md` for notable user-visible, platform, or release-worthy
  changes.
- Keep release intent in milestones and durable docs, not only in PR text.
- Validate build, test, and deployment commands that the release path depends
  on.

## Add or revise a Banyan Labs agent workflow

Use this workflow when adding or changing reusable AI-assisted development
guidance such as this `skills.md` file, `AGENTS.md`, or workflow documentation.

- Put durable repo-local rules in `AGENTS.md`, `CONTRIBUTING.md`, `skills.md`,
  or focused docs. Do not add personal Codex runtime settings to the repo.
- Prefer trigger-focused workflow names and descriptions. The first lines
  should make it clear when the workflow applies.
- Keep entries concise and Banyan Labs-specific. Link to focused docs for
  longer policy instead of duplicating it.
- Align examples with Banyan Labs conventions: `basectl gh`, `origin/main`,
  `<category>/<issue>-<YYYYMMDD>-<slug>`, and
  `~/work/banyanlabs-worktrees/<slug>`.
- Review the workflow against likely pressure cases: time pressure, ambiguous
  review feedback, failing tests, dirty worktrees, and temptation to move
  Base-owned workspace behavior into Banyan Labs.
- Validate documentation-only workflow changes with `git diff --check`. If a
  CI workflow validates the guidance, run or update that validation too.
