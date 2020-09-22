package cart

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/account"
	b_cart "github.com/gingerxman/ginger-product/business/cart"
)

type ProductCount struct {
	eel.RestResource
}

func (this *ProductCount) Resource() string {
	return "cart.product_count"
}

func (this *ProductCount) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *ProductCount) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	corp := account.GetCorpFromContext(bCtx)
	if corp == nil {
		ctx.Response.Error("cart.product_count:invalid_corp", "")
		return
	}

	user := account.GetUserFromContext(bCtx)
	shoppingCart := b_cart.NewCartRepository(bCtx).GetCartForUserInCorp(user, corp)
	count := shoppingCart.GetProductCount()
	
	ctx.Response.JSON(eel.Map{
		"count": count,
	})
}
