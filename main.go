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
	// Gunakan "." sebagai root directory agar semua file di direktori saat ini & subfolder ikut diproses
	rootDir := "."

	// 1. Transpile VMK → Lua untuk semua file
	fmt.Println("Transpiling VMK → Lua...")
	err := transpilers.TranspileVMKtoLua(rootDir)
	if err != nil {
		fmt.Println("Error transpiling:", err)
		return
	}

	// 2. Jalankan semua file .lua yang ditemukan di rootDir
	fmt.Println("Running LuaJIT...")
	err = runLuaJIT(rootDir)
	if err != nil {
		fmt.Println("Error running LuaJIT:", err)
		return
	}

	// 3. Reverse transpile Lua → VMK untuk semua file
	fmt.Println("Reversing Lua → VMK...")
	err = transpilers.ReverseLuaToVMK(rootDir)
	if err != nil {
		fmt.Println("Error reverse transpiling:", err)
		return
	}

	fmt.Println("Process complete!")
}

// runLuaJIT menjalankan semua file .lua dalam rootDir menggunakan LuaJIT
func runLuaJIT(rootDir string) error {
	luajitPath, err := exec.LookPath("luajit")
	if err != nil {
		return fmt.Errorf("LuaJIT not found. Make sure it is installed and accessible.")
	}

	// Cari semua file .lua dalam rootDir dan jalankan satu per satu
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
