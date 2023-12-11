go_mod:
	go mod init tinygo-leaves
    go mod tidy

build: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./go/wasm.go
	cp ./go/wasm.js ./html/
	cp ./go/index.html ./html/

wasm_exec:
	cp `tinygo env TINYGOROOT`/targets/wasm_exec.js ./html/

clean:
	rm -rf ./html
	mkdir ./html
