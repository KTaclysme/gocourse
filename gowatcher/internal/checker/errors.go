package checker

import "fmt"

type UnreachebleURLError struct {
	URL string
	Err error
}

func (e *UnreachebleURLError) Error() string {
	return fmt.Sprintf("URL inaccessible : %s (%v)", e.URL, e.Err)
}

// func (e *UnreachebleURLError) Unwrap() error {
// 	return
// }
