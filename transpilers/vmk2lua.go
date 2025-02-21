package transpilers

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Transpile VMK to Lua changes the .vmk file to .lua
func TranspileVMKtoLua(filePath string) (string, error) {
	if !strings.HasSuffix(filePath, ".vmk") {
		return "", fmt.Errorf("invalid file type: %s", filePath)
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Change VMK Keywords with Lua
	content = regexp.MustCompile(`\block\b`).ReplaceAll(content, []byte("local"))
	content = regexp.MustCompile(`\bfn\b`).ReplaceAll(content, []byte("function"))
	content = regexp.MustCompile(`\bstr.\b`).ReplaceAll(content, []byte("string."))

	// Save as a .lua file
	luaFilePath := strings.Replace(filePath, ".vmk", ".lua", 1)
	err = ioutil.WriteFile(luaFilePath, content, 0644)
	if err != nil {
		return "", err
	}

	fmt.Println("Transpile VMK → Lua complete!")
	return luaFilePath, nil
}
