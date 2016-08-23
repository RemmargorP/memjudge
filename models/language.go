package models

import (
	"errors"
)

type Language struct {
	Id        Id // Not UUID! Common names like "G++ 5.1 C++ 11", "Java 8.??", "Go 1.6", "Python 3.4.3", etc.
	ShortName Id // cpp, python, go, java (For readability)
}

func (lang *Language) Process(sandboxPath string) (TestingResult, error) {
	return errors.New("Not implemented")
}
