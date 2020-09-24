package product

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/product"
)

type Products struct {
	eel.RestResource
}

func (this *Products) Resource() string {
	return "product.products"
}

func (this *Products) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"ids:json-array", "?fill_options:json-array"},
	}
}

func (this *Products) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	ids := req.GetIntArray("ids")
	poolProducts := product.GetGlobalProductPool(bCtx).GetPoolProductsByIds(ids)
	
	fillOptionStrs := req.GetStringArray("fill_options")
	fillOptions := eel.FillOption{}
	for _, optionStr := range fillOptionStrs {
		fillOptions[optionStr] = true
	}
	
	fillService := product.NewFillPoolProductServiceForCorp(bCtx, nil)
	fillService.Fill(poolProducts, fillOptions)

	encodeService := product.NewEncodePoolProductService(bCtx)
	rows := encodeService.EncodeMany(poolProducts)
	
	ctx.Response.JSON(eel.Map{
		"products": rows,
	})
}

