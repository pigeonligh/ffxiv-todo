package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/ffxiv-todo/pkg/algorithm"
	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
)

type CalcRequest struct {
	Products    map[string]int `json:"products"`
	Inventories map[string]int `json:"inventories"`
}

func Calculate(c *gin.Context) {
	request := &CalcRequest{}
	if err := c.BindJSON(request); err != nil {
		response(c, failed(err))
		return
	}
	skips := map[string]struct{}{}

	m := cache.GetInstance()
	w := algorithm.NewWorkshop(m)

	for name, amount := range request.Products {
		item := m.Item.Get(name)
		if item == nil {
			skips[name] = struct{}{}
		} else {
			w.AddProduct(item.ID, amount)
		}
	}
	for name, amount := range request.Inventories {
		item := m.Item.Get(name)
		if item == nil {
			skips[name] = struct{}{}
		} else {
			w.AddInventory(item.ID, amount)
		}
	}

	if err := w.Calculate(); err != nil {
		response(c, failed(err))
		return
	}

	report := w.Report()
	response(c, success(gin.H{
		"report": report,
	}))
}
