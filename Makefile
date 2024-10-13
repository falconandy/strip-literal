build-wasm:
	GOOS=wasip1 GOARCH=wasm tinygo build -no-debug -scheduler=none -panic=trap -o ./demo/strip-literal.wasm ./cmd/strip-literal.wasm/ && du -sh -B1 ./docs/strip-literal.wasm
