package interfaces

import (
	"github.com/RemmargorP/mjudge/models"
)

type Repos struct {
	Contests       ContestsRepo
	Languages      LanguagesRepo
	Problems       ProblemsRepo
	Submissions    SubmissionsRepo
	TestingResults TestingResultsRepo
	Users          UsersRepo
}

type ContestsRepo interface {
	GetById(models.Id) (*models.Contest, error)
	Get(limit int) ([]*models.Contest, error)
	Save(*models.Contest) error
}

type LanguagesRepo interface {
	GetById(models.Id) (*models.Language, error)
	Get() ([]*models.Language, error)
}

type ProblemsRepo interface {
	GetById(models.Id) (*models.Problem, error)
	Get(limit int) ([]*models.Problem, error)
	Save(*models.Problem) error
}

type SubmissionsRepo interface {
	GetById(models.Id) (*models.Submission, error)
	Get(limit int) ([]*models.Submission, error)
	Save(*models.Submission) error
}

type TestingResultsRepo interface {
	GetById(models.Id) (*models.TestingResult, error)
	Get(limit int) ([]*models.TestingResult, error)
	Save(*models.TestingResult) error
}

type UsersRepo interface {
	GetById(models.Id) (*models.User, error)
	Get(limit int) ([]*models.User, error)
	Save(*models.User) error
}
