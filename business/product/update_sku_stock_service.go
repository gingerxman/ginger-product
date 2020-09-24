package product

import (
	"context"
	"github.com/gingerxman/gorm"
	"github.com/gingerxman/eel"
	m_product "github.com/gingerxman/ginger-product/models/product"
)

type UpdateSkuStockService struct {
	eel.ServiceBase
}

func NewUpdateSkuStockService(ctx context.Context) *UpdateSkuStockService {
	service := new(UpdateSkuStockService)
	service.Ctx = ctx
	return service
}

func (this *UpdateSkuStockService) Use(skuId int, count int) error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_product.ProductSku{}).Where("id", skuId).Update("stocks", gorm.Expr("stocks - ?", count))
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func (this *UpdateSkuStockService) Add(skuId int, count int) error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_product.ProductSku{}).Where("id", skuId).Update("stocks", gorm.Expr("stocks + ?", count))
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func init() {
}
