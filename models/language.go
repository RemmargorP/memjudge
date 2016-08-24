package models

type Language struct {
	Id        Id // Common names like "G++ 5.1 C++ 11", "Java 8.??", "Go 1.6", "Python 3.4.3", etc.
	ShortName Id // cpp, python, go, java (For readability)
	LanguageInterface
}

type LanguageInterface interface {
	Process(*Submission, *Problem) (*TestingResult, error)
}
