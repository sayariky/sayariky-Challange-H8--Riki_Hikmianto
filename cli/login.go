package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
	"time"
)

func Login(message string) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	if message != "" {
		fmt.Println("Username atau password salah")
		time.Sleep(3 * time.Second)
		fmt.Println("silahkan login kembali")
	}

	var username string
	var password string
	fmt.Println("Masukan username anda")
	username, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukan password anda")
	password, _ = consoleReader.ReadString('\n')

	user := entity.User{
		Username: username,
		Password: password,
	}

	err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if err != nil {
		Login(err.Error())
	}

	fmt.Println("Login berhasil")

	time.Sleep(3 * time.Second)
	ListProduct(user.Username)
}
