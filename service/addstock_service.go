// models/stock_transaction.go
type StockTransaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID string    `json:"product_id"`
	Amount    int       `json:"amount"` // จำนวนที่เพิ่ม
	Type      string    `json:"type"`   // เช่น "IN", "OUT", "ADJUST"
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
}

// services/stock_service.go
func AddStock(db *gorm.DB, productID string, qty int, note string) error {
	// ใช้ Transaction เพื่อความปลอดภัยของข้อมูล
	return db.Transaction(func(tx *gorm.DB) error {
		
		// 1. ตรวจสอบว่าสินค้ามีจริงไหม และ Lock Row ไว้ป้องกัน Race Condition
		var product Product
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&product, "id = ?", productID).Error; err != nil {
			return err // Product Not Found
		}

		// 2. อัปเดตจำนวนสต็อกในตารางหลัก
		newStock := product.Stock + qty
		if err := tx.Model(&product).Update("stock", newStock).Error; err != nil {
			return err
		}

		// 3. บันทึกประวัติการเพิ่มสต็อก (History Log)
		log := StockTransaction{
			ProductID: productID,
			Amount:    qty,
			Type:      "IN",
			Remark:    note,
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}

		return nil // สำเร็จ! Transaction จะ Commit อัตโนมัติ
	})
}