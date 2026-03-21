package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// โครงสร้างสำหรับรับข้อมูล JSON
type AddStockRequest struct {
	ProductID   string `json:"product_id" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,gt=0"`
	Remark      string `json:"remark"`
	RecordedBy  string `json:"recorded_by"`
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/stocks", func(c *gin.Context) {
		var req AddStockRequest

		// 1. Bind JSON เข้ากับ Struct และ Validate เบื้องต้น
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง: " + err.Error()})
			return
		}

		// 2. Logic การเพิ่ม Stock (เช่น ติดต่อ Database)
		TODO: db.Model(&Product{}).Where("id = ?", req.ProductID).Update("stock", ...)

		// 3. ส่ง Response กลับ
		c.JSON(http.StatusOK, gin.H{
			"message": "เพิ่มสต็อกสำเร็จ",
			"data":    req,
		})
	})

	r.Run(":8080")
}