package main

import (
	"bufio"
	"fmt"
	"os" 
  // "path/filepath"
)
import 	"crud_1_copy/db"



func main() {
	scanner := bufio.NewScanner(os.Stdin)
  
	db.ClearScreen();
  db.ReadFileNamedb();

  fmt.Print("PILIH  \033[33m DB :\033[0m ")
  scanner.Scan()
  
	nmDB := scanner.Text()
  
  db.OPENDB(nmDB)

  
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Create User")
		fmt.Println("2. List User")
		fmt.Println("3. Update User")
		fmt.Println("4. Delete User")
		fmt.Println("5. Exit")
		fmt.Print("Pilih: ")

		scanner.Scan()
		menu := scanner.Text()

		switch menu {
		case "1":
			db.CreateListBuy(scanner)
		case "2":
			db.ListBuy()
		case "3":
			db.UpdateListBuy(scanner)
		case "4":
			db.DeleteListMoney(scanner)
		case "n1":
			db.CeateListNote(scanner)
		case "5":
			fmt.Println("Keluar...")
			db.ClearScreen()
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}