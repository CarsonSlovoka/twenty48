# build for Linux

如果您想用[replit](https://replit.com/)來發佈，請手動編譯

```
export GOOS=js
export GOARCH=wasm
go build -o ../pages/2048.wasm twenty48
export GOOS=linux
export GOARCH=amd64

// 記下GOROOT的路徑
go env GOROOT

// 確認wasm_exec.js真的存在  其中xxx...xxx-go-1.17.5請自行更換成自己的機器代號和go版本
ls /nix/store/xxx...xxx-go-1.17.5/share/go/misc/wasm/wasm_exec.js

// 複製到../pages/wasm_exec.js路徑之中
cp /nix/store/xxx...xxx-go-1.17.5/share/go/misc/wasm/wasm_exec.js ../pages/wasm_exec.js
```
