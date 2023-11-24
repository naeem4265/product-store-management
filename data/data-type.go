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

/*
{
  "product_id": 1,
  "product_name": "Example Product",
  "product_description": "A detailed description of the product.",
  "product_specifications": "Specifications of the product.",
  "product_brand_id": 123,
  "product_category_id": 456,
  "product_supplier_id": 789,
  "product_unit_price": 49.99,
  "product_discount_price": 39.99,
  "product_tags": "electronics, gadgets",
  "product_status_id": 1
}
*/

type Supplier struct {
	ID         int    `json:"supplier_id" bson:"supplier_id"`
	Name       string `json:"supplier_name" bson:"supplier_name"`
	Email      string `json:"supplier_email" bson:"supplier_email"`
	Phone      string `json:"supplier_phone" bson:"supplier_phone"`
	StatusID   int    `json:"supplier_status_id" bson:"supplier_status_id"`
	IsVerified bool   `json:"supplier_is_verified_supplier" bson:"supplier_is_verified_supplier"`
	CreatedAt  string `json:"supplier_created_at" bson:"supplier_created_at"`
}

type Stock struct {
	ID            int    `json:"stock_id" bson:"stock_id"`
	ProductID     int    `json:"product_id" bson:"product_id"`
	StockQuantity int    `json:"stock_quantity" bson:"stock_quantity"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
}
