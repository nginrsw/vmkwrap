package transpilers

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Reverse Lua to VMK returns the .lua file to .vmk
func ReverseLuaToVMK(filePath string) error {
	if !strings.HasSuffix(filePath, ".lua") {
		return fmt.Errorf("invalid file type: %s", filePath)
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// return the keyword lua to VMK
	content = regexp.MustCompile(`\blocal\b`).ReplaceAll(content, []byte("lock"))
	content = regexp.MustCompile(`\bfunction\b`).ReplaceAll(content, []byte("fn"))
	content = regexp.MustCompile(`\bstring.\b`).ReplaceAll(content, []byte("str."))

	// Save as a .vmk file
	vmkFilePath := strings.Replace(filePath, ".lua", ".vmk", 1)
	err = ioutil.WriteFile(vmkFilePath, content, 0644)
	if err != nil {
		return err
	}

	// Delete the .lua file after completion of reverse transpile
	os.Remove(filePath)

	fmt.Println("Reverse Lua → VMK complete!")
	return nil
}
