package main

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-finance-api/internal/config"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/db"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/types"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util"
	"github.com/leetcode-golang-classroom/golang-finance-api/pkg/password"
)

func main() {
	dbInstance, err := db.Connect(config.AppConfig.DbURL)
	if err != nil {
		util.FailOnError(err, "failed to connect")
	}
	passwdHandler := password.Handler{}
	passwd, err := passwdHandler.HashPassword(config.AppConfig.DefaultAdminPassword)
	if err != nil {
		log.Fatalf("error generating password: %v", err)
	}
	users := []types.User{
		{
			Username: "admin",
			Password: passwd,
		},
	}
	_, err = dbInstance.NamedExec(
		`INSERT INTO users (username, password)
		VALUES (:username, :password);
		`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}
	log.Printf("Successfully inserted user: %v\n", users)
}
