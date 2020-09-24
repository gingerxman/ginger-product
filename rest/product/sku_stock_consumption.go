package product

import (
	"github.com/gingerxman/eel"
	b_product "github.com/gingerxman/ginger-product/business/product"
)

type SkuStockConsumption struct {
	eel.RestResource
}

func (this *SkuStockConsumption) Resource() string {
	return "product.sku_stock_consumption"
}

func (this *SkuStockConsumption) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"sku_id:int",
			"count:int",
		},
		"DELETE": []string{
			"sku_id:int",
			"count:int",
		},
	}
}

func (this *SkuStockConsumption) Put(ctx *eel.Context) {
	req := ctx.Request
	skuId, _ := req.GetInt("sku_id")
	count, _ := req.GetInt("count")
	

	bCtx := ctx.GetBusinessContext()
	err := b_product.NewUpdateSkuStockService(bCtx).Use(skuId, count)
	if err != nil {
		ctx.Response.Error("sku_stock_consumption:use_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}

func (this *SkuStockConsumption) Delete(ctx *eel.Context) {
	req := ctx.Request
	skuId, _ := req.GetInt("sku_id")
	count, _ := req.GetInt("count")
	
	
	bCtx := ctx.GetBusinessContext()
	err := b_product.NewUpdateSkuStockService(bCtx).Add(skuId, count)
	if err != nil {
		ctx.Response.Error("sku_stock_consumption:add_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}
