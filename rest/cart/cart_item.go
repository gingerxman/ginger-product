package cart

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/account"
	b_cart "github.com/gingerxman/ginger-product/business/cart"
	"github.com/gingerxman/ginger-product/business/product"
)

type CartItem struct {
	eel.RestResource
}

func (this *CartItem) Resource() string {
	return "cart.cart_item"
}

func (this *CartItem) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"pool_product_id:int", "sku_name", "count:int"},
		"DELETE": []string{"id:int"},
	}
}

func (this *CartItem) Put(ctx *eel.Context) {
	req := ctx.Request
	poolProductId, _ := req.GetInt("pool_product_id")
	skuName := req.GetString("sku_name")
	count, _ := req.GetInt("count")

	bCtx := ctx.GetBusinessContext()
	user := account.GetUserFromContext(bCtx)
	corp := account.GetCorpFromContext(bCtx)
	productPool := product.GetProductPoolForCorp(bCtx, corp)
	poolProduct := productPool.GetPoolProduct(poolProductId)
	
	shoppingCart := b_cart.NewCartRepository(bCtx).GetCartForUserInCorp(user, corp)
	err := shoppingCart.AddProduct(poolProduct, skuName, count)
	
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("cart_item:create_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{
			"count": shoppingCart.GetProductCount(),
		})
	}
}

func (this *CartItem) Delete(ctx *eel.Context) {
	req := ctx.Request
	
	bCtx := ctx.GetBusinessContext()
	user := account.GetUserFromContext(bCtx)
	corp := account.GetCorpFromContext(bCtx)

	id, _ := req.GetInt("id")
	shoppingCart := b_cart.NewCartRepository(bCtx).GetCartForUserInCorp(user, corp)
	err := shoppingCart.DeleteItem(id)
	
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("cart_item:delete_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}

