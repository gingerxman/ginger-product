package cart

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-product/business/product"
	m_cart "github.com/gingerxman/ginger-product/models/cart"
)

type CartProduct struct {
	eel.EntityBase
	Id int
	UserId int
	PoolProductId int
	ProductSkuName string
	ProductSkuDisplayName string
	Count int
	
	PoolProduct *product.PoolProduct
	
	_isFillValidity bool
	_isValid bool
}

func (this *CartProduct) isValid() bool {
	if !this.PoolProduct.CanPurchase() {
		return false
	}
	
	sku := this.PoolProduct.GetSku(this.ProductSkuName)
	if sku == nil || sku.IsDeleted || !sku.HasStocks() {
		return false
	}
	
	if this.PoolProduct.Product.IsDeleted {
		return false
	}
	
	return true
}

// IsValid 判断购物车商品是否可以购买
func (this *CartProduct) IsValid() bool {
	if !this._isFillValidity {
		this._isValid = this.isValid()
	}
	
	return this._isValid
}

//根据model构建对象
func NewCartProductFromModel(ctx context.Context, model *m_cart.CartItem) *CartProduct {
	instance := new(CartProduct)
	instance.Ctx = ctx
	instance.Model = model
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.PoolProductId = model.PoolProductId
	instance.ProductSkuName = model.ProductSkuName
	instance.ProductSkuDisplayName = model.ProductSkuDisplayName
	instance.Count = model.Count

	return instance
}

func init() {
}
