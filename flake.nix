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
        lib = pkgs.lib;
        treefmtEval = treefmt-nix.lib.evalModule pkgs ./treefmt.nix;
        rust = pkgs.rust-bin.fromRustupToolchainFile ./rust-toolchain.toml;
        craneLib = (crane.mkLib pkgs).overrideToolchain rust;
        src = ./.;
        cargoArtifacts = craneLib.buildDepsOnly {
          inherit src;
          buildInputs = bevyengine-dependencies;
          nativeBuildInputs = [ pkgs.pkg-config ];
        };
        kosu = craneLib.buildPackage {
          inherit src cargoArtifacts;
          buildInputs = bevyengine-dependencies;
          nativeBuildInputs = [ pkgs.pkg-config ];
          strictDeps = true;

          doCheck = true;
        };
        cargo-clippy = craneLib.cargoClippy {
          inherit cargoArtifacts src;
          buildInputs = bevyengine-dependencies;
          nativeBuildInputs = [ pkgs.pkg-config ];
          cargoClippyExtraArgs = "--verbose -- --deny warnings";
        };
        cargo-doc = craneLib.cargoDoc {
          inherit cargoArtifacts src;
          buildInputs = bevyengine-dependencies;
          nativeBuildInputs = [ pkgs.pkg-config ];
        };
        bevyengine-dependencies = with pkgs; [
          udev
          alsa-lib
          vulkan-loader

          # To use the x11 feature
          xorg.libX11
          xorg.libXcursor
          xorg.libXi
          xorg.libXrandr

          # To use the wayland feature
          libxkbcommon
          wayland
        ];
        llvm-cov = pkgs.stdenv.mkDerivation {
          name = "llvm-cov";
          inherit src;

          buildInputs = [
            rust
            pkgs.cargo-llvm-cov
          ];

          buildPhase = ''
            cargo llvm-cov --html --offline
            mkdir $out/share
          '';

          installPhase = ''
            cp target/llvm-cov/html/* $out/share
            chmod -R 666 $out/share/*
          '';

          meta.allowSubstitutes = true;
          meta.allowUnfree = true;
        };
      in
      {
        formatter = treefmtEval.config.build.wrapper;

        packages.default = kosu;
        packages.doc = cargo-doc;
        packages.cov = llvm-cov;

        apps.default = flake-utils.lib.mkApp {
          drv = self.packages.${system}.default;
        };

        checks = {
          inherit kosu cargo-clippy cargo-doc;
          formatting = treefmtEval.config.build.check self;
        };

        devShells.default = pkgs.mkShell rec {
          buildInputs = bevyengine-dependencies ++ [
            rust
          ];

          nativeBuildInputs = [
            pkgs.pkg-config
          ];

          LD_LIBRARY_PATH = lib.makeLibraryPath buildInputs;

          shellHook = ''
            export PS1="\n[nix-shell:\w]$ "
          '';
        };
      });
}
