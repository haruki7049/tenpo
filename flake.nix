{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-compat.url = "github:edolstra/flake-compat";
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
      ];

      imports = [
        inputs.treefmt-nix.flakeModule
      ];

      perSystem =
        {
          config,
          lib,
          pkgs,
          ...
        }:
        let
          env.LD_LIBRARY_PATH = lib.makeLibraryPath buildInputs;
          buildInputs = lib.optionals pkgs.stdenv.isLinux [
            pkgs.libx11
            pkgs.libxrandr
            pkgs.libGL
            pkgs.libxcursor
            pkgs.libxinerama
            pkgs.libxi
            pkgs.libxxf86vm
            pkgs.libglvnd
          ];
          nativeBuildInputs = [
            pkgs.go # Golang
            pkgs.nil # Nix LSP
            pkgs.gopls # Golang LSP
            pkgs.nushell # For scripts runner
          ];
        in
        {
          treefmt = {
            projectRootFile = ".git/config";

            # Nix
            programs.nixfmt.enable = true;

            # Go
            programs.gofmt.enable = true;

            # GitHub Actions
            programs.actionlint.enable = true;

            # Markdown
            programs.mdformat.enable = true;

            # ShellScript
            programs.shellcheck.enable = true;
            programs.shfmt.enable = true;
          };

          devShells.default = pkgs.mkShell {
            inherit buildInputs nativeBuildInputs env;
            inputsFrom = [ config.treefmt.build.devShell ];
          };
        };
    };
}
