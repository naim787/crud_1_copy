package db

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func UpdateListBuy(scanner *bufio.Scanner) {
	fmt.Print("ID: ")
	scanner.Scan()
	id := scanner.Text()

	var lmb ListMoneyBuy

	result := DB.Where("id = ?", id).First(&lmb)
	if result.Error != nil {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Printf("Nota (%s): ", lmb.Note)
	scanner.Scan()
	note := strings.TrimSpace(scanner.Text())
	if note != "" {
		lmb.Note = note
	}

	fmt.Printf("Date (%s): ", lmb.Date)
	scanner.Scan()
	date := strings.TrimSpace(scanner.Text())
	if date != "" {
		lmb.Date = date
	}

	fmt.Printf("Rp (%d): ", lmb.Price)
	scanner.Scan()
	priceText := strings.TrimSpace(scanner.Text())
	if priceText != "" {
		rp, err := strconv.Atoi(priceText)
		if err != nil {
			fmt.Println("Harga harus berupa angka")
			return
		}
		lmb.Price = rp
	}

	DB.Save(&lmb)

	fmt.Println("Data berhasil diupdate")
}
