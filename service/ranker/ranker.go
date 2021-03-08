package main

import (
	"fmt"
	"sort"
	"traefik/service/db"
)

func SortInt64(s []int64) []int64 {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

func rankPrice(id int64) (db.Item, int64) {
	itemArr := db.GetDb()
	temp := itemArr[0]
	var IncRank int
	var Item db.Item

	for i := 0; i < len(itemArr); i++ {
		for j := i; j < len(itemArr); j++ {
			if itemArr[i].Price > itemArr[j].Price {
				temp = itemArr[i]
				itemArr[i] = itemArr[j]
				itemArr[j] = temp
			}
		}
	}
	for i := 0; i < len(itemArr); i++ {
		if itemArr[i].ID == id {
			IncRank = i
			Item = itemArr[i]
		}
	}
	return Item, int64(IncRank)
}

func main() {
	// fmt.Println(db.InitConn())
	// itemArr := db.GetDb()

	// fmt.Println(itemArr[1])
	fmt.Println(rankPrice(2675340))
}
