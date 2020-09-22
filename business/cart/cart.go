package cart

import (
	"errors"
	"fmt"
	"github.com/gingerxman/ginger-product/business/account"
	"github.com/gingerxman/ginger-product/business/product"
	m_cart "github.com/gingerxman/ginger-product/models/cart"
	"time"
	
	
	"github.com/gingerxman/gorm"
	"github.com/gingerxman/eel"
)

type Cart struct {
	eel.EntityBase
	Id int
	UserId int
	CorpId int
	CreatedAt time.Time

	//foreign key
	validProducts []*CartProduct
	invalidProducts []*CartProduct
	isProductsSplitted bool
}

func (this *Cart) fillProducts(cartProducts []*CartProduct) {
	if len(cartProducts) == 0 {
		return
	}
	
	poolProductIds := make([]int, 0)
	for _, cartProduct := range cartProducts {
		poolProductIds = append(poolProductIds, cartProduct.PoolProductId)
	}
	
	//fill product
	corp := account.NewCorpFromOnlyId(this.Ctx, this.CorpId)
	poolProducts := product.GetProductPoolForCorp(this.Ctx, corp).GetPoolProductsByIds(poolProductIds)
	if len(poolProducts) == 0 {
		return
	}
	product.NewFillPoolProductService(this.Ctx).Fill(poolProducts, eel.FillOption{
		"with_sku": true,
	})
	
	//构建<ppid, poolProduct>，因为不同的cart product可能对应相同的pool product，所以通过构建<ppid, poolProduct>完成填充
	ppid2pp := make(map[int]*product.PoolProduct, 0)
	for _, poolProduct := range poolProducts {
		ppid2pp[poolProduct.Id] = poolProduct
	}
	
	//填充CartProduct.Product
	for _, cartProduct := range cartProducts {
		if poolProduct, ok := ppid2pp[cartProduct.PoolProductId]; ok {
			cartProduct.PoolProduct = poolProduct
		}
	}
}

func (this *Cart) splitProducts() error {
	this.isProductsSplitted = true
	o := eel.GetOrmFromContext(this.Ctx)
	
	models := make([]*m_cart.CartItem, 0)
	db := o.Model(&m_cart.CartItem{}).Where(eel.Map{
		"user_id": this.UserId,
		"corp_id": this.CorpId,
	}).Find(&models)
	err := db.Error
	
	if err != nil {
		eel.Logger.Error(err)
		return err
	}
	
	cartProducts := make([]*CartProduct, 0)
	for _, model := range models {
		cartProducts = append(cartProducts, NewCartProductFromModel(this.Ctx, model))
	}
	
	this.fillProducts(cartProducts)
	
	//根据cartProduct.IsValid切分products
	for _, cartProduct := range cartProducts {
		if cartProduct.IsValid() {
			this.validProducts = append(this.validProducts, cartProduct)
		} else {
			this.invalidProducts = append(this.invalidProducts, cartProduct)
		}
	}
	
	return nil
}

// GetValidProducts 获得购物车中有效商品集合
func (this *Cart) GetValidProducts() []*CartProduct {
	if !this.isProductsSplitted {
		err := this.splitProducts()
		if err != nil {
			eel.Logger.Error(err)
		}
	}
	return this.validProducts
}

// GetInvalidProducts 获得购物车中无效商品集合
func (this *Cart) GetInvalidProducts() []*CartProduct {
	if !this.isProductsSplitted {
		err := this.splitProducts()
		if err != nil {
			eel.Logger.Error(err)
		}
	}
	return this.invalidProducts
}

// GetProductCount 获得购物车中商品的数量
func (this *Cart) GetProductCount() int {
	return len(this.GetValidProducts())
}

// AddProduct 向购物车中添加商品
func (this *Cart) AddProduct(poolProduct *product.PoolProduct, skuName string, count int) error {
	isValidProductSku := product.NewProductRepository(this.Ctx).IsValidProductSku(poolProduct.ProductId, skuName)
	if !isValidProductSku {
		return errors.New(fmt.Sprintf("invalid product sku(%d-%s)", poolProduct.ProductId, skuName))
	}
	
	o := eel.GetOrmFromContext(this.Ctx)
	isCartItemExists := o.Model(&m_cart.CartItem{}).Where(eel.Map{
		"pool_product_id": poolProduct.Id,
		"product_sku_name": skuName,
		"user_id": this.UserId,
	}).Exist()
	
	if isCartItemExists {
		//cart item已存在，更新count
		db := o.Model(&m_cart.CartItem{}).Where(eel.Map{
			"user_id": this.UserId,
			"corp_id": this.CorpId,
			"pool_product_id": poolProduct.Id,
			"product_sku_name": skuName,
		}).Update("count", gorm.Expr("count + ?", count))
		
		if db.Error != nil {
			eel.Logger.Error(db.Error)
			return db.Error
		}
	} else {
		//cart item不存在，创建之
		product.NewFillPoolProductService(this.Ctx).FillOne(poolProduct, eel.FillOption{
			"with_sku": true,
		})
		skuDisplayName := poolProduct.GetSku(skuName).GetDisplayName()
		model := m_cart.CartItem{
			UserId: this.UserId,
			CorpId: this.CorpId,
			PoolProductId: poolProduct.Id,
			ProductSkuName: skuName,
			ProductSkuDisplayName: skuDisplayName,
			Count: count,
		}
		
		db := o.Create(&model)
		err := db.Error
		if err != nil {
			eel.Logger.Error(err)
			return err
		}
	}
	
	return nil
}

// DeleteItems 从购物车中删除商品项
func (this *Cart) DeleteItems(ids []int) error {
	if len(ids) == 0 {
		return nil
	}
	
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Where(eel.Map{
		"user_id": this.UserId,
		"corp_id": this.CorpId,
		"id__in": ids,
	}).Delete(&m_cart.CartItem{})
	err := db.Error
	
	if err != nil {
		eel.Logger.Error(err)
		return err
	}
	
	return nil
}

// DeleteItem 从购物车中删除商品项
func (this *Cart) DeleteItem(id int) error {
	return this.DeleteItems([]int{id})
}

func init() {
}
