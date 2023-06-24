package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
)

func main() {
	cfg := config.Load()
	conn, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, conn)

	// Task 1

	sortedCart, err := con.Sort(&models.ShopCartGetListRequest{})
	for _, val := range sortedCart.ShopCarts {
		fmt.Println(val)
	}

	// Task 2

	shopCart, err := con.Filter(&models.ShopCartGetListRequest{Offset: 0, Limit: 0, FromTime: "2022-09-07 20:16:58", ToTime: "2023-09-07 20:16:58"})
	for _, val := range shopCart {
		fmt.Println(val)
	}

	// Task 3

	userHistory, err := con.UserHistory(&models.UserPrimaryKey{Id: "27457ac2-74dd-4656-b9b0-0d46b1af10dc"})
	if err != nil {
		return
	}
	fmt.Println(userHistory)
	for i, val := range userHistory {
		for key, val := range val {
			fmt.Println(i)
			fmt.Println(key, val)
		}
	}

	// Task 4

	topProduct, err := con.TopProducts()
	for _, val := range topProduct {
		fmt.Println(val)
	}

	// Task 5

	topTime, err := con.TopTime()
	if err != nil {
		return
	}
	for _, val := range topTime {
		fmt.Println(val)
	}

	// Task 9
	history, err := con.CategoryHistory()
	if err != nil {
		return
	}
	for _, val := range history {
		fmt.Println(val)
	}

}
