build-wasm:
	GOOS=wasip1 GOARCH=wasm tinygo build -no-debug -scheduler=none -panic=trap -o ./docs/strip-literal.wasm ./cmd/strip-literal.wasm/ && du -sh -B1 --apparent-size ./docs/strip-literal.wasm
