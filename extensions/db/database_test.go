package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/RemmargorP/mjudge/interfaces"
	"github.com/RemmargorP/mjudge/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var repos *interfaces.Repos
var db *mgo.Database

func TestMain(m *testing.M) {
	rawConfig, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		log.Fatalln(err)
	}

	var Config struct {
		Encryption map[string]interface{}
		DB_MongoDB map[string]interface{}
	}
	err = json.Unmarshal(rawConfig, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	var dbConfig []byte
	dbConfig, err = json.Marshal(Config.DB_MongoDB)

	repos, err = DBReposFromConfig(dbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	var DBConfig struct {
		URL      string
		Login    string
		Password string
		Database string
	}

	if err = json.Unmarshal(dbConfig, &DBConfig); err != nil {
		log.Fatalln(err)
	}

	var session *mgo.Session
	session, err = mgo.Dial(DBConfig.URL)
	if err != nil {
		log.Fatalln(err)
	}

	session.SetSafe(&mgo.Safe{})

	db = session.DB(DBConfig.Database)
	err = db.Login(DBConfig.Login, DBConfig.Password)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestContestsGetById(t *testing.T) {
	defer db.C("contests").DropCollection()
	db.C("contests").Insert(bson.M{"_id": "abacaba"})
	res, err := repos.Contests.GetById(models.Id("abacaba"))
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Fatal("nil result")
	}
	if res.Id != "abacaba" {
		t.Fatalf("ids differ %+v", res)
	}
}

func TestContestsGet(t *testing.T) {
	defer db.C("contests").DropCollection()
	for i := 0; i < 10; i++ {
		db.C("contests").Insert(models.Contest{Id: models.Id(i)})
	}
	res, err := repos.Contests.Get(-1)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 10 {
		t.Fatal("Lenghts differ")
	}
	res, err = repos.Contests.Get(3)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 3 {
		t.Fatal("Lenghts diffed")
	}
}
func TestContestsSave(t *testing.T) {
	defer db.C("contests").DropCollection()
	err := repos.Contests.Save(&models.Contest{
		Id:       "lel",
		Problems: []models.Id{"aaas", "asdas", "ioi", "хых"},
		Rights:   []models.Id{"nope", "nope", "mid", "mda kek "},
	})
	if err != nil {
		t.Fatal(err)
	}
	var res *models.Contest
	res, err = repos.Contests.GetById("lel")
	if res == nil {
		t.Fatal("not found")
	}
}
