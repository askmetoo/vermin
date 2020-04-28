1. <s>Build CI</s>
2. <s>Build Linux/mac install script</s>
3. Use progress (https://github.com/schollz/progressbar/issues/57) 
4. Eliminate need to command dependecies (most of them):
    * arp   (exists on win/nix)
    * ping  (exists on win/nix)
    * vboxmanage (a dependency)
    * wget      (exists on win/nix) - on windows (wget <file url> -o <file output>)
    * ssh		(exists on win/nix)
    * scp       (exists on win/nix)
2. Build Powershell install script (https://raw.githubusercontent.com/habitat-sh/habitat/master/components/hab/install.ps1)
5. Logo drawing