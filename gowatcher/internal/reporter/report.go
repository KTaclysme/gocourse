package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/KTaclysme/gowatcher/internal/checker"
)

func ExportResultatsToJsonFile(filePath string, results []checker.ReportEntry) error {
	data, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return fmt.Errorf("impossible d'encoder le ficher %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire le rapport json dans le ficher %s: %w", filePath, err)
	}
	return nil
}
