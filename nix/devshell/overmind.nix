{ pkgs, lib, processes }:
let
   processText = processName: processAttr:
       let
         execSet = processAttr.exec;
         scripts = lib.mapAttrsToList (name: value: "${processName}-${name}: exec ${value}") execSet;
       in
         builtins.concatStringsSep "\n" scripts;

   processList = lib.mapAttrsToList processText processes;

   procfile = pkgs.writeText "Procfile" (builtins.concatStringsSep "\n" processList);

   beforeScripts = lib.mapAttrsToList (name: process: "${process.beforeExec}") processes;

   beforeScript = if beforeScripts != [] then ''
    ${builtins.concatStringsSep "\n" beforeScripts}
   '' else "";

   afterScripts = builtins.concatStringsSep "\n" (lib.mapAttrsToList (name: process: "${process.afterExec}") processes);

   afterScript = if afterScripts != "" then ''
    stop_up() {
      ${afterScripts}
    }
    trap stop_up SIGINT SIGTERM
   '' else "";

   # Had to wrap overmind so that I could set the root directory. It defaults to the dir where the procfile is located
   # and there's no env var to override this dir value, for some reason.
   # The wrapping also allows me to run scripts before and after execution.
   overmind = pkgs.writeShellScriptBin "overmind" ''
     extraParams=""
     case "$@" in
       'start')
         ;;
       's')
         extraParams="--root $PROJECT_ROOT"
         ;;
     esac

     ${beforeScript}

     ${afterScript}

     OVERMIND_PROCFILE=${procfile} ${pkgs.overmind}/bin/overmind "$@" $extraParams
   '';
in
{
  buildInputs = [
    overmind
  ];
}
