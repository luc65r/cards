{ lib
, buildGoModule
, pkg-config
, xorg
, glfw
}:

buildGoModule rec {
  pname = "cards";
  version = "v0.1";
  src = ./.;
  vendorSha256 = null;

  nativeBuildInputs = [
    pkg-config
  ];

  propagatedBuildInputs = [
    glfw
  ];

  buildInputs = with xorg; [
    libX11
    libXcursor
    libXrandr
    libXinerama
    libXi
    libXext
    libXxf86vm
  ];
}
