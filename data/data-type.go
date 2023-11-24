package data

type Brand struct {
	BrandId        int    `json:"brand_id"`
	BrandName      string `json:"brand_name"`
	BrandStatusId  int    `json:"brand_status_id"`
	BrandCreatedAt string `json:"brand_created_at"`
}

type Category struct {
	CategoryId        int    `json:"category_id"`
	CategoryParentId  int    `json:"category_parent_id"`
	CategorySequence  int    `json:"category_sequence"`
	CategoryName      string `json:"category_name"`
	CategoryStatusId  int    `json:"category_status_id"`
	CategoryCreatedAt string `json:"category_created_at"`
}
