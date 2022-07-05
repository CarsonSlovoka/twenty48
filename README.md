## twenty48

此項目僅是一個範例，將go編寫出來的遊戲(2048)，編譯成wasm格式，使其能在瀏覽器直接運行

> 特別申明: 2048的遊戲代碼都是從[ebiten/examples/2048/](https://github.com/hajimehoshi/ebiten/tree/e687865c8c8f1ce10e7bdf47548d8657c50d8cae/examples/2048)修改而得

## Demo

我把成品放到gh-pages之中，請點擊連結來觀看範例

> https://carsonslovoka.github.io/twenty48/

### [在Github Page套用PWA的注意事項](https://stackoverflow.com/a/72867561/9935654)

因為Github的根路徑是使用者名稱

> https://{github_username}.github.io/{repository}

所以在設定PWA的時候，Service Workers也要做調整，要加上[{repository}](https://github.com/CarsonSlovoka/twenty48/blob/81d0440244e95ce84e7a5f177477071fe024d0f5/docs/sw.js#L8-L14)

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

或者您可以用我們所建立的簡單Server來啟動

> go run main/server

其中[main](https://github.com/CarsonSlovoka/twenty48/blob/a4b5c29/go.mod#L1)是我們根路徑所定義的package名稱

至於[server](https://github.com/CarsonSlovoka/twenty48/blob/a4b5c297c1bf43cd0c004e11134579461b71c91e/server/main.go#L1)則是我們實際包含main的go代碼

## 自行手動建立步驟

1. 編寫您用Go所寫的主程式，本範例是用[2048](src/)當成範例
2. 加入`wasm_exec.js`到您的[script](https://github.com/CarsonSlovoka/twenty48/blob/c0b028e/pages/index.html#L1)

   wasm_exec.js在GoRoot的資料夾中可以找到
   > %GOROOT%/misc/wasm/wasm_exec.js
3. 在您的script中連結[您的主wasm檔案](https://github.com/CarsonSlovoka/twenty48/blob/c0b028e/pages/index.html#L11)(第一點)

以上的東西，可以參考此[commit](https://github.com/CarsonSlovoka/twenty48/commit/c0b028e8a6304b9b7ef260f610797cf780097ac2)

## Publish to [replit](https://replit.com/)

請手動編譯，並複製wasm_exec.js到pages資料夾內

請參考[指令](src/README.md)

同時請設定環境變[Debug改為false](https://github.com/CarsonSlovoka/twenty48/blob/a4b5c29/server/main.go#L19)

## 參考資料
- [hajimehoshi/ebiten.examples/2048](https://github.com/hajimehoshi/ebiten/tree/e687865/examples/2048)
- [WebAssembly Compiling](https://ebiten.org/documents/webassembly.html)
- [github: hajimehoshi/wasmserve](https://github.com/hajimehoshi/wasmserve)
- [Uniquely identifying PWAs with the web app manifest id property](https://developer.chrome.com/blog/pwa-manifest-id/?utm_source=devtools)
