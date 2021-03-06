package main

/*
 * This file serves as a database mock. We just declare some constants here for
 * later use.
 */

type user struct {
	Name string
	Age  int
}

var users = []user{
	{"Viktor", 21},
	{"Marta", 35},
	{"Eugene", 26},
	{"George", 43},
	{"Antony", 48},
}

func findUser(users []user, name string) *user {
	for _, u := range users {
		if u.Name == name {
			return &u
		}
	}
	return nil
}
