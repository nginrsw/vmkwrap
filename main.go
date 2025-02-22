package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"vmkwrap/transpilers"
)

func main() {
	// Use "." As a root directory, all files in the current directory & subfolder are processed
	rootDir := "."

	// 1. Transpile VMK → Lua for all files
	fmt.Println("Transpiling VMK → Lua...")
	err := transpilers.TranspileVMKtoLua(rootDir)
	if err != nil {
		fmt.Println("Error transpiling:", err)
		return
	}

	// 2. Run all. Lua files found in rootdir
	fmt.Println("Running LuaJIT...")
	err = runLuaJIT(rootDir)
	if err != nil {
		fmt.Println("Error running LuaJIT:", err)
		return
	}

	// 3. Reverse Transpile Lua → VMK for all files
	fmt.Println("Reversing Lua → VMK...")
	err = transpilers.ReverseLuaToVMK(rootDir)
	if err != nil {
		fmt.Println("Error reverse transpiling:", err)
		return
	}

	fmt.Println("Process complete!")
}

// runLuaJIT runs all .lua files in rootDir using LuaJIT
func runLuaJIT(rootDir string) error {
	luajitPath, err := exec.LookPath("luajit")
	if err != nil {
		return fmt.Errorf("LuaJIT not found. Make sure it is installed and accessible.")
	}

	// Search for all .lua files in rootDir and run them one by one
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".lua") {
			cmd := exec.Command(luajitPath, path)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running %s: %v\n", path, err)
			}
		}
		return nil
	})

	return err
}
