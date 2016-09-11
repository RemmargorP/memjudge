package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/RemmargorP/mjudge/extensions/db"
	"github.com/RemmargorP/mjudge/interfaces"
	"github.com/RemmargorP/mjudge/models"
)

func main() {
	logfile, err := os.OpenFile("mjudgectl.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)

	var rawConfig []byte
	rawConfig, err = ioutil.ReadFile("config.json")
	if err != nil {
		logger.Fatalln(err)
	}

	var Config struct {
		Encryption map[string]interface{}
		DB_MongoDB map[string]interface{}
	}
	err = json.Unmarshal(rawConfig, &Config)
	if err != nil {
		logger.Fatalln(err)
	}

	var dbConfig []byte
	dbConfig, err = json.Marshal(Config.DB_MongoDB)

	var repos *interfaces.Repos
	repos, err = db.DBReposFromConfig(dbConfig)
	logger.Printf("%+v\n", repos)
	logger.Printf("%+v\n", repos.Contests.Save(&models.Contest{"mda kek", []models.Id{}, []models.Id{}, []models.Id{}}))
	if err != nil {
		logger.Fatalln(err.Error())
	}
	logger.Println(repos)
}
