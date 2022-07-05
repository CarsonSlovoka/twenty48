## twenty48

此項目僅是一個範例，將go編寫出來的遊戲(2048)，編譯成wasm格式，使其能在瀏覽器直接運行

> 特別申明: 2048的遊戲代碼都是從[ebiten/examples/2048/](https://github.com/hajimehoshi/ebiten/tree/e687865c8c8f1ce10e7bdf47548d8657c50d8cae/examples/2048)修改而得

## build

它會做以下事情:

1. 從`%GOROOT%/misc/wasm/wasm_exec.js`複製wasm_exec.js到[pages](pages/)資料夾
2. 編譯用Go寫的2048遊戲檔案放置到`pages/2048.wasm`

```
git clone https://github.com/CarsonSlovoka/twenty48.git
cd twenty48
go run build/main.go
```

## Run

您需要準備一個Server，來開啟[主頁文件](pages/index.html)

## 自行手動建立步驟

1. 編寫您用Go所寫的主程式，本範例是用[2048](src/)當成範例
2. 加入`wasm_exec.js`到您的[script](https://github.com/CarsonSlovoka/twenty48/blob/c0b028e/pages/index.html#L1)

   wasm_exec.js在GoRoot的資料夾中可以找到
   > %GOROOT%/misc/wasm/wasm_exec.js
3. 在您的script中連結[您的主wasm檔案](https://github.com/CarsonSlovoka/twenty48/blob/c0b028e/pages/index.html#L11)(第一點)

以上的東西，可以參考此[commit](https://github.com/CarsonSlovoka/twenty48/commit/c0b028e8a6304b9b7ef260f610797cf780097ac2)


## 參考資料
- [hajimehoshi/ebiten.examples/2048](https://github.com/hajimehoshi/ebiten/tree/e687865/examples/2048)
- [WebAssembly Compiling](https://ebiten.org/documents/webassembly.html)
- [github: hajimehoshi/wasmserve](https://github.com/hajimehoshi/wasmserve)
