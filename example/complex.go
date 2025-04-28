package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"

	"github.com/Nadim147c/go-tree"
)

type User struct {
	Username string
	Nickname string
	Age      uint
}

func main() {
	type tree map[string]any
	data := map[string]any{
		"contents": tree{
			"list": tree{
				"users": map[string][]any{
					"bob454":  {"bob454", "Bob", "19"},
					"alice13": {"alice13", "Alice", "25"},
					"nick23":  {"nick23", "Nick", "25"},
				},
			},
		},
	}

	// Find the users map
	usersMap, _ := gotree.Find(data, func(n gotree.Node) bool {
		return n.Key == "users"
	})

	// get all the users
	usersData, _ := gotree.Traverse(usersMap, func(n gotree.Node) bool {
		return n.Value.Kind() == reflect.Slice
	})

	var users []User

	for _, ud := range usersData {
		u := ud.([]any)
		var user User
		user.Username = u[0].(string)
		user.Nickname = u[1].(string)
		age, _ := strconv.Atoi(u[2].(string))
		user.Age = uint(age)

		users = append(users, user)
	}

	json.NewEncoder(os.Stdout).Encode(users)
	// [
	//	{ "Username": "bob454", "Nickname": "Bob", "Age": 19 },
	//	{ "Username": "alice13", "Nickname": "Alice", "Age": 25 },
	//	{ "Username": "nick23", "Nickname": "Nick", "Age": 25 }
	// ]
}
