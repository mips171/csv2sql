package main

import (
	"os"
	"strings"
	"time"
)

func GetUserGroupID(groupName string) string {
	switch groupName {
	case "Retail":
		return "1"
	case "Trade":
		return "2"
	default:
		return "1"
	}
}

func GetNewsletterStatus(subscriberStatus string) string {
	if subscriberStatus == "y" {
		return "1"
	}
	return "0"
}

func GetStatus(activeStatus string) string {
	if activeStatus == "y" {
		return "1"
	}
	return "0"
}

func GetDateAdded() string {
	dateAdded := time.Now().Format("2006-01-02 15:04:05")
	return dateAdded
}

func DoNothing() func(string, string) string {
	return func(value string, _ string) string { return TransformIdentity(value) }
}

func GetFirstName(value string, email string) string {
	if value != "" {
		return value
	}
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return "firstname"
}

func GetLastName(value string, email string) string {
	if value != "" {
		return value
	}
	parts := strings.Split(email, "@")
	if len(parts) > 1 {
		domainParts := strings.Split(parts[1], ".")
		if len(domainParts) > 0 {
			return domainParts[0]
		}
	}
	return "surname"
}

func MapBrandToManufacturerID(brand string, nothing string) string {
	// Example: map brand names to manufacturer IDs
	// brandToID := map[string]string{
	// 	"Brand1": "1",
	// 	"Brand2": "2",
	// 	// ... and so on for other brands
	// }

	// return brandToID[brand]
	return "1"
}

func MapImageFilePath(sku string, email string) string {
	// Define the base path where the images will be stored
	basePath := "catalog/images/products/"

	// Generate the possible file paths for jpg and png formats
	jpgPath := basePath + sku + ".jpg"
	pngPath := basePath + sku + ".png"

	// In this example, we are assuming the jpg format as the default
	// If png images are more common in your dataset, you can check for the png path first
	if _, err := os.Stat(jpgPath); err == nil {
		return jpgPath
	} else if _, err := os.Stat(pngPath); err == nil {
		return pngPath
	} else {
		// If neither file exists, return an empty string or handle the error as needed
		return "catalog/journal3/placeholder-1100x1100.png"
	}
}
