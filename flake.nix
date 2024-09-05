{
  description = "A basic gomod2nix flake";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";
  inputs.gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
  inputs.gomod2nix.inputs.flake-utils.follows = "flake-utils";
  inputs.flake-compat.url =
    "https://flakehub.com/f/edolstra/flake-compat/1.tar.gz";

  outputs = { self, nixpkgs, flake-utils, gomod2nix, flake-compat }:
    (flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        # The current default sdk for macOS fails to compile go projects, so we use a newer one for now.
        # This has no effect on other platforms.
        callPackage =
          pkgs.darwin.apple_sdk_11_0.callPackage or pkgs.callPackage;
      in {
        packages.default =
          gomod2nix.legacyPackages.${system}.buildGoApplication {
            pname = "compare_versions";
            version = "0.1";
            pwd = ./.;
            src = ./.;
            modules = ./gomod2nix.toml;
          };
        devShells.default = callPackage ./shell.nix {
          inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
        };
      }));
}
