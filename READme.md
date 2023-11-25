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
[files](https://github.com/naeem4265/product-store-management)

   mongorestore --db productStore /path/to/backup_directory/your_database_name

### Step-7: Go to github directory and mount this your local pc
[Repository](https://github.com/naeem4265/product-store-management)

    git clone git@github.com:naeem4265/product-store-management.git


### Step-8: Run the go files from terminal
    go run *.go
If the go run command runs successfully, you will see the message "Server started at :8080."


### Step-9: Open postman 
<li> Make a GET request to the URL localhost:8080/products using Postman. You will receive a list of products.</li>

<br><li>For inserting a product, make a POST request to localhost:8080/products and pass a JSON body with the required information.</li>

```
{
  "product_id": 105,
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
<br> <li> For update a product, make a PUT request to localhost:8080/products/id and pass a JSON body with the required information.</li>
```
{
  "product_id": id,
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

<br> <li> For Delete a product, make a DELETE request to localhost:8080/products/id. Product will be deleted. </li>


### Using a similar approach, you can perform CRUD operations on the "brands," "suppliers," "categories," and "stocks" databases.

</br>

### Note: The project is currently under construction, and there are many ongoing optimization efforts.