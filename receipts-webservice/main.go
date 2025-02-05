package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	ID           string `json:"id,omitempty"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
	Points       int    `json:"points,omitempty"`
}

var receipts []Receipt

func main() {
	router := gin.Default()
	router.GET("/receipts/:id/points", getReceiptPoints)
	router.POST("/receipts/process", postReceipt)

	fmt.Println("Server listening on port 8080")
	router.Run("localhost:8080")
}

// Finds the receipt whose ID matches the ID given
func getReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	for _, receipt := range receipts {
		if receipt.ID == id {
			c.IndentedJSON(http.StatusAccepted, gin.H{"points": receipt.Points})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that ID."})
}

// Takes a given JSON receipt,
func postReceipt(c *gin.Context) {
	var newReceipt Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid."})
		return
	}

	// Generate new ID and attach it to the receipt
	id := uuid.New()
	newReceipt.ID = id.String()

	// Calculate receipt points
	newReceipt.Points = calculatePoints(newReceipt)

	// Add the receipt to the list
	receipts = append(receipts, newReceipt)

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func calculatePoints(r Receipt) int {
	var points int = 0

	//One point for every alphanumeric character in the retailer name.
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	alphanumeric := re.FindAllString(r.Retailer, -1)
	joined := strings.Join(alphanumeric, "")
	points += len(joined)

	//50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(r.Total, 64)
	if total == math.Ceil(total) {
		points += 50
	}

	//25 points if the total is a multiple of 0.25.
	modTotal := math.Mod(total, 0.25)
	if modTotal == 0.0 {
		points += 25
	}

	//5 points for every two items on the receipt.
	pairs := math.Floor(float64(len(r.Items) / 2))
	points += (int(pairs) * 5)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range r.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	splitDate := strings.Split(r.PurchaseDate, "-")
	day, _ := strconv.Atoi(splitDate[2])
	if day%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")

	purchaseTime, _ := time.Parse("15:04", r.PurchaseTime)
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	return points
}
