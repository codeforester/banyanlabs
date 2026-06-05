# Contributing to Banyan Labs

Banyan Labs uses the same issue-first GitHub workflow as the sibling
[Base](https://github.com/codeforester/base#readme) repository. Keep work
visible in GitHub Issues, keep pull requests small, and use Base's `basectl gh`
helper when it supports the operation.

## AI-Assisted Development

Coding agents should follow [AGENTS.md](AGENTS.md). It points to the same
workflow in this guide while capturing Banyan Labs-specific instructions for
Base tooling, validation, and design-only sessions.

## Workflow

1. Create or choose a GitHub issue before starting implementation work.
2. Apply one primary category label:
   - `bug` for defects, regressions, or correctness issues.
   - `enhancement` for new capabilities, product improvements, refactors, and
     most maintenance work.
   - `documentation` for documentation-only work.
   - `ci` for GitHub Actions, tests, release automation, or CI reliability.
   - `security` for security hardening, dependency pinning, static analysis, or
     permission tightening.
3. Create a branch from the issue using this convention:

   ```text
   <category>/<issue>-<YYYYMMDD>-<slug>
   ```

   Example:

   ```text
   enhancement/42-20260529-add-url-shortener-api
   ```

4. Use an isolated Git worktree for each pull request:

   ```bash
   git fetch origin main
   git worktree add ../banyanlabs-worktrees/<slug> -b <branch> origin/main
   ```

5. Keep the PR scoped to one issue. Avoid unrelated refactors.
6. Link the PR back to the issue with `Fixes #<issue>`.
7. After merge, sync `main`, remove the worktree, and delete the local and
   remote branches.

For the full policy, including milestone and GitHub Project guidance, see
[GitHub Workflow](docs/github-workflow.md).

## Base Tooling

Use the sibling Base checkout as the shared workspace control plane:

```text
~/work/base
~/work/banyanlabs
```

When `basectl` is on `PATH`, prefer:

```bash
basectl gh issue create --category enhancement --title "Add ..."
basectl gh issue start <issue-number>
basectl gh pr create
basectl gh pr checks
```

If `basectl` is not yet on `PATH`, invoke it from the sibling checkout:

```bash
~/work/base/bin/basectl gh issue list
```

Fallback to `gh` only when `basectl gh` does not support the needed operation.

## Validation

Run the narrowest relevant checks first, then broaden when the change touches
shared behavior.

Current general checks:

```bash
git diff --check
```

As Banyan Labs code returns to the repository, add project-specific checks here
for Go, frontend, infrastructure, database migrations, and service tests.

## Pull Request Checklist

Before opening a PR:

- The branch name follows `<category>/<issue>-<YYYYMMDD>-<slug>`.
- The PR is scoped to one issue.
- The PR body explains what changed and how it was validated.
- Documentation is updated when behavior, setup, or workflow changes.
- The PR includes `Fixes #<issue>` when it should close the issue.
