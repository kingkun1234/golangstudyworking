package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := Customer{"king", "kingkun", "1234567890", 123, 23}
	str, _ := json.Marshal(user)
	fmt.Println(string(str))
	var customer Customer
	json.Unmarshal(str, &customer)
	fmt.Println(customer)
}

type Customer struct {
	UserName     string `json:"userName"`
	UserAccount  string `json:"userAccount"`
	UserMobile   string `json:"userMobile"`
	VotesNumber  int    `json:"votesNumber"`
	ActivityRank int    `json:"activityRank"`
}
