package checker

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckURL(url string) CheckResult {
	client := http.Client{
		Timeout: time.Second * 3,
	}
	res, err := client.Get(url)
	if err != nil {
		return CheckResult{Target: url, Err: fmt.Errorf("Failed to fetch URL: %w", err)}
	}
	defer res.Body.Close()
	result := CheckResult{
		Target: url,
		Err:    &UnreachebleURLError{URL: url, Err: err},
	}
	return result
}
