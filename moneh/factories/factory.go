package factories

import (
	"fmt"
	"log"
	"moneh/factories/seeders"
	"strconv"
	"strings"
)

func Factory() {
	var purpose string
	var factMod int16
	var total int
	var isPrint bool
	msgFailed := "Failed to generate"

	log.Println("\nWelcome to Moneh Data Factories\n ")

	fmt.Print("Factories Name : ")
	fmt.Scanln(&purpose)

	trimmedPurpose := strings.TrimSpace(purpose)
	if len(trimmedPurpose) > 0 {
		fmt.Print("Factories Module : \n[1] Dictionary\n[2] Tag\n[3] Wishlist\nChoose a module to auto generate dummy : ")
		fmt.Scanln(&factMod)

		fmt.Print("How many data (Max : 100) : ")
		fmt.Scanln(&total)

		var toogle string
		fmt.Print("\nShow result in command [Y/n] : ")
		fmt.Scanln(&toogle)

		if toogle == "Y" || toogle == "n" {
			if toogle == "Y" {
				isPrint = true
			} else if toogle == "n" {
				isPrint = false
			}

			switch factMod {
			case 1:
				seeders.SeedDictionaries(total, isPrint)
			case 2:
				seeders.SeedTags(total, isPrint)
				// ID is empty after run
			case 3:
				seeders.SeedWishlists(total, isPrint)
				// ID is empty after run
			default:
				log.Println("\n" + msgFailed + " : Invalid module")
			}

			if total > 0 && total <= 100 {
				log.Println("\nSuccess run : ", purpose)
				num := strconv.Itoa(total)
				log.Println("With " + num + " data created")
			} else {
				if total <= 0 {
					log.Println("\n" + msgFailed + " : Total is invalid")
				} else {
					log.Println("\n" + msgFailed + " : Data too many")
				}
			}
		} else {
			log.Println("\n" + msgFailed + " : Command not valid")
		}
	} else {
		log.Println("\n" + msgFailed + " : Factories name cant be empty")
	}
}
