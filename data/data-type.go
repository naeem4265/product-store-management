package data

type Brand struct {
	BrandId        int    `json:"brand_id" bson:"brand_id"`
	BrandName      string `json:"brand_name" bson:"brand_name"`
	BrandStatusId  int    `json:"brand_status_id" bson:"brand_status_id"`
	BrandCreatedAt string `json:"brand_created_at" bson:"brand_created_at"`
}

type Category struct {
	CategoryId        int    `json:"category_id" bson:"category_id"`
	CategoryParentId  int    `json:"category_parent_id" bson:"category_parent_id"`
	CategorySequence  string `json:"category_sequence" bson:"category_sequence"`
	CategoryName      string `json:"category_name" bson:"category_name"`
	CategoryStatusId  int    `json:"category_status_id" bson:"category_status_id"`
	CategoryCreatedAt string `json:"category_created_at" bson:"category_created_at"`
}

type Product struct {
	ID             int     `json:"product_id" bson:"product_id"`
	Name           string  `json:"product_name" bson:"product_name"`
	Description    string  `json:"product_description" bson:"product_description"`
	Specifications string  `json:"product_specifications" bson:"product_specifications"`
	BrandID        int     `json:"product_brand_id" bson:"product_brand_id"`
	CategoryID     int     `json:"product_category_id" bson:"product_category_id"`
	SupplierID     int     `json:"product_supplier_id" bson:"product_supplier_id"`
	UnitPrice      float64 `json:"product_unit_price" bson:"product_unit_price"`
	DiscountPrice  float64 `json:"product_discount_price" bson:"product_discount_price"`
	Tags           string  `json:"product_tags" bson:"product_tags"`
	StatusID       int     `json:"product_status_id" bson:"product_status_id"`
}
