{
  inputs.flake-utils.url = "github:numtide/flake-utils";
  outputs = {
    self,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import <nixpkgs> {};
    in {
      devShell = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [go gopls];
        buildInputs = [];
      };
    });
}
