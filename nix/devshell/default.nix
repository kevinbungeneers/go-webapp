{ pkgs, lib }:

let
  modulesData = [
    { file = ./go.nix; params = { inherit pkgs; }; }
    { file = ./nodejs.nix; params = { inherit pkgs; }; }
  ];
  modules = map ({file, params}: import file params) modulesData;

  processes = builtins.foldl' (a: b: a // b.process) {} modules;

  overmind = import ./overmind.nix {
    inherit pkgs lib processes;
  };

  shellHook = builtins.concatStringsSep "\n" (map (module: module.shellHook) modules);

  buildInputs = (builtins.foldl' (a: b: a ++ b.buildInputs) [] modules) ++ overmind.buildInputs;
in
  pkgs.mkShell {
    inherit shellHook buildInputs;
  }
