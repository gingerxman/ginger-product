package cart

import (
	"github.com/gingerxman/eel"
)

//CartItem Model
type CartItem struct {
	eel.Model
	UserId int
	CorpId int
	PoolProductId int `gorm:"index"`
	ProductSkuName string `gorm:"size:256"`
	ProductSkuDisplayName string `gorm:"size:256"`
	Count int
}
func (self *CartItem) TableName() string {
	return "cart_item"
}
func (this *CartItem) TableIndex() [][]string {
	return [][]string{
		[]string{"UserId", "CorpId", "PoolProductId"},
	}
}





func init() {
	eel.RegisterModel(new(CartItem))
}
