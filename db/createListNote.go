package db

import (
  "time"
// 	"strconv"
	"fmt"
	"bufio"
  )

func CeateListNote(scanner *bufio.Scanner) {
  ClearScreen()
	fmt.Print("Text : ")
	scanner.Scan()
	note := scanner.Text()

  loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		panic(err)
	}

	// Ambil waktu sekarang sesuai WITA
	n := time.Now().In(loc)
	listN := ListNote{
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
	}

	DB.Create(&listN)

	fmt.Println("DATA BERHASIL DI TAMBAHKAN....")
	ClearScreen()
}
