package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/RemmargorP/mjudge/extensions/db"
	"github.com/RemmargorP/mjudge/interfaces"
	"github.com/RemmargorP/mjudge/logic"
	"github.com/RemmargorP/mjudge/models"

	"github.com/howeyc/gopass"
)

func main() {
	var rawConfig []byte
	var err error
	rawConfig, err = ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	var Config struct {
		Encryption map[string]string
		DB_MongoDB map[string]interface{}
	}
	err = json.Unmarshal(rawConfig, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	var dbConfig []byte
	dbConfig, err = json.Marshal(Config.DB_MongoDB)

	var repos *interfaces.Repos
	repos, err = db.DBReposFromConfig(dbConfig)
	repos.Contests.Get(1)
	fmt.Printf("Connected to the DB.\n")

	var login, password string
	fmt.Printf("Login [admin]: ")
	scanner := bufio.NewReader(os.Stdin)
	login, _ = scanner.ReadString('\n')
	if len(login) > 0 {
		login = login[:len(login)-1]
	}
	if len(login) == 0 {
		login = "admin"
	}
	fmt.Printf("Password: ")
	tmp, _ := gopass.GetPasswd()
	password = string(tmp)

	fmt.Printf("Email: ")
	email, _ := scanner.ReadString('\n')
	if len(email) > 0 {
		email = email[:len(email)-1]
	}

	var user *models.User
	user, err = repos.Users.GetByLogin(login)
	if user != nil {
		log.Fatalln("User with such login already exists")
	}

	user, err = repos.Users.GetByEmail(email)
	if user != nil {
		log.Fatalln("User with such email already exists")
	}

	err = repos.Users.Save(&models.User{
		Id:            logic.GenerateId(),
		Login:         login,
		Email:         email,
		Password:      logic.Encrypt(password, Config.Encryption["PasswordLeftSalt"], Config.Encryption["PasswordRightSalt"]),
		Administrator: true})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Success!")
}
