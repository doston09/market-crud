package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) CategoryCreate(req *models.CategoryCreate) (*models.Category, error) {

	log.Printf("Category create req: %+v\n", req)

	resp, err := c.Strg.Category().Create(req)
	if err != nil {
		log.Printf("error while Category Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) CategoryGetById(req *models.CategoryPrimaryKey) (*models.Category, error) {

	resp, err := c.Strg.Category().GetById(req)
	if err != nil {
		log.Printf("error while get category by id: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) CategoryGetList(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {

	resp, err := c.Strg.Category().GetList(req)
	if err != nil {
		log.Printf("error while get category list: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) CategoryUpdate(req *models.CategoryUpdate) (*models.Category, error) {

	resp, err := c.Strg.Category().Update(req)
	if err != nil {
		log.Printf("error while updating category: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) CategoryDelete(req *models.CategoryPrimaryKey) error {

	err := c.Strg.Category().Delete(req)
	if err != nil {
		log.Printf("error while deleting category: %+v\n", err)
		return err
	}

	return nil
}
