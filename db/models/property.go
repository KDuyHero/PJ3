package gormModel

type Property struct {
	Id        int64  `gorm:"primary_key;column:id"`
	ProductId int64  `gorm:"column:product_id"`
	Name      string `gorm:"column:name"`
	Value     string `gorm:"column:value"`
}

func (Property) TableName() string { return "properties" }
