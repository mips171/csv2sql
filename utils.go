package main

import (
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
