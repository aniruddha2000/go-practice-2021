package database

import (
	"encoding/json"
	"os"
	"strings"
)

type User struct {
	Username    string        `json:"username"`
	Balance     int64         `json:"balance"`
	Transaction []Transaction `json:"transaction"`
}

type Transaction struct {
	Amount    int64  `json:"amount"`
	Type      string `json:"string"`
	Narration string `json:"narration"`
}

func getUsers() ([]User, error) {
	var users []User
	data, err := os.ReadFile("database/db.json")
	if err == nil {
		json.Unmarshal(data, &users)
	}
	return users, err
}

func updateDB(data []User) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	os.WriteFile("database/db.json", bytes, 0644)
}

func FindUser(username string) (*User, error) {
	users, err := getUsers()
	if err == nil {
		for i := range users {
			user := users[i]
			if strings.EqualFold(user.Username, username) {
				return &user, nil
			}
		}
	}

	return nil, err
}

func FindOrCreateUser(username string) (*User, error) {
	user, err := FindUser(username)
	if user == nil {
		var newUser = User{
			Username:    strings.ToLower(username),
			Balance:     0,
			Transaction: []Transaction{},
		}
		users, err := getUsers()
		if err == nil {
			users = append(users, newUser)
			updateDB(users)
		}
		return &newUser, err
	}
	return user, err
}

func UpdateUser(user *User) {
	users, err := getUsers()
	if err == nil {
		for i := range users {
			if strings.EqualFold(users[i].Username, user.Username) {
				users[i] = *user
			}
		}
		updateDB(users)
	}
}
