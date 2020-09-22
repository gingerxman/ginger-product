package cart

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/account"
	b_cart "github.com/gingerxman/ginger-product/business/cart"
)

type Cart struct {
	eel.RestResource
}

func (this *Cart) Resource() string {
	return "cart.cart"
}

func (this *Cart) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *Cart) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	user := account.GetUserFromContext(bCtx)
	corp := account.GetCorpFromContext(bCtx)
	if corp == nil {
		ctx.Response.Error("cart:invalid_corp", fmt.Sprintf("%d", corp.Id))
		return
	}
	
	shoppingCart := b_cart.NewCartRepository(bCtx).GetCartForUserInCorp(user, corp)
	
	if shoppingCart == nil {
		ctx.Response.Error("cart:invalid_shopping_cart", fmt.Sprintf("user(%d), corp(%d)", user.GetId(), corp.GetId()))
	} else {
		encodeService := b_cart.NewEncodeCartService(bCtx)
		respData := encodeService.Encode(shoppingCart)
		
		ctx.Response.JSON(respData)
	}
}
