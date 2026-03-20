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
  

# 📑 Stock Transaction

Get Transactions

EndPoint GET /api/transactions

Query Parameters

{

"itemId": "ITEM001",

"type": "ADD"

}

Expected Response

[

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

}

# 🧾 Order

Create Order

EndPoint POST /api/orders

Request

{

"itemId": "ITEM001",

"quantity": 5

}

Expected Response

{

"orderId": "ORDER001",

"orderDate": "2026-03-20",

"status": "CREATED"

}

Update Order Status

EndPoint PUT /api/orders/{orderId}/status

Request

{

"status": "CONFIRMED"

}

Expected Response

{

"message": "Order status updated successfully",

"orderId": "ORDER001",

"status": "CONFIRMED"

}

Get Order

EndPoint GET /api/orders/{orderId}

Expected Response

{

"orderId": "ORDER001",

"orderDate": "2026-03-20",

"status": "CONFIRMED"

}


# 🚚 Delivery Controller

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

# 👨‍💼 Staff Controller

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


# 📝 Design Notes
API นี้ถูกออกแบบตามหลัก RESTful API โดยแยกทรัพยากรออกเป็นหมวดหมู่ เช่น items, inventories, transactions, orders, deliveries และ staff เพื่อให้โครงสร้างของระบบชัดเจนและเข้าใจง่าย

ในแต่ละ endpoint มีการเลือกใช้ HTTP Method ให้เหมาะสมกับการทำงาน ได้แก่

GET สำหรับดึงข้อมูล

POST สำหรับสร้างข้อมูลหรือดำเนินการเกี่ยวกับสต็อก

PUT สำหรับอัปเดตข้อมูลที่มีอยู่แล้ว

การจัดการสต็อก เช่น add, deduct, lock, release และ adjust ถูกแยกเป็น endpoint ตามหน้าที่อย่างชัดเจน เพื่อให้ง่ายต่อการพัฒนาและดูแลระบบในอนาคต

ทุกครั้งที่มีการเปลี่ยนแปลงสต็อก ควรมีการสร้าง StockTransaction เพื่อใช้บันทึกประวัติการเคลื่อนไหวของสินค้า และรองรับการตรวจสอบย้อนหลัง

Order มีความเชื่อมโยงกับการจองสินค้าในคลัง ส่วน Delivery เชื่อมโยงกับกระบวนการจัดส่งสินค้าให้ลูกค้า

ระบบมีการแยกบทบาทของพนักงาน (WarehouseStaff, DeliveryStaff, Admin) เพื่อรองรับการกำหนดสิทธิ์และหน้าที่ของผู้ใช้งานแต่ละประเภท

รูปแบบข้อมูลที่ใช้รับและส่งระหว่าง client และ server คือ JSON เนื่องจากเป็นรูปแบบที่อ่านง่าย ใช้งานแพร่หลาย และเหมาะกับการพัฒนา API

การออกแบบระบบนี้มุ่งเน้นให้มีความชัดเจน รองรับการขยายระบบในอนาคต และสามารถเชื่อมการทำงานระหว่าง Controller, Service, Repository และ Model ได้อย่างเป็นระบบ

# Interface Layor
# 1. Item Service Interface
public interface ItemService {

    Item createItem(String name, String sku, String description);

    void updateItem(String itemId, String name, String sku, String description);

    Item getItemById(String itemId);
}
# 2. Inventory Service Interface
public interface InventoryService {

    void addStock(String itemId, int qty);
    
    void deductStock(String itemId, int qty);
    
    void lockStock(String itemId, int qty, String referenceId);
    
    void releaseStock(String itemId, int qty, String referenceId);
    
    void adjustStock(String itemId, int newQty);
    
    boolean checkLowStock(String itemId);
}
# 3. Stock Transaction Service Interface
import java.util.List;

public interface StockTransactionService {

    void createTransaction(String itemId, String staffId, 
    
                           StockActionType type, int quantity, String referenceId);

    List<StockTransaction> getTransactions(String itemId, StockActionType type);
}
# 4. Order Service Interface
public interface OrderService {
    Order createOrder(String itemId, int quantity);
    void updateOrderStatus(String orderId, OrderStatus status);
    Order getOrderById(String orderId);
}
# 5. Delivery Service Interface
public interface DeliveryService {
    Delivery createDelivery(String orderId, String staffId);
    void updateDeliveryStatus(String deliveryId, DeliveryStatus status);
    Delivery getDeliveryById(String deliveryId);
}
# 6. Staff Service Interface
public interface StaffService {
    boolean login(String staffId, String password);
    Role getRole(String staffId);
}
# 7. Supporting Enums
public enum StockActionType {
    ADD, DEDUCT, ADJUST, LOCK, RELEASE
}
public enum OrderStatus {
    CREATED, CONFIRMED, CANCELLED
}
public enum DeliveryStatus {
    PENDING, SHIPPED, DELIVERED
}
public enum Role {
    WarehouseStaff, DeliveryStaff, Admin
}
