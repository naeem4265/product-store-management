# Project Description: E-commerce Management System

The E-commerce Management System is a comprehensive software solution designed to streamline and organize the operations of an online retail business. This system encompasses various modules, each tailored to efficiently handle specific aspects of the e-commerce ecosystem. The core functionalities revolve around managing products, suppliers, brands, categories, and stocks to ensure a seamless and effective online shopping experience.

**Key Features:**

1. **Product Management:**
   - Facilitates the addition, modification, and removal of products in the inventory.
   - Enables detailed product descriptions, specifications, and pricing.
   - Manages product availability and discounts.

2. **Supplier Management:**
   - Tracks information about suppliers, including names, contact details, and verification status.
   - Provides a centralized platform to manage relationships with multiple suppliers.

3. **Brand and Category Management:**
   - Organizes products into brands and categories for easy navigation.
   - Allows for the creation, modification, and deletion of brands and categories.

4. **Stock Management:**
   - Monitors stock levels and updates in real-time.
   - Ensures accurate tracking of inventory to prevent stockouts or overstock situations.
   - Logs updates to stock quantities with timestamps.

5. **User-Friendly Interfaces:**
   - Offers intuitive interfaces for administrators to manage the system efficiently.
   - Provides an easy-to-navigate storefront for customers to browse and purchase products.

6. **Data Integrity and Security:**
   - Ensures data integrity through secure storage and retrieval mechanisms.
   - Implements user authentication and authorization to control access to sensitive information.

7. **Scalability:**
   - Designed to scale with the growth of the business, accommodating an increasing number of products, suppliers, and customers.

8. **Reporting and Analytics:**
   - Generates comprehensive reports on sales, product popularity, and stock status.
   - Provides analytics to assist in decision-making and strategy planning.

**Technologies Used:**
- **Backend:** Go (Golang)
- **Database:** MongoDB
- **APIs:** RESTful APIs for communication between frontend and backend components.

**Purpose:**
The E-commerce Management System is developed to enhance the efficiency, organization, and scalability of e-commerce businesses. By centralizing and automating various processes, the system aims to optimize the management of products, suppliers, and stocks, ultimately providing a superior shopping experience for customers.

**Note:** This project description is a generalized template. Actual details and features may vary based on the specific requirements and scope of your e-commerce management system.

# Steps to run the project

### Step-1: Download and install mongodb. Check and start mongodb
https://www.youtube.com/watch?v=HSIh8UswVVY&t=527s
[mongodb](https://www.mongodb.com/try/download/community)
```
sudo systemctl status mongod
sudo systemctl start mongod
```

### Step-2: Download and Install mongodb Shell and mongodb database tool
[shell](https://www.mongodb.com/try/download/shell)
[tool](https://www.mongodb.com/try/download/database-tools)

### Step-3: Command mongosh to connect database successfully 
    mongosh
### Step-4: Download and install Go and visual studio code
[Go](https://go.dev/dl/) 
[vscode](https://code.visualstudio.com/) 

### Step-5: Download and Install postman for curl request
[postman](https://www.postman.com/downloads/)

### Step-6: Download mongodb database files and restore to the database
[folder](https://github.com/naeem4265/product-store-management/tree/master/database)

    mongorestore --db productStore /path/to/backup_directory/productStore

### Step-7: Go to github directory and mount this your local pc
[Repository](https://github.com/naeem4265/product-store-management)

    git clone git@github.com:naeem4265/product-store-management.git


### Step-8: Run the go files from terminal
```
   go mod tidy && go mod vendor
   go run *.go
```
If the go run command runs successfully, you will see the message "Server started at :8080."


### Step-9: Open postman
#### Products CRUD operation
<li> Make a GET request to the URL 'localhost:8080/products' using Postman. You will receive a list of products.</li><br>

<li> Make a GET request by id to the URL 'localhost:8080/products/id' using Postman. You will receive a product of product id = id.</li>

<br><li>For inserting a product, make a POST request to 'localhost:8080/products' and pass a JSON body with the required information.</li>

```
{
  "product_id": 100,
  "product_name": "Running Shoes - Model A105",
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
```
<br> 

<li> For update a product, make a PUT request to 'localhost:8080/products/id' and pass a JSON body with the required information.</li>


```
{
  "product_id": 100,
  "product_name": "Update name",
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
```

<br> 
<li> For Delete a product, make a DELETE request to 'localhost:8080/products/id'. Product will be deleted. </li>


#### Brands CRUD operation
<li> Make a GET request to the URL 'localhost:8080/brands' using Postman. You will receive a list of brands.</li><br>

<li> Make a GET request by id to the URL 'localhost:8080/brands/id' using Postman. You will receive a brand of brand id = id.</li><br>

<li>For inserting a brand, make a POST request to 'localhost:8080/brands' and pass a JSON body with the required information.</li>

```
{
  "brand_id": 500,
  "brand_name": "New Brand",
  "brand_status_id": 1,
  "brand_created_at": "2023-11-25T09:30:00Z"
}
```
<br> 
<li> For update a brand, make a PUT request to 'localhost:8080/brands/id' and pass a JSON body with the required information.</li>

```
{
  "brand_id": 500,
  "brand_name": "Update Brand",
  "brand_status_id": 1,
  "brand_created_at": "2023-11-25T09:30:00Z"
}
```

<br> <li> For Delete a brand, make a DELETE request to 'localhost:8080/brands/id'. brand will be deleted. </li>


#### Categories CRUD operation
<li> Make a GET request to the URL 'localhost:8080/categories' using Postman. You will receive a list of categories.</li><br>

<li> Make a GET request by id to the URL 'localhost:8080/categories/id' using Postman. You will receive a category of category id = id.</li><br>

<li>For inserting a category, make a POST request to 'localhost:8080/categories' and pass a JSON body with the required information.</li>

```
{
  "category_id": 100,
  "category_parent_id": 0,
  "category_sequence": "A",
  "category_name": "Running Shoes",
  "category_status_id": 1,
  "category_created_at": "2023-11-25T12:30:00Z"
}
```
<br> 
<li> For update a category, make a PUT request to 'localhost:8080/categories/id' and pass a JSON body with the required information.</li>

```
{
  "category_id": 100,
  "category_parent_id": 0,
  "category_sequence": "A",
  "category_name": "Updated name",
  "category_status_id": 1,
  "category_created_at": "2023-11-25T12:30:00Z"
}
```

<br> <li> For Delete a category, make a DELETE request to 'localhost:8080/categories/id'. category will be deleted. </li>


#### Suppliers CRUD operation
<li> Make a GET request to the URL 'localhost:8080/suppliers' using Postman. You will receive a list of suppliers.</li><br>

<li> Make a GET request by id to the URL 'localhost:8080/suppliers/id' using Postman. You will receive a supplier of supplier id = id.</li><br>

<li>For inserting a supplier, make a POST request to 'localhost:8080/suppliers' and pass a JSON body with the required information.</li>

```
{
  "supplier_id": 500,
  "supplier_name": "InnoTech Innovations500",
  "supplier_email": "inno@example.com",
  "supplier_phone": "+1 (555) 678-9012",
  "supplier_status_id": 1,
  "supplier_is_verified_supplier": true,
  "supplier_created_at": "2023-11-25T19:00:00Z"
}
```
<br> 
<li> For update a supplier, make a PUT request to 'localhost:8080/suppliers/id' and pass a JSON body with the required information.</li>

```
{
  "supplier_id": 500,
  "supplier_name": "Update supplier",
  "supplier_email": "inno@example.com",
  "supplier_phone": "+1 (555) 678-9012",
  "supplier_status_id": 1,
  "supplier_is_verified_supplier": true,
  "supplier_created_at": "2023-11-25T19:00:00Z"
}
```

<br> <li> For Delete a supplier, make a DELETE request to 'localhost:8080/suppliers/id'. supplier will be deleted. </li>


#### Stocks CRUD operation
<li> Make a GET request to the URL 'localhost:8080/stocks' using Postman. You will receive a list of stocks.</li><br>

<li> Make a GET request by id to the URL 'localhost:8080/stocks/id' using Postman. You will receive a stock of stock id = id.</li><br>

<li>For inserting a stock, make a POST request to 'localhost:8080/stocks' and pass a JSON body with the required information.</li>

```
{
    "stock_id": 100,
    "product_id": 123,
    "stock_quantity": 100,
    "updated_at": "2023-01-01T12:00:00Z"
}
```
<br> 
<li> For update a stock, make a PUT request to 'localhost:8080/stocks/id' and pass a JSON body with the required information.</li>

```
{
    "stock_id": 100,
    "product_id": 123,
    "stock_quantity": 100,
    "updated_at": "Update time 2023-01-01T12:00:00Z"
}
```

<br> <li> For Delete a stock, make a DELETE request to 'localhost:8080/stocks/id'. stock will be deleted. </li>

</br>

### Note: The project is currently under construction, and there are many ongoing optimization efforts.
