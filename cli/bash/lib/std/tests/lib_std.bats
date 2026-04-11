#!/usr/bin/env bats

load ../../../tests/test_helper.bash

setup() {
    setup_test_tmpdir
    source "$BANYAN_BASH_DIR/lib/std/lib_std.sh"
}

@test "add_to_path appends an existing directory only once" {
    mkdir -p "$TEST_TMPDIR/bin"
    PATH="/usr/bin"

    add_to_path "$TEST_TMPDIR/bin"
    add_to_path "$TEST_TMPDIR/bin"

    [ "$PATH" = "/usr/bin:$TEST_TMPDIR/bin" ]
}

@test "add_to_path prepends when requested" {
    mkdir -p "$TEST_TMPDIR/bin"
    PATH="/usr/bin"

    add_to_path -p "$TEST_TMPDIR/bin"

    [ "$PATH" = "$TEST_TMPDIR/bin:/usr/bin" ]
}

@test "add_to_path skips missing directories unless -n is used" {
    PATH="/usr/bin"

    add_to_path "$TEST_TMPDIR/missing"
    [ "$PATH" = "/usr/bin" ]

    add_to_path -n "$TEST_TMPDIR/missing"
    [ "$PATH" = "/usr/bin:$TEST_TMPDIR/missing" ]
}

@test "run honors dry-run mode without executing the command" {
    local target="$TEST_TMPDIR/dry-run.txt"
    DRY_RUN=true

    run touch "$target"

    [ "$status" -eq 0 ]
    [ ! -e "$target" ]
}

@test "safe_touch creates files" {
    local target="$TEST_TMPDIR/touched.txt"

    safe_touch "$target"

    [ -f "$target" ]
}

@test "assert_arg_count accepts exact and ranged matches" {
    assert_arg_count 2 2
    assert_arg_count 2 1 3
}

@test "assert_arg_count exits when the count is out of range" {
    bats_run assert_arg_count 4 1 3

    [ "$status" -eq 1 ]
    [[ "$output" == *"Argument count mismatch"* ]]
}

@test "safe_unalias removes aliases and ignores missing ones" {
    alias ll='ls -l'

    safe_unalias ll missing_alias

    ! alias ll >/dev/null 2>&1
}
