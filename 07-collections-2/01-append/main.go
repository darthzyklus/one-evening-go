package main

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
}

func AddUser(user string) {
	users = append(users, user)
}
