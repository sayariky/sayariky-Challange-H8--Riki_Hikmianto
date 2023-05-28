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

func Register() {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

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

	err := config.DB.Create(&user).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Pesanan berhasil dibuat")

	time.Sleep(3 * time.Second)
	Login("")
}
