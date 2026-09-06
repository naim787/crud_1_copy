package db 

import (
    // "crypto/rand"
// 	"encoding/hex"
// 	"time"
// 	"strconv"
	"fmt"
// 	"bufio"
  )

func ListBuy() {
  ClearScreen()
  //
	var list []ListMoneyBuy

	DB.Find(&list)

  SortDate(list)

	fmt.Println("\n=== DATA PEGGELUARAN DAN PEMASUKAN N ===")

  index := 1
  var totalPrice = 0

  fmt.Println("no  |  id    |  note          date  |   price \n")
  fmt.Println("___________________________________________________ \n")
	for _, list := range list {
		fmt.Printf(
			"%s ~ \033[33m%s\033[0m ~ %s ~ \033[33m%d\033[0m \n",
			list.Id,
			list.Note,
			list.Date,
			list.Price,
		)
    fmt.Println("___________________________________________________ \n")
		totalPrice += list.Price
		// fmt.Println(totalPrice)
		index+=1
	}
	fmt.Println("total semuara harga RP :","\033[33m", totalPrice, "\033[0m")
}