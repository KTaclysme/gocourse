package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/KTaclysme/gowatcher/internal/config"
)

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string
	ErrMsg string
}

type CheckResult struct {
	Target config.InputTarget
	Status string
	Err    error
}

func CheckURL(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: time.Second * 3,
	}
	res, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{Target: target, Err: fmt.Errorf("Failed to fetch URL: %w", err)}
	}
	defer res.Body.Close()
	return CheckResult{Target: target, Status: res.Status}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   res.Target.Name,
		URL:    res.Target.URL,
		Owner:  res.Target.Owner,
		Status: res.Status, // Statut par défaut
	}

	if res.Err != nil {
		var unreachable *UnreachebleURLError
		if errors.As(res.Err, &unreachable) {
			report.Status = "Inaccessible"
			report.ErrMsg = fmt.Sprintf("Unreachable URL: %v", unreachable.Err)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("Erreur générique: %v", res.Err)
		}
	}

	return report
}
