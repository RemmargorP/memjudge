package db

// Mongo DB

import (
	"encoding/json"

	"github.com/RemmargorP/mjudge/interfaces"
	_ "github.com/RemmargorP/mjudge/logic"
	"github.com/RemmargorP/mjudge/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func DBReposFromConfig(config []byte) (*interfaces.Repos, error) {
	var DBConfig struct {
		URL      string
		Login    string
		Password string
		Database string
		Prefix   string
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
		Contests:       DBContestsRepo{db, DBConfig.Prefix},
		Languages:      DBLanguagesRepo{},
		Problems:       DBProblemsRepo{db, DBConfig.Prefix},
		Submissions:    DBSubmissionsRepo{db, DBConfig.Prefix},
		TestingResults: DBTestingResultsRepo{db, DBConfig.Prefix},
		Users:          DBUsersRepo{db, DBConfig.Prefix},
	}, nil
}

// Contests
type DBContestsRepo struct {
	DB     *mgo.Database
	Prefix string
}

func (r DBContestsRepo) GetById(id models.Id) (*models.Contest, error) {
	var res *models.Contest
	err := r.DB.C(r.Prefix + "contests").FindId(id).One(&res)
	return res, err
}
func (r DBContestsRepo) Get(limit int) ([]*models.Contest, error) {
	var res []*models.Contest
	var err error
	if limit == -1 {
		err = r.DB.C(r.Prefix + "contests").Find(bson.M{}).All(&res)
	} else {
		err = r.DB.C(r.Prefix + "contests").Find(bson.M{}).Limit(limit).All(&res)
	}
	return res, err
}
func (r DBContestsRepo) Save(contest *models.Contest) error {
	_, err := r.DB.C(r.Prefix+"contests").Upsert(bson.M{"_id": contest.Id}, bson.M{
		"$set": contest,
	})
	return err
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
	DB     *mgo.Database
	Prefix string
}

func (r DBProblemsRepo) GetById(id models.Id) (*models.Problem, error) {
	var res *models.Problem
	err := r.DB.C(r.Prefix + "problems").FindId(id).One(&res)
	return res, err
}
func (r DBProblemsRepo) Get(limit int) ([]*models.Problem, error) {
	var res []*models.Problem
	var err error
	if limit == -1 {
		err = r.DB.C(r.Prefix + "problems").Find(bson.M{}).All(&res)
	} else {
		err = r.DB.C(r.Prefix + "problems").Find(bson.M{}).Limit(limit).All(&res)
	}
	return res, err
}
func (r DBProblemsRepo) Save(problem *models.Problem) error {
	_, err := r.DB.C(r.Prefix+"problems").Upsert(bson.M{"_id": problem.Id}, bson.M{
		"$set": problem,
	})
	return err
}

// Submissions
type DBSubmissionsRepo struct {
	DB     *mgo.Database
	Prefix string
}

func (r DBSubmissionsRepo) GetById(id models.Id) (*models.Submission, error) {
	var res *models.Submission
	err := r.DB.C(r.Prefix + "submissions").FindId(id).One(&res)
	return res, err
}
func (r DBSubmissionsRepo) Get(limit int) ([]*models.Submission, error) {
	var res []*models.Submission
	var err error
	if limit == -1 {
		err = r.DB.C(r.Prefix + "submissions").Find(bson.M{}).All(&res)
	} else {
		err = r.DB.C(r.Prefix + "submissions").Find(bson.M{}).Limit(limit).All(&res)
	}
	return res, err
}
func (r DBSubmissionsRepo) Save(submission *models.Submission) error {
	_, err := r.DB.C(r.Prefix+"submissions").Upsert(bson.M{"_id": submission.Id}, bson.M{
		"$set": submission,
	})
	return err
}

// Testing Results
type DBTestingResultsRepo struct {
	DB     *mgo.Database
	Prefix string
}

func (r DBTestingResultsRepo) GetById(id models.Id) (*models.TestingResult, error) {
	var res *models.TestingResult
	err := r.DB.C(r.Prefix + "testingresults").FindId(id).One(&res)
	return res, err
}
func (r DBTestingResultsRepo) Get(limit int) ([]*models.TestingResult, error) {
	var res []*models.TestingResult
	var err error
	if limit == -1 {
		err = r.DB.C(r.Prefix + "testingresults").Find(bson.M{}).All(&res)
	} else {
		err = r.DB.C(r.Prefix + "testingresults").Find(bson.M{}).Limit(limit).All(&res)
	}
	return res, err
}
func (r DBTestingResultsRepo) Save(tr *models.TestingResult) error {
	_, err := r.DB.C(r.Prefix+"testingresults").Upsert(bson.M{"_id": tr.Id}, bson.M{
		"$set": tr,
	})
	return err
}

// Submissions
type DBUsersRepo struct {
	DB     *mgo.Database
	Prefix string
}

func (r DBUsersRepo) GetById(id models.Id) (*models.User, error) {
	var res *models.User
	err := r.DB.C(r.Prefix + "users").FindId(id).One(&res)
	return res, err
}
func (r DBUsersRepo) GetByLogin(login string) (*models.User, error) {
	var res *models.User
	err := r.DB.C(r.Prefix + "users").Find(bson.M{"login": login}).One(&res)
	return res, err
}
func (r DBUsersRepo) GetByEmail(email string) (*models.User, error) {
	var res *models.User
	err := r.DB.C(r.Prefix + "users").Find(bson.M{"email": email}).One(&res)
	return res, err
}

func (r DBUsersRepo) Get(limit int) ([]*models.User, error) {
	var res []*models.User
	var err error
	if limit == -1 {
		err = r.DB.C(r.Prefix + "users").Find(bson.M{}).All(&res)
	} else {
		err = r.DB.C(r.Prefix + "users").Find(bson.M{}).Limit(limit).All(&res)
	}
	return res, err
}
func (r DBUsersRepo) Save(user *models.User) error {
	_, err := r.DB.C(r.Prefix+"users").Upsert(bson.M{"_id": user.Id}, bson.M{
		"$set": user,
	})
	return err
}
