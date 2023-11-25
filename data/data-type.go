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

/* Brand insert
{
  "brand_id": 5,
  "brand_name": "New Balance",
  "brand_status_id": 1,
  "brand_created_at": "2023-11-25T09:30:00Z"
}
*/

/* Category insert
{
  "category_id": 1,
  "category_parent_id": 0,
  "category_sequence": "A",
  "category_name": "Running Shoes",
  "category_status_id": 1,
  "category_created_at": "2023-11-25T12:30:00Z"
}
*/

/* supplier insert
{
  "supplier_id": 5,
  "supplier_name": "InnoTech Innovations",
  "supplier_email": "inno@example.com",
  "supplier_phone": "+1 (555) 678-9012",
  "supplier_status_id": 1,
  "supplier_is_verified_supplier": true,
  "supplier_created_at": "2023-11-25T19:00:00Z"
}
*/

/*
{
    "stock_id": 1,
    "product_id": 123,
    "stock_quantity": 100,
    "updated_at": "2023-01-01T12:00:00Z"
}
*/

/* Product insert
{
  "product_id": 1,
  "product_name": "Running Shoes - Model A",
  "product_description": "High-performance running shoes for all terrains.",
  "product_specifications": "Size: 9, Color: Blue",
  "product_brand_id": 1,
  "product_category_id": 1,
  "product_supplier_id": 1,
  "product_unit_price": 89.99,
  "product_discount_price": 79.99,
  "product_tags": "running, sports, shoes",
  "product_status_id": 1
}
*/
