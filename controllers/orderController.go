package controllers

import (
	"assignment1/database"
	"assignment1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()

	var orders []models.Order

	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var Order models.Order

	err := ctx.ShouldBindJSON(&Order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Create(&Order).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Order,
	})
}

func GetOrderByID(ctx *gin.Context) {
	db := database.GetDB()
	orderID := ctx.Param("id")

	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Invalid order ID",
		})
		return
	}

	var order models.Order

	err = db.Preload("Items").First(&order, orderIDInt).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Order Not Found",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrderByID(ctx *gin.Context) {
	db := database.GetDB()
	orderID := ctx.Param("id")

	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Invalid order ID",
		})
		return
	}

	result := db.Delete(&models.Order{}, orderIDInt)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Order Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func UpdateOrderByID(ctx *gin.Context) {
	db := database.GetDB()
	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Invalid order ID",
		})
		return
	}

	var ExistingOrder models.Order
	err = db.First(&ExistingOrder, orderIDInt).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Order Not Found",
		})
		return
	}

	var Order models.Order
	err = ctx.ShouldBindJSON(&Order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// clear old items
		err = tx.Where("order_id = ?", orderIDInt).Delete(&models.Item{}).Error
		if err != nil {
			return err
		}

		ExistingOrder.CustomerName = Order.CustomerName
		ExistingOrder.OrderedAt = Order.OrderedAt
		ExistingOrder.Items = Order.Items

		err = tx.Save(&ExistingOrder).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed updating the order",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": ExistingOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var Order models.Order
	var ExistingOrder models.Order
	var NewItems []models.Item

	err := ctx.ShouldBindJSON(&Order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Preload("Items").First(&ExistingOrder, Order.ID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// Update Order karena ada order.id
		// Upsert item sesuai yang diinput user.
		// Jika ada item.id -> update
		// tanpa item.id -> akan create
		err = tx.Save(&Order).Error
		if err != nil {
			return err
		}

		NewItems = Order.Items
		ItemsID := make(map[uint]bool)
		for _, item := range NewItems {
			ItemsID[item.ID] = true
		}

		// Hapus item lama yang tidak digunakan lagi
		for _, item := range ExistingOrder.Items {
			if ItemsID[item.ID] {
				continue
			}
			// Karena tidak ada di ItemsID maka bisa dihapus
			tx.Delete(&item)
		}
		return nil
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Preload("Items").First(&ExistingOrder, Order.ID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": ExistingOrder,
	})
}
