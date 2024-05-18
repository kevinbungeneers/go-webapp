{ pkgs }:
let
  config = {
    package = pkgs.go;
  };
in
{
  process.go = {
    beforeExec = "";
    exec = {};
    afterExec = "";
  };

  buildInputs = [
    config.package
  ];

  shellHook = ''
    if [[ ! -d "$(pwd)/.devshell/data/go" ]]; then
      mkdir -p "$(pwd)/.devshell/data/go"
    fi

    export GOPATH="$(pwd)/.devshell/data/go"
    export GOROOT="${config.package}/share/go/"
    export PATH=$GOPATH/bin:$PATH
  '';
}
