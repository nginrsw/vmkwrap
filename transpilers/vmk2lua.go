package transpilers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// TranspileVMKtoLua converts all .vmk files in a directory (including subdirectories) to .lua
func TranspileVMKtoLua(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Pastikan ini bukan direktori dan file berakhiran .vmk
		if !info.IsDir() && strings.HasSuffix(path, ".vmk") {
			err := convertVMKtoLua(path)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", path, err)
			}
		}
		return nil
	})
}

// convertVMKtoLua handles individual file conversion (VMK → Lua)
func convertVMKtoLua(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Replace VMK keywords with Lua equivalents
	content = regexp.MustCompile(`\block\b`).ReplaceAll(content, []byte("local"))
	content = regexp.MustCompile(`\bfn\b`).ReplaceAll(content, []byte("function"))
	content = regexp.MustCompile(`\bstr\.`).ReplaceAll(content, []byte("string."))

	// Simpan sebagai .lua file
	luaFilePath := strings.Replace(filePath, ".vmk", ".lua", 1)
	err = os.WriteFile(luaFilePath, content, 0644)
	if err != nil {
		return err
	}

	// Hapus file .vmk yang lama
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Println("Converted:", filePath, "→", luaFilePath)
	return nil
}
