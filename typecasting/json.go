package typecasting

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id     string   `json:"id"`
	Age    int32    `json:"age"`
	Skills []string `json:"languages"`
}

func castUserToJson(u *User) string {
	s, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	return string(s)
}

func castJsonToUser(s string) User {
	var u User
	if err := json.Unmarshal([]byte(s), &u); err != nil {
		panic(err)
	}

	return u
}

func Test() {
	s := castUserToJson(&User{
		Id:     "1",
		Age:    31,
		Skills: []string{"javascript", "react", "go"},
	})
	fmt.Println(s)

	u := castJsonToUser(s)
	fmt.Println(u.Id)
	fmt.Println(u.Age)
	fmt.Println(u.Skills)
}
