package main

import (
	"fmt"
)

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {

		// return fmt.Errorf("fmt.Errorf...")
		// return errors.New("errors.New...")
		return PasswordError{
			Len:        len(password),
			RequireLen: 8,
		}
	}
	// ...
	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d ReqLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입 완료.")
	}
}
