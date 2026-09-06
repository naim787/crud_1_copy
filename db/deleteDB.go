package db

import (
    // "crypto/rand"
// 	"encoding/hex"
// 	"time"
  // "strings"
// 	"strconv"
	"fmt"
	"bufio"
  )

func DeleteListMoney(scanner *bufio.Scanner) {
	fmt.Print("ID ? : ")
	scanner.Scan()

	idStr := scanner.Text()

	DB.Where("id = ?", idStr).Delete(&ListMoneyBuy{})

	fmt.Println("User berhasil dihapus")
	ClearScreen()
}