package product

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/account"
	"github.com/gingerxman/ginger-product/business/product"
)

type CorpProductCount struct {
	eel.RestResource
}

func (this *CorpProductCount) Resource() string {
	return "product.corp_product_count"
}

func (this *CorpProductCount) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"?corp_id:int"},
	}
}

func (this *CorpProductCount) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	corpId, _ := req.GetInt("corp_id", 0)
	var corp *account.Corp
	if corpId != 0{
		corp = account.NewCorpFromOnlyId(bCtx, corpId)
	}else{
		corp = account.GetCorpFromContext(bCtx)
	}

	productPool := product.GetProductPoolForCorp(bCtx, corp)
	onSaleCount := productPool.GetOnSaleCount()
	forSaleCount := productPool.GetForSaleCount()
	lowStockCount := productPool.GetLowStockProductCount()
	
	ctx.Response.JSON(eel.Map{
		"onsale_count": onSaleCount,
		"forsale_count": forSaleCount,
		"low_stock_count": lowStockCount,
	})
}

