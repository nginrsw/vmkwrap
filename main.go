package main

import (
	"fmt"
	"os"
	"os/exec"

	"vmkwrap/transpilers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vmkwrap <file.vmk>")
		return
	}

	filePath := os.Args[1]

	// 1. Transpile VMK → Lua
	luaFile, err := transpilers.TranspileVMKtoLua(filePath)
	if err != nil {
		fmt.Println("Error transpiling:", err)
		return
	}

	// 2. Run the .lua file on luajit
	fmt.Println("Running LuaJIT...")
	if err := runLuaJIT(luaFile); err != nil {
		fmt.Println("Error running LuaJIT:", err)
		return
	}

	// 3. Return the .lua file to .vmk
	if err := transpilers.ReverseLuaToVMK(luaFile); err != nil {
		fmt.Println("Error reverse transpiling:", err)
		return
	}

	fmt.Println("Process complete!")
}

// Run Luajit with a transpiled file
func runLuaJIT(luaFile string) error {
	luajitPath, err := exec.LookPath("luajit")
	if err != nil {
		return fmt.Errorf("LuaJIT not found. Make sure it is installed and accessible.")
	}

	cmd := exec.Command(luajitPath, luaFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
