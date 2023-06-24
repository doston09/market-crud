package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) ProductCreate(req *models.ProductCreate) (*models.Product, error) {

	log.Printf("Create product request: %+v\n", req)

	resp, err := c.Strg.Product().Create(req)
	if err != nil {
		log.Printf("error while creating product: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) ProductGetById(req *models.ProductPrimaryKey) (*models.Product, error) {

	resp, err := c.Strg.Product().GetById(req)
	if err != nil {
		log.Printf("error while get product by id: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductGetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	resp, err := c.Strg.Product().GetList(req)
	if err != nil {
		log.Printf("error while get product list: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductUpdate(req *models.ProductUpdate) (*models.Product, error) {

	resp, err := c.Strg.Product().Update(req)
	if err != nil {
		log.Printf("error while Product Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductDelete(req *models.ProductPrimaryKey) error {

	err := c.Strg.Product().Delete(req)
	if err != nil {
		log.Printf("error while deleting product: %+v\n", err)
		return err
	}

	return nil
}
