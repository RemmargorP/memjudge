package models

// Testing Result - total result based on each test result
type TestingResult struct {
	Id             Id      `bson:"_id"`
	Submission     Id      `bson:"submission"`
	Status         int8    `bson:"status"`
	TestResults    []Test  `bson:"tests"`
	TestedBy       Id      `bson:"tested_by"` //Invoker Id
	FailedTest     int32   `bson:"failed_test"`
	MaxTimeElapsed float32 `bson:"max_time_elapsed"`
	MaxMemoryUsed  int32   `bson:"max_memory_used"`
	Reason         string  `bson:"reason"`
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
