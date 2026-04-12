# `setup`

Bootstrap the local Banyan Labs CLI environment on macOS.

## What It Does

The command is intentionally small and idempotent.

## Commands

- `install`
  Installs Homebrew, Xcode Command Line Tools, Python, and creates `$HOME/.banyan_venv`.
- `update-profile`
  Reserved for future shell profile updates. This subcommand is not implemented yet.

## Install Behavior

The `install` command performs these steps:

1. install Homebrew if it is not already installed
2. install Xcode Command Line Tools if they are not already installed
3. install Python via Homebrew if it is not already installed
4. create `$HOME/.banyan_venv` if it does not already exist

## What It Does Not Do Yet

- update shell profiles such as `~/.bashrc` or `~/.zshrc`
- uninstall previously installed tools
- manage application-specific Python packages inside the virtual environment

## Usage

Via the wrapper:

```bash
cli/bash/bin/bash-wrapper setup.sh install
```

Via the symlinked entrypoint:

```bash
cli/bash/bin/setup.sh install
```

Help:

```bash
cli/bash/bin/setup.sh --help
```

Dry run:

```bash
cli/bash/bin/setup.sh install --dry-run
```

Verbose/debug logging:

```bash
cli/bash/bin/setup.sh -v install
```

## Configuration

The command supports a few environment-variable overrides, mainly for automation and tests:

- `BANYAN_SETUP_VENV_DIR`
- `BANYAN_SETUP_PYTHON_FORMULA`
- `BANYAN_SETUP_PYTHON_BIN`
- `BANYAN_SETUP_BREW_BIN`
- `BANYAN_SETUP_HOMEBREW_INSTALLER_SCRIPT`
- `BANYAN_SETUP_XCODE_COMMAND_LINE_TOOLS_DIR`

## Tests

Run the command test suite with:

```bash
bats cli/bash/commands/setup/tests/setup.bats
```
