with import <nixpkgs> {};

import ./default.nix {
  inherit lib buildGoModule pkg-config xorg glfw;
}
