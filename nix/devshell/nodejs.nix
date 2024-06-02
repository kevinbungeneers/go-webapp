{ pkgs }:

let
  config = {
    package = pkgs.nodejs_20;
  };
in
{
  process.nodejs = {
    beforeExec = "";
    exec = {
        "tailwind_watch" = "npx tailwindcss -i ./web/assets/styles/app.css -o ./web/public/css/app.css --watch";
    };
    afterExec = "";
  };

  buildInputs = [
    config.package
  ];

  shellHook = ''
    if [[ ! -f "package-lock.json" ]]; then
        echo "package-lock.json was not found"
        exit
    fi

    ACTUAL_NPM_CHECKSUM="${config.package.version}:$(${pkgs.nix}/bin/nix-hash --type sha256 package-lock.json)"
    NPM_CHECKSUM_FILE="node_modules/package-lock.json.checksum"
    if [ -f "$NPM_CHECKSUM_FILE" ]
      then
        read -r EXPECTED_NPM_CHECKSUM < "$NPM_CHECKSUM_FILE"
      else
        EXPECTED_NPM_CHECKSUM=""
    fi

    if [[ "$ACTUAL_NPM_CHECKSUM" != "$EXPECTED_NPM_CHECKSUM" ]]; then
      if ${config.package}/bin/npm install; then
        echo "$ACTUAL_NPM_CHECKSUM" > "$NPM_CHECKSUM_FILE"
      else
        echo "Npm install failed. Run 'npm install' manually."
      fi
    fi
  '';
}
