package cart

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business"
)

type CartRepository struct {
	eel.RepositoryBase
}

func NewCartRepository(ctx context.Context) *CartRepository {
	repository := new(CartRepository)
	repository.Ctx = ctx
	return repository
}

//GetShipInfoInCorp 根据id和user获得ShipInfo对象
func (this *CartRepository) GetCartForUserInCorp(user business.IUser, corp business.ICorp) *Cart {
	cart := Cart{
		UserId: user.GetId(),
		CorpId: corp.GetId(),
	}
	cart.Ctx = this.Ctx
	
	return &cart
}


func init() {
}
