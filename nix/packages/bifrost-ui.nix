{
  pkgs,
  src,
  version,
}:
pkgs.buildNpmPackage {
  pname = "bifrost-ui";
  inherit version;
  inherit src;
  sourceRoot = "source/ui";

  npmDepsHash = "sha256-JtaSXBfz7toZ8gF9Jjq5iveV/b5obQNhA9VvjyvDky0=";

  # Avoid the upstream build script's copy step (writes outside $PWD).
  npmBuildScript = "build-enterprise";

  installPhase = ''
    runHook preInstall

    mkdir -p "$out/ui"
    cp -R --no-preserve=mode,ownership,timestamps out/. "$out/ui/"

    runHook postInstall
  '';
}