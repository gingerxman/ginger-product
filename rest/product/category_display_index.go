package product

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/product"
)

type CategoryDisplayIndex struct {
	eel.RestResource
}

func (this *CategoryDisplayIndex) Resource() string {
	return "product.category_display_index"
}

func (this *CategoryDisplayIndex) GetParameters() map[string][]string {
	return map[string][]string{
		"POST": []string{
			"id:int",
			"action:string",
		},
	}
}

func (this *CategoryDisplayIndex) Post(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")
	action := req.GetString("action")

	bCtx := ctx.GetBusinessContext()
	repository := product.NewProductCategoryRepository(bCtx)
	productCategory := repository.GetProductCategory(id)
	if productCategory == nil {
		ctx.Response.Error("category_display_index:invalid_category", fmt.Sprintf("id=%d", id))
		return
	}
	
	err := productCategory.UpdateDisplayIndex(action)
	if err != nil {
		eel.Logger.Error(err)
	}
	
	ctx.Response.JSON(eel.Map{})
}
