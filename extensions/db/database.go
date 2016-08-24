package db

// Mongo DB

import (
	"encoding/json"

	"github.com/RemmargorP/mjudge/interfaces"
	_ "github.com/RemmargorP/mjudge/logic"
	"github.com/RemmargorP/mjudge/models"

	"gopkg.in/mgo.v2"
)

func DBReposFromConfig(config []byte) (*interfaces.Repos, error) {
	var DBConfig struct {
		URL      string
		Login    string
		Password string
		Database string
	}

	if err := json.Unmarshal(config, &DBConfig); err != nil {
		return nil, err
	}

	session, err := mgo.Dial(DBConfig.URL)
	if err != nil {
		return nil, err
	}

	session.SetSafe(&mgo.Safe{})

	db := session.DB(DBConfig.Database)
	err = db.Login(DBConfig.Login, DBConfig.Password)
	if err != nil {
		return nil, err
	}

	return &interfaces.Repos{
		Contests:       DBContestsRepo{db},
		Languages:      DBLanguagesRepo{},
		Problems:       DBProblemsRepo{db},
		Submissions:    DBSubmissionsRepo{db},
		TestingResults: DBTestingResultsRepo{db},
		Users:          DBUsersRepo{db},
	}, nil
}

// Contests
type DBContestsRepo struct {
	DB *mgo.Database
}

func (r DBContestsRepo) GetById(models.Id) (*models.Contest, error) {
	return nil, nil
}
func (r DBContestsRepo) GetByData(sample map[string]interface{}) (*models.Contest, error) {
	return nil, nil
}
func (r DBContestsRepo) Get(limit int) ([]*models.Contest, error) {
	return nil, nil
}
func (r DBContestsRepo) Save(*models.Contest) error {
	return nil
}

// Languages
type DBLanguagesRepo struct{}

func (r DBLanguagesRepo) GetById(models.Id) (*models.Language, error) {
	return nil, nil
}
func (r DBLanguagesRepo) Get() ([]*models.Language, error) {
	return nil, nil
}

// Problems
type DBProblemsRepo struct {
	DB *mgo.Database
}

func (r DBProblemsRepo) GetById(models.Id) (*models.Problem, error) {
	return nil, nil
}
func (r DBProblemsRepo) GetByData(sample map[string]interface{}) (*models.Problem, error) {
	return nil, nil
}
func (r DBProblemsRepo) Get(limit int) ([]*models.Problem, error) {
	return nil, nil
}
func (r DBProblemsRepo) Save(*models.Problem) error {
	return nil
}

// Submissions
type DBSubmissionsRepo struct {
	DB *mgo.Database
}

func (r DBSubmissionsRepo) GetById(models.Id) (*models.Submission, error) {
	return nil, nil
}
func (r DBSubmissionsRepo) GetByData(sample map[string]interface{}) (*models.Submission, error) {
	return nil, nil
}
func (r DBSubmissionsRepo) Get(limit int) ([]*models.Submission, error) {
	return nil, nil
}
func (r DBSubmissionsRepo) Save(*models.Submission) error {
	return nil
}

// Testing Results
type DBTestingResultsRepo struct {
	DB *mgo.Database
}

func (r DBTestingResultsRepo) GetById(models.Id) (*models.TestingResult, error) {
	return nil, nil
}
func (r DBTestingResultsRepo) GetByData(sample map[string]interface{}) (*models.TestingResult, error) {
	return nil, nil
}
func (r DBTestingResultsRepo) Get(limit int) ([]*models.TestingResult, error) {
	return nil, nil
}
func (r DBTestingResultsRepo) Save(*models.TestingResult) error {
	return nil
}

// Submissions
type DBUsersRepo struct {
	DB *mgo.Database
}

func (r DBUsersRepo) GetById(models.Id) (*models.User, error) {
	return nil, nil
}
func (r DBUsersRepo) GetByData(sample map[string]interface{}) (*models.User, error) {
	return nil, nil
}
func (r DBUsersRepo) Get(limit int) ([]*models.User, error) {
	return nil, nil
}
func (r DBUsersRepo) Save(*models.User) error {
	return nil
}
