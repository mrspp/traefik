package main

import (
	"fmt"
	"traefik/service/db"
)

func sort(field string) []db.Item {
	itemArr := db.GetDb()
	temp := itemArr[0]

	switch field {
	case "price":
		for i := 0; i < len(itemArr); i++ {
			for j := i; j < len(itemArr); j++ {
				if itemArr[i].Price > itemArr[j].Price {
					temp = itemArr[i]
					itemArr[i] = itemArr[j]
					itemArr[j] = temp
				}
			}
		}
	case "rating":
		fmt.Println("rating")
	default:
	}
	return itemArr
}
