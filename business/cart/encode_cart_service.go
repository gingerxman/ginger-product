package cart

import (
	"context"
	
	"github.com/gingerxman/eel"
)

type EncodeCartService struct {
	eel.ServiceBase
}

func NewEncodeCartService(ctx context.Context) *EncodeCartService {
	service := new(EncodeCartService)
	service.Ctx = ctx
	return service
}

func (this *EncodeCartService) encodeProducts(cartProducts []*CartProduct) []*RCartProduct {
	rValidProducts := make([]*RCartProduct, 0)
	for _, cartProduct := range cartProducts {
		price := 0
		stocks := 0
		if cartProduct.IsValid() {
			productSku := cartProduct.PoolProduct.GetSku(cartProduct.ProductSkuName)
			price = productSku.Price
			stocks = productSku.Stocks
		}
		
		rValidProduct := RCartProduct{
			Id: cartProduct.PoolProduct.Id,
			CartItemId: cartProduct.Id,
			Name: cartProduct.PoolProduct.Product.Name,
			SkuName: cartProduct.ProductSkuName,
			SkuDisplayName: cartProduct.ProductSkuDisplayName,
			Thumbnail: cartProduct.PoolProduct.Product.Thumbnail,
			Price: price,
			Stocks: stocks,
			PurchaseCount: cartProduct.Count,
		}
		rValidProducts = append(rValidProducts, &rValidProduct)
	}
	
	return rValidProducts
}

//Encode 对单个实体对象进行编码
func (this *EncodeCartService) Encode(cart *Cart) *RCart {
	if cart == nil {
		return nil
	}

	//编码product_groups
	rProductGroups := make([]*RCartProductGroup, 0)
	validProducts := cart.GetValidProducts()
	if len(validProducts) > 0 {
		rProductGroups = append(rProductGroups, &RCartProductGroup{
			Products: this.encodeProducts(validProducts),
		})
	}
	
	//编码invalid products
	return &RCart{
		ProductGroups: rProductGroups,
		InvalidProducts: this.encodeProducts(cart.GetInvalidProducts()),
	}
}

func init() {
}
