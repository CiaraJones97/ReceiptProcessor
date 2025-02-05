package main

import (
	"testing"
)

var receipt1 Receipt = Receipt{
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
}

var receipt2 Receipt = Receipt{
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
}

func TestCalcuatePoints1(t *testing.T) {
	expected := 28
	var result int = calculatePoints(receipt1)
	if expected != result {
		t.Fatalf(`calculatePoints(receipt) = %d, expect = %d`, result, expected)
	}
}

func TestCalcuatePoints2(t *testing.T) {
	expected := 109
	var result int = calculatePoints(receipt2)
	if expected != result {
		t.Fatalf(`calculatePoints(receipt) = %d, expect = %d`, result, expected)
	}
}
