package db 

import (
  // "crypto/rand"
// 	"encoding/hex"
	"time"
	"strconv"
	"fmt"
	"bufio"
  )

func CreateListBuy(scanner *bufio.Scanner) {
  ClearScreen()
	fmt.Print("NOTE: ")
	scanner.Scan()
	note := scanner.Text()

	fmt.Print("RP: ")
	scanner.Scan()
	price := scanner.Text()

	pr, _ := strconv.Atoi(price)

  loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		panic(err)
	}

	// Ambil waktu sekarang sesuai WITA
	n := time.Now().In(loc)
	listMONEY := ListMoneyBuy{
	  Id: RandomString(),
		Note: note,
		Date: fmt.Sprintf(
    	"%02d-%02d-%d | %02d:%02d:%02d",
    	n.Day(),
    	n.Month(),
    	n.Year(),
    	n.Hour(),
    	n.Minute(),
    	n.Second(),
   ),
		Price:  pr,
	}

	DB.Create(&listMONEY)

	fmt.Println("DATA BERHASIL DI TAMBAHKAN....")
	ClearScreen()
}