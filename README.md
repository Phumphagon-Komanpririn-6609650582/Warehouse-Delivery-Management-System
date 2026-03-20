# Warehouse-Delivery-Management-System

# 📦Item

**Item**

EndPoint POST /api/items

- Request

  {
  
    "name": "iPhone 15",
  
    "sku": "IP15-001",
    
    "description": "Apple smartphone"
  
  }

- Expected Response

  {
  
    "itemId": "ITEM001",
    
    "name": "iPhone 15",
    
    "sku": "IP15-001",
    
    "description": "Apple smartphone"
  
  }

  **Update Item**

  EndPoint PUT /api/items/{itemId}

- Request
  
  {
    
    "name": "iPhone 15 Pro",
    
    "sku": "IP15-002",
    
    "description": "Updated version"
  
  }

- Expected Response

  {
   
    "message": "Item updated successfully"
  
  }

  **Get Item**

  EndPoint GET /api/items/{itemId}

- Request

  {
  
    itemId = ITEM001

  }

- Expected Response

  {
  
    "itemId": "ITEM001",
  
    "name": "iPhone 15",
  
    "sku": "IP15-001",
  
    "description": "Apple smartphone"
  
  }

  #  🗄️Inventory

  **Add Stock**

  EndPoint POST /api/inventories/{itemId}/add

- Request

  {
  
  "quantity": 100
  
  }

- Expected Response

  {
  
  "availableQuantity": 150,
  
  "reservedQuantity": 5

  }

  **Deduct Stock**

  EndPoint POST /api/inventories/{itemId}/deduct

- Request

  {
  
  "quantity": 10
  
  }

- Expected Response

  {
  
  "availableQuantity": 90

  "reservedQuantity": 5

  }
  
   **Lock Stock (reserve)**

  EndPoint POST /api/inventories/{itemId}/lock

- Request

  {
  
  "quantity": 5,
  
  }

- Expected Response

  {
  
  "availableQuantity": 85,

  "reservedQuantity": 10

  }
  
   **Release Stock**

  EndPoint POST /api/inventories/{itemId}/release

- Request

  {
  
  "quantity": 5,
  
  }

- Expected Response

  {
  
  "availableQuantity": 90,

  "reservedQuantity": 5

  }

  **Adjust Stock**

  EndPoint POST /api/inventories/{itemId}/adjust

- Request

  {
  
  "newQuantity": 200
  
  }

- Expected Response

  {
  
  "availableQuantity": 200,
  
  "reservedQuantity": 0

  }
  
  **Check Low Stock**

  EndPoint GET /api/inventories/{itemId}/low-stock

- Request

  {
  
  itemId = ITEM001
  
  }

- Expected Response

  {
  
  "itemId": "ITEM001",
  
  "lowStock": true,
  
  "availableQuantity": 5,
  
  "threshold": 10
  
  }
  

📑 Stock Transaction


 Get Transactions
**Endpoint:** `GET /api/transactions`

**Query Parameters**
json
{
  "itemId": "ITEM001",
  "type": "ADD"
}
Expected Response

  {
    "transactionId": "TX001",
    "itemId": "ITEM001",
    "type": "ADD",
    "quantity": 50,
    "timestamp": "2026-03-19T10:00:00",
    "referenceId": "ADMIN001"
  },
  {
    "transactionId": "TX002",
    "itemId": "ITEM001",
    "type": "DEDUCT",
    "quantity": 10,
    "timestamp": "2026-03-19T11:00:00",
    "referenceId": "ORDER001"
  }


### 🧾 Order


#### Create Order
**Endpoint:** `POST /api/orders`

**Request**
json
{
  "itemId": "ITEM001",
  "quantity": 5
}

**Expected Response**
{
  "orderId": "ORDER001",
  "orderDate": "2026-03-20",
  "status": "CREATED"
}

**Update Order Status**
Endpoint: PUT /api/orders/{orderId}/status
Request
{
  "status": "CONFIRMED"
}
**Expected Response**
{
  "message": "Order status updated successfully",
  "orderId": "ORDER001",
  "status": "CONFIRMED"
}

Get Order
Endpoint: GET /api/orders/{orderId}
Expected Response
{
  "orderId": "ORDER001",
  "orderDate": "2026-03-20",
  "status": "CONFIRMED"
}


🚚 Delivery Controller

Create Delivery

EndPoint POST /api/deliveries

Request

{

"orderId": "ORDER001",

"staffId": "STAFF002"

}

Expected Response

{

"deliveryId": "DEL001",

"orderId": "ORDER001",

"staffId": "STAFF002",

"status": "PENDING",

"shippedDate": null,

"deliveredDate": null

}

Update Delivery Status

EndPoint PUT /api/deliveries/{deliveryId}/status

Request

{

"status": "SHIPPED"

}

Expected Response

{

"message": "Delivery status updated successfully",

"deliveryId": "DEL001",

"status": "SHIPPED",

"shippedDate": "2026-03-20T10:00:00",

"deliveredDate": null

}

Get Delivery

EndPoint GET /api/deliveries/{deliveryId}

Expected Response

{

"deliveryId": "DEL001",

"orderId": "ORDER001",

"staffId": "STAFF002",

"status": "SHIPPED",

"shippedDate": "2026-03-20T10:00:00",

"deliveredDate": null

}

👨‍💼 Staff Controller

Login

EndPoint POST /api/staff/login

Request

{

"staffId": "STAFF001",

"password": "1234"

}

Expected Response

{

"success": true,

"staffId": "STAFF001",

"role": "Admin"

}

Get Staff Profile

EndPoint GET /api/staff/{staffId}

Expected Response

{

"staffId": "STAFF001",

"name": "John Doe",

"role": "WarehouseStaff"

}

Get Staff Deliveries

EndPoint GET /api/staff/{staffId}/deliveries

Expected Response

[

{

"deliveryId": "DEL001",

"orderId": "ORDER001",

"status": "SHIPPED"

},

{

"deliveryId": "DEL002",

"orderId": "ORDER002",

"status": "DELIVERED"

}

}
