package main

import "fmt"

// Item ... 商品の品目と値段
type Item struct {
	Category string
	Price    int
}

func main() {
	item := inputItem()

	fmt.Println("============")
	fmt.Printf("%sに%d円使いました\n", item.Category, item.Price)
}

// 入力したItemを返す
func inputItem() Item {
	// Item型のitemを定義
	var item Item

	fmt.Print("品目>")
	// 入力した値をitemのCategoryフィールドに入れる
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	// 入力した値をitemのPriceフィールドに入れる
	fmt.Scan(&item.Price)

	return item
}
