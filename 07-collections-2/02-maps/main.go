package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	Stats["create"]++
	fmt.Println("Creating user", user)
}

func UpdateUser(user string) {
	Stats["update"]++
	fmt.Println("Updating user", user)
}

func PurgeStats() {
	Stats["create"] = 0
	Stats["update"] = 0
}
