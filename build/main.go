//go:build windows

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	PkgName        = "twenty48"
	OutputWasmName = "2048.wasm"
	Ldflags        = "-H=windowsgui -s -w"
)

var (
	SrcDir         string
	OutputWasmPath string
)

func init() {
	SrcDir, _ = filepath.Abs("./src")
	OutputWasmPath, _ = filepath.Abs(fmt.Sprintf("./pages/%s", OutputWasmName))
}

func main() {
	if err := os.Setenv("GOOS", "js"); err != nil { // The environment only works while the program is running; that is, it will disappear after the program stops.
		panic(err)
	}
	if err := os.Setenv("GOARCH", "wasm"); err != nil {
		panic(err)
	}
	fmt.Printf("GOOS:%s\n", os.Getenv("GOOS"))
	fmt.Printf("GOARCH:%s\n", os.Getenv("GOARCH"))

	dstWasmExecJsPath := filepath.Join(filepath.Dir(OutputWasmPath), "wasm_exec.js")
	if _, err := os.Stat(dstWasmExecJsPath); os.IsNotExist(err) {
		goRoot := os.Getenv("GoRoot") // env is case Insensitive
		if goRoot == "" {
			log.Fatal("Please check the env: 'GoROOT' exists.")
		}

		wasmExecJsPath := filepath.Join(goRoot, "misc/wasm/wasm_exec.js")
		if _, err := os.Stat(wasmExecJsPath); os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("File not found error: %q\n", wasmExecJsPath))
		}

		log.Printf("copy %q to %q\n", "%GOROOT%/misc/wasm/wasm_exec.js", dstWasmExecJsPath)
		srcBytes, err := ioutil.ReadFile(wasmExecJsPath)
		if err != nil {
			panic(err)
		}
		if err = ioutil.WriteFile(dstWasmExecJsPath, srcBytes, 0666); err != nil {
			panic(err)
		}
	}

	log.Println("go build...")
	cmd := exec.Command("go", "build", "-ldflags", Ldflags, "-o", OutputWasmPath, PkgName)
	cmd.Dir = SrcDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	log.Println("done!")
	_ = exec.Command("explorer", filepath.Dir(OutputWasmPath)).Start()
}
