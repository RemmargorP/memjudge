package models

// Testing Result - total result based on each test result
type TestingResult struct {
	Id             Id
	Submission     Id
	Status         int8
	TestResults    []Test
	TestedBy       Id //Invoker Id
	FailedTest     int32
	MaxTimeElapsed float32
	MaxMemoryUsed  int32
	Reason         string
}

const (
	TestingResult_Ok = iota
	TestingResult_WrongAnswer
	TestingResult_CompilationError
	TestingResult_TimeLimitExceeded
	TestingResult_MemoryLimitExceeded
	TestingResult_RuntimeError
	TestingResult_TestingFailed
)
