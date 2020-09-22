package cart

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_cart "github.com/gingerxman/ginger-product/models/cart"
)

type CartService struct {
	eel.ServiceBase
}

func NewCartService(ctx context.Context) *CartService {
	service := new(CartService)
	service.Ctx = ctx
	return service
}

func (this *CartService) DeleteShoppingCartItems(ids[] int) {
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Model(&m_cart.CartItem{}).Where("id__in", ids).Delete(&m_cart.CartItem{})
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
}


func init() {
}
