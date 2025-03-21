package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,numeric"`
}

type Receipt struct {
	ID           string `json:"id,omitempty"`
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []Item `json:"items" validate:"required,dive"`
	Total        string `json:"total" validate:"required,numeric"`
	Points       int    `json:"points,omitempty"`
}

type ReceiptQueryParams struct {
	PointsGreater  int    `form:"points_greater_than" validate:"omitempty,numeric,gte=0"`
	PurchasedAfter string `form:"purchased_after" validate:"omitempty,datetime=2006-01-02 15:04"`
}

type ReceiptQueryReturn struct {
	ReceiptsReturned int       `json:"resultsReturned"`
	Receipts         []Receipt `json:"receipts"`
}

var validate *validator.Validate

var receipts []Receipt

func main() {
	validate = validator.New()

	router := gin.Default()
	router.GET("/receipts/:id/points", getReceiptPoints)
	router.GET("/receipts/query", queryReceipt)
	router.POST("/receipts/process", postReceipt)

	fmt.Println("Server listening on port 8080")
	router.Run("localhost:8080")
}

// Finds the receipt whose ID matches the ID given
func getReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	for _, receipt := range receipts {
		if receipt.ID == id {
			c.JSON(http.StatusAccepted, gin.H{"points": receipt.Points})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "No receipt found for that ID."})
}

// Takes a given JSON receipt,
func postReceipt(c *gin.Context) {
	var newReceipt Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid."})
		return
	}

	if err := validate.Struct(newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid."})
		return
	}

	// Generate new ID and attach it to the receipt
	id := uuid.New()
	newReceipt.ID = id.String()

	// Calculate receipt points
	newReceipt.Points = calculatePoints(newReceipt)

	// Add the receipt to the list
	receipts = append(receipts, newReceipt)

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Get receipts bases on specified criteria
func queryReceipt(c *gin.Context) {
	var params ReceiptQueryParams

	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "query format invalid"})
		return
	}

	if err := validate.Struct(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resultQuery := searchReceipts(params.PointsGreater, params.PurchasedAfter)
	message := []ReceiptQueryReturn{
		{
			ReceiptsReturned: len(resultQuery),
			Receipts:         resultQuery,
		},
	}

	c.JSON(http.StatusAccepted, gin.H{"result": message})

}

func searchReceipts(pointsGreater int, purchased string) []Receipt {
	var resultQuery []Receipt

	purchasedAfter, _ := time.Parse("2006-01-02 15:04", purchased)

	for _, receipt := range receipts {
		receiptPurchaseDateTime, _ := time.Parse("2006-01-02 15:04", (receipt.PurchaseDate + " " + receipt.PurchaseTime))

		if receipt.Points >= pointsGreater && receiptPurchaseDateTime.After(purchasedAfter) {
			resultQuery = append(resultQuery, receipt)
		}
	}

	return resultQuery
}

func calculatePoints(r Receipt) int {
	var points int = 0

	// Add the total rounded up to the points
	total, _ := strconv.ParseFloat(r.Total, 64)
	points += int(math.Ceil(total))

	//5 points for every two items on the receipt.
	pairs := math.Floor(float64(len(r.Items) / 2))
	points += (int(pairs) * 5)

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")

	purchaseTime, _ := time.Parse("15:04", r.PurchaseTime)
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	return points
}
