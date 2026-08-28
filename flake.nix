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
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
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
          system,
          ...
        }:
        let
          overlays = [ inputs.gomod2nix.overlays.default ];
          env.LD_LIBRARY_PATH = lib.makeLibraryPath buildInputs;
          buildInputs = lib.optionals pkgs.stdenv.hostPlatform.isLinux [
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
            pkgs.gomod2nix # gomod2nix for creating hashes (./gomod2nix.toml)
          ];

          tenpo = pkgs.buildGoApplication {
            name = "tenpo";
            src = lib.cleanSource ./.;
            modules = ./gomod2nix.toml;

            inherit buildInputs nativeBuildInputs env;
          };
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system overlays;
          };

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

          packages = {
            inherit tenpo;
            default = tenpo;
          };

          checks = {
            inherit tenpo;
          };

          devShells.default = pkgs.mkShell {
            inherit buildInputs nativeBuildInputs env;
            inputsFrom = [ config.treefmt.build.devShell ];
          };
        };
    };
}
