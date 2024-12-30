# shell.nix
{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # Go development
    go                    # Go compiler and tools
    gopls                 # Go language server
    go-tools              # Additional Go tools
    golangci-lint        # Go linter
    delve                # Go debugger

    # Development tools
    git
    gnumake
    pre-commit

    # Optional but useful tools
    air                   # Live reload for Go development
    go-mockery           # Mock generator
    gomodifytags         # Modify struct field tags
    gore                 # Go REPL

    # Shell utilities
    direnv               # Per-directory environment variables
  ];

  shellHook = ''
    # Set GOPATH to a directory within the project
    export GOPATH="$PWD/.go"
    export PATH="$GOPATH/bin:$PATH"

    # Create necessary directories
    mkdir -p .go/bin

    # Initialize git hooks
    if [ -e .git/hooks ]; then
      pre-commit install --install-hooks
    fi

    # Set up direnv if .envrc doesn't exist
    if [ ! -e .envrc ]; then
      echo "use nix" > .envrc
      direnv allow
    fi

    echo "Go development environment loaded!"
    echo "Go version: $(go version)"
    echo "GOPATH: $GOPATH"
  '';

  # Environment variables
  LANG = "en_US.UTF-8";
  GO111MODULE = "on";
  CGO_ENABLED = "1";
}
