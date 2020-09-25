package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-product/rest/area"
	"github.com/gingerxman/ginger-product/rest/cart"
	"github.com/gingerxman/ginger-product/rest/dev"
	"github.com/gingerxman/ginger-product/rest/product"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	/*
	 product
	 */
	//category
	eel.RegisterResource(&product.Category{})
	eel.RegisterResource(&product.CategoryDisplayIndex{})
	eel.RegisterResource(&product.DisabledCategory{})
	eel.RegisterResource(&product.SubCategories{})
	//label
	eel.RegisterResource(&product.ProductLabel{})
	eel.RegisterResource(&product.ProductLabels{})
	eel.RegisterResource(&product.CorpProductLabels{})
	eel.RegisterResource(&product.DisabledCategory{})
	//property
	eel.RegisterResource(&product.ProductProperty{})
	eel.RegisterResource(&product.ProductPropertyValue{})
	eel.RegisterResource(&product.ProductProperties{})
	eel.RegisterResource(&product.CorpProductProperties{})
	//product
	eel.RegisterResource(&product.Product{})
	eel.RegisterResource(&product.Products{})
	eel.RegisterResource(&product.OffshelfProducts{})
	eel.RegisterResource(&product.OnshelfProducts{})
	eel.RegisterResource(&product.CorpProducts{})
	eel.RegisterResource(&product.CreateOptions{})
	eel.RegisterResource(&product.CorpProductCount{})
	eel.RegisterResource(&product.SkuStockConsumption{})
	
	/*
	 cart
	 */
	eel.RegisterResource(&cart.CartItem{})
	eel.RegisterResource(&cart.Cart{})
	eel.RegisterResource(&cart.ProductCount{})
	
	/*
	 material
	*/
	//eel.RegisterResource(&material.Image{})
	
	/*
	 area
	 */
	eel.RegisterResource(&area.Area{})
	eel.RegisterResource(&area.AreaCode{})
	eel.RegisterResource(&area.YouzanAreaList{})

	/*
	 dev
	 */
	eel.RegisterResource(&dev.BDDReset{})
}