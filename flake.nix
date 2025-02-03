{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_23       # Go compiler and tools
    pkgs.gopls    # Go language server for editor support
    pkgs.git      # Version control system
  ];

  shellHook = ''
    export GOPATH=$PWD/go
    export PATH=$GOPATH/bin:$PATH
    echo "Go development environment loaded."
  '';
}

