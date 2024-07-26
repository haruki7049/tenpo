{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    rust-overlay.url = "github:oxalica/rust-overlay";
    crane.url = "github:ipetkov/crane";
  };

  outputs = { self, nixpkgs, treefmt-nix, rust-overlay, flake-utils, crane }:
    flake-utils.lib.eachSystem [ "x86_64-linux" ] (system:
      let
        overlays = [ (import rust-overlay) ];
        pkgs = import nixpkgs {
          inherit overlays system;
        };
        treefmtEval = treefmt-nix.lib.evalModule pkgs ./treefmt.nix;
        rust = pkgs.rust-bin.fromRustupToolchainFile ./rust-toolchain.toml;
        craneLib = (crane.mkLib pkgs).overrideToolchain rust;
        src = ./.;
        cargoArtifacts = craneLib.buildDepsOnly {
          inherit src;
        };
        kosu = craneLib.buildPackage {
          inherit src cargoArtifacts;
          strictDeps = true;

          doCheck = true;
        };
        cargo-clippy = craneLib.cargoClippy {
          inherit cargoArtifacts src;
          cargoClippyExtraArgs = "--verbose -- --deny warnings";
        };
        cargo-doc = craneLib.cargoDoc {
          inherit cargoArtifacts src;
        };
      in
      {
        formatter = treefmtEval.config.build.wrapper;

        packages.default = kosu;
        packages.doc = cargo-doc;

        apps.default = flake-utils.lib.mkApp {
          drv = self.packages.${system}.default;
        };

        checks = {
          inherit kosu cargo-clippy cargo-doc;
          formatting = treefmtEval.config.build.check self;
        };

        devShells.default = pkgs.mkShell {
          packages = [
            rust
          ];

          shellHook = ''
            export PS1="\n[nix-shell:\w]$ "
          '';
        };
      });
}
