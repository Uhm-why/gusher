package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Uhm-why/gusher/fileio"
)

func SaveFileInfo(fi fileio.FileInfo) error {

	baseName := filepath.Base(fi.Name)

	outputDir := fmt.Sprintf("%v_gusherFileInfo", fi.Name)

	if err := os.MkdirAll(outputDir, 0777); err != nil {
		return fmt.Errorf("failed to create directory: %v", outputDir)
	}

	jsonPath := filepath.Join(outputDir, baseName+".gusher.json")

	data, err := json.MarshalIndent(fi, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal FileInfo to JSON: %v", err)
	}

	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON to %s: %v", jsonPath, err)
	}

	return nil
}

func LoadFileInfo(fname string) {

}
