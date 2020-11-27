package main

import "fmt"

var db map[string]string

// User struct
type User struct {
	uid int
	id  string
	pw  string
}

//SignUp for signup
func SignUp() (u User) {
	var id string
	var pw string
	fmt.Printf("ID:")
	fmt.Scanf("%s", &id)
	fmt.Printf("PW:")
	fmt.Scanf("%s", &pw)

	u = User{id: id, pw: pw}

	fmt.Println("account created.\n", u)
	db[id] = pw

	return

}

// SignIn is login
func SignIn() bool {
	var id string
	var pw string
	fmt.Println("Log In")
	fmt.Printf("ID:")
	fmt.Scanf("%s", &id)
	fmt.Printf("PW:")
	fmt.Scanf("%s", &pw)

	val, exist := db[id]
	if !exist {
		fmt.Println("ID doesn`t exist")
		return false
	} else if val != pw {
		fmt.Println("wrong password.")
		return false
	}
	fmt.Println("Success!")
	return true

}

func main() {
	var i int
	login := false

	db = map[string]string{}
	for login == false {
		fmt.Printf("1: SignUp\n2: SignIn")
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			SignUp()
		case 2:
			login = SignIn()

		}
	}
	fmt.Println("Connected")

}
