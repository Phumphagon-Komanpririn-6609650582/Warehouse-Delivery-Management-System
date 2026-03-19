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

  #  🗄️Inventory

  **Add Stock**

  EndPoint POST /api/inventories/{itemId}/add

- Request

  {
  
  "quantity": 100
  
  }

- Expected Response

  {
  
  "availableQuantity": 150

  }

  **Deduct Stock**

  EndPoint POST /api/inventories/{itemId}/deduct

  - Request

  {
  
  "quantity": 10
  
  }

   **Lock Stock (reserve)**

  EndPoint POST /api/inventories/{itemId}/lock

  - Request

  {
  
  "quantity": 5,
  
  "referenceId": "ORDER001"
  
  }

   **Release Stock**

  EndPoint POST /api/inventories/{itemId}/release

  **Adjust Stock**

  EndPoint POST /api/inventories/{itemId}/adjust

  - Request

  {
  
  "newQuantity": 200
  
  }

  **Check Low Stock**

  EndPoint GET /api/inventories/{itemId}/low-stock

  - Request

  {
  
  "lowStock": true
  
  }
