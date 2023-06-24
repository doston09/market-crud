package controller

import (
	"fmt"
	"sort"

	"app/models"
)

// Sort Task 1
func (c *Controller) Sort(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	var resp = &models.ShopCartGetListResponse{}
	var orderDateFilter []*models.ShopCart
	getOrder, err := c.ShopCartGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getOrder.ShopCarts {
		orderDateFilter = append(orderDateFilter, ord)

	}
	sort.Slice(orderDateFilter, func(i, j int) bool {
		return orderDateFilter[i].Time > orderDateFilter[j].Time
	})
	resp.Count = len(orderDateFilter)
	resp.ShopCarts = orderDateFilter
	return resp, nil
}

// Filter Task 2
func (c *Controller) Filter(req *models.ShopCartGetListRequest) ([]*models.ShopCart, error) {
	var orderDateFilter []*models.ShopCart
	getOrder, err := c.ShopCartGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getOrder.ShopCarts {
		if ord.Time >= req.FromTime && ord.Time < req.ToTime {
			orderDateFilter = append(orderDateFilter, ord)
		}
	}
	return orderDateFilter, nil
}

// UserHistory Task 3
func (c *Controller) UserHistory(req *models.UserPrimaryKey) (map[string][]models.History, error) {
	var (
		orders   []models.History
		orderMap = make(map[string][]models.History)
	)
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	getUser, err := c.UserGetById(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	for _, v := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return nil, err
		}

		if v.UserId == req.Id {
			if v.Status == true {
				orders = append(orders, models.History{
					ProductName: getProduct.Name,
					Count:       v.Count,
					Total:       v.Count * getProduct.Price,
					Time:        v.Time,
				})
			}
		}
	}
	orderMap[getUser.Name] = orders
	return orderMap, nil
}

// UserCash Task 4
func (c *Controller) UserCash(req *models.UserPrimaryKey) (map[string]int, error) {
	user := make(map[string]int)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	getUser, err := c.UserGetById(req)

	for _, value := range getOrder.ShopCarts {
		if value.UserId == req.Id {
			if value.Status == true {
				getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
				if err != nil {
					return nil, err
				}
				user[getUser.Name] += value.Count * getProduct.Price
			}
		}
	}
	return user, nil
}

// ProductCountSold Task 4
func (c *Controller) ProductCountSold() (map[string]int, error) {
	product := make(map[string]int)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			product[getProduct.Name] += value.Count
		}

	}
	return product, nil
}

// TopProducts Task 6
func (c *Controller) TopProducts() ([]*models.ProductsHistory, error) {
	var (
		productsMap = make(map[string]int)
		products    []*models.ProductsHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			productsMap[getProduct.Name] += value.Count
		}
	}
	for k, v := range productsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count > products[j].Count
	})

	return products, nil
}

// FailureProducts Task 7
func (c *Controller) FailureProducts() ([]*models.ProductsHistory, error) {
	var (
		productsMap = make(map[string]int)
		products    []*models.ProductsHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			productsMap[getProduct.Name] += value.Count
		}
	}
	for k, v := range productsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count < products[j].Count
	})

	return products, nil
}

// TopTime Task 8
func (c *Controller) TopTime() ([]*models.DateHistory, error) {
	var (
		topTimes = make(map[string]int)
		result   []*models.DateHistory
	)

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.ShopCarts {
		if value.Status == true {
			topTimes[value.Time] += value.Count
		}
	}

	for k, v := range topTimes {
		result = append(result, &models.DateHistory{
			Date:  k,
			Count: v,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result, nil
}

// CategoryHistory Task 9
func (c *Controller) CategoryHistory() ([]*models.CategoryHistory, error) {

	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return nil, err
	}
	for _, value := range getOrder.ShopCarts {
		getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
		if err != nil {
			return nil, err
		}
		fmt.Println(getProduct.CategoryId)
	}
	return nil, nil
}

// ActiveUser Task 10
func (c *Controller) ActiveUser() (string, error) {
	users := make(map[string]int)
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return "", err
	}
	for _, value := range getOrder.ShopCarts {
		if value.Status == true {
			getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
			if err != nil {
				return "", err
			}
			users[value.UserId] += value.Count * getProduct.Price
		}
	}
	user, sum := "", 0
	for key, value := range users {
		if sum < value {
			user = key
			sum = value
		}
	}
	getUser, err := c.UserGetById(&models.UserPrimaryKey{
		Id: user,
	})
	if err != nil {
		return "", err
	}
	return getUser.Name, nil
}

// Bonus Task 11
func (c *Controller) Bonus(req *models.UserPrimaryKey) (int, error) {
	getOrder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{})
	if err != nil {
		return 0, err
	}
	getUser, err := c.UserGetById(&models.UserPrimaryKey{Id: req.Id})
	if err != nil {
		return 0, err
	}
	ProductPrices := []int{}
	sum := 0
	for _, value := range getOrder.ShopCarts {
		if getUser.Id == value.UserId {
			if value.Status == true {
				getProduct, err := c.ProductGetById(&models.ProductPrimaryKey{Id: value.ProductId})
				if err != nil {
					return 0, err
				}
				if value.Count > 9 {
					sum = value.Count * getProduct.Price
					ProductPrices = append(ProductPrices, getProduct.Price)
				}
			}
		}
	}
	sort.Ints(ProductPrices)
	sum -= ProductPrices[0]
	return sum, nil
}
