{
	description = "defines extropic-art-backend dev environment in a nix flake";
	inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
	# TODO insert tests
	outputs = { self, nixpkgs }:
		let
      supportedSystems = [ "x86_64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
		in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
						packages = [
							pkgs.go 
						];
          shellHook = 
          ''
            unset GOPATH
            unset GOROOT
          '';
          };
        });
    };
}
