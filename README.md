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
