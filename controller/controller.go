package controller

import (
	"app/config"
	"app/storage"
)

type Controller struct {
	Cfg  *config.Config   `json:"cfg,omitempty" :"cfg"`
	Strg storage.StorageI `json:"strg,omitempty" :"strg"`
}

func NewController(cfg *config.Config, storage storage.StorageI) *Controller {
	return &Controller{
		Cfg:  cfg,
		Strg: storage,
	}
}
