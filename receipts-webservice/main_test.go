package main

import (
	"testing"
)

var testReceipts = []Receipt{
	Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			{
				ShortDescription: "Klarbrunn 12-PK 12 FL OZ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	},
	Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
		Total: "9.00",
	},
	Receipt{
		ID:           "jp2n4jh6-e17d-4411-bf3e-87be9e352db6",
		Retailer:     "Walgreens",
		PurchaseDate: "2024-03-20",
		PurchaseTime: "10:30",
		Items: []Item{
			{
				ShortDescription: "Chips",
				Price:            "2.00",
			},
			{
				ShortDescription: "Pepsi",
				Price:            "5.00",
			},
			{
				ShortDescription: "Water",
				Price:            "1.50",
			},
		},
		Total:  "8.00",
		Points: 13,
	},
}

func TestCalcuatePoints1(t *testing.T) {
	expected := 46
	var result int = calculatePoints(testReceipts[0])
	if expected != result {
		t.Fatalf(`calculatePoints(receipt) = %d, expect = %d`, result, expected)
	}
}

func TestCalcuatePoints2(t *testing.T) {
	expected := 29
	var result int = calculatePoints(testReceipts[1])
	if expected != result {
		t.Fatalf(`calculatePoints(receipt) = %d, expect = %d`, result, expected)
	}
}
