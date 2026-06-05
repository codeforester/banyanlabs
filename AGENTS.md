# Codex Guidance

This file gives coding agents the repository-specific rules for Banyan Labs. It
is a navigation layer over the existing contributor docs, not a replacement for
them.

## Working Agreement

- Follow `CONTRIBUTING.md` for workflow and `README.md` for the current project
  shape.
- Treat Banyan Labs as the project repository for services, infrastructure,
  manifests, application code, tests, and project-specific automation.
- Treat the sibling `~/work/base` checkout as the shared developer workspace
  control plane.
- Keep shared bootstrap and reusable workspace orchestration changes in Base
  only when they are genuinely cross-project behavior.
- When the user explicitly says a session is design-only or asks for no code
  changes, stay in discussion mode and do not edit files.
- Surface unresolved product or architecture decisions instead of silently
  choosing defaults for broad changes.

## GitHub Workflow

- Create or choose a GitHub issue before implementation work.
- Use one primary category label: `bug`, `enhancement`, `documentation`, `ci`,
  or `security`.
- Do not create or apply `type:*` issue labels.
- Assign Codex-created issues to `codeforester` when GitHub allows it.
- Prefer `basectl gh` for supported issue, branch, PR, check, and cleanup
  operations.
- Use `~/work/base/bin/basectl gh ...` when `basectl` is not on `PATH`.
- Fall back to the GitHub connector, raw `gh`, or `git` when `basectl gh` does
  not support the needed operation or local tooling is unavailable.
- Branch from `origin/main` with `<category>/<issue>-<YYYYMMDD>-<slug>`.
- Use a dedicated worktree under `~/work/banyanlabs-worktrees/<slug>` for PR
  work.
- Link PRs with `Fixes #<issue>` or `Closes #<issue>` when merge should close
  the issue.
- After merge, sync `main`, remove the worktree, and delete local and remote
  branches.

See `docs/github-workflow.md` for the full policy, including milestones,
GitHub Projects, and cleanup rules.

## Local Commands

- Set up the project with `basectl setup banyanlabs`.
- Run the general test entrypoint with `basectl test banyanlabs`.
- Run declared project commands with `basectl run banyanlabs <command>`.
- Current declared commands include `build` and `dev`.

## Validation

- Run the narrowest relevant checks first, then broaden when shared behavior is
  touched.
- For documentation-only changes, run `git diff --check`.
- For general project changes, run `basectl test banyanlabs` and
  `git diff --check`.
- For service changes, include the relevant language-native checks in addition
  to the Base-managed project check.
- For setup, manifest, or command changes, verify the corresponding
  `basectl setup`, `basectl test`, or `basectl run` path.
- If a required check cannot be run locally, say so in the PR and final
  summary.

## Change Boundaries

- Keep URL shortener and future service logic under the owning service
  directory.
- Keep project-specific manifests, infrastructure, and automation in Banyan
  Labs.
- Do not move reusable Base bootstrap behavior back into this repository.
- Use structured parsers and established project tooling instead of ad hoc text
  manipulation when a reasonable option exists.
- Do not add repo-level Codex settings for personal model, approval, or sandbox
  defaults. Those belong in the user's Codex configuration unless the change is
  explicitly about shared repository runtime behavior.
