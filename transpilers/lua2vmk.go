package transpilers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ReverseLuaToVMK converts all .lua files in a directory (including subdirectories) to .vmk
func ReverseLuaToVMK(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Pastikan ini bukan direktori dan file berakhiran .lua
		if !info.IsDir() && strings.HasSuffix(path, ".lua") {
			err := reverseConvertFile(path)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", path, err)
			}
		}
		return nil
	})
}

// reverseConvertFile handles individual file conversion (Lua → VMK)
func reverseConvertFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Replace Lua keywords with VMK equivalents
	content = regexp.MustCompile(`\blocal\b`).ReplaceAll(content, []byte("lock"))
	content = regexp.MustCompile(`\bfunction\b`).ReplaceAll(content, []byte("fn"))
	content = regexp.MustCompile(`\bstring\.`).ReplaceAll(content, []byte("str."))

	// Simpan sebagai .vmk file
	vmkFilePath := strings.Replace(filePath, ".lua", ".vmk", 1)
	err = os.WriteFile(vmkFilePath, content, 0644)
	if err != nil {
		return err
	}

	// Hapus file .lua yang lama
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Println("Converted:", filePath, "→", vmkFilePath)
	return nil
}
