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

func ListProduct(name string) {
	helpers.ClearScreen()
	consolReader := bufio.NewReader(os.Stdin)

	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("selamat datang", name)
	fmt.Println("---List Product---")

	for _, product := range products {
		product.PrintDetail()
	}

	var input string
	fmt.Println("Masukan Id product yang ingin dipesan")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consolReader.ReadString('\n')

	switch input {
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		OrderProduct(input)
	}
}

func OrderProduct(id string) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("id = ?", id).First(&product).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	product.PrintDetail()
	var input string
	fmt.Println("Tekan (y) untuk memesan produk ini")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consoleReader.ReadString('\n')

	switch input {
	case "y\n":
		CreateOrder(product)
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		// 1 is true, 0 is false
		os.Exit(1)
	default:
		OrderProduct(id)
	}
}

func CreateOrder(product entity.Product) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	var email string
	var adress string
	fmt.Println("Masukan email anda")
	email, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukan alamat anda")
	adress, _ = consoleReader.ReadString('\n')

	order := entity.Order{
		ProductId:   int(product.ID),
		BuyerEmail:  email,
		BuyerAdrees: adress,
		OrderDate:   time.Now(),
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Pesanan berhasil dibuat")

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")
	_, err = fmt.Scanln(&input)

	if err != nil {
		MainMenu()
	}

	switch input {
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		// 1 is true, 0 is false
		os.Exit(1)
	default:
		MainMenu()
	}
}
