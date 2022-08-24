package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	Stats["create"] += 1
	fmt.Println("Creating user", user)
}

func UpdateUser(user string) {
	Stats["update"] += 1
	fmt.Println("Updating user", user)
}

func PurgeStats() {
	delete(Stats, "create")
	delete(Stats, "update")
}
