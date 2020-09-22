package cart

type RCartProduct struct {
	Id int `json:"id"`
	CartItemId int `json:"cart_item_id"`
	Name string `json:"name"`
	SkuName string `json:"sku_name"`
	SkuDisplayName string `json:"sku_display_name"`
	Thumbnail string `json:"thumbnail"`
	Price int `json:"price"`
	Stocks int `json:"stocks"`
	PurchaseCount int `json:"purchase_count"`
}

type RCartProductGroup struct {
	Products []*RCartProduct `json:"products"`
}

type RCart struct {
	ProductGroups []*RCartProductGroup `json:"product_groups"`
	InvalidProducts []*RCartProduct `json:"invalid_products"`
}


func init() {
}
