package main

import (
	"fmt"
	"strings"
)

// This EntityRecord struct maps from CSV to a struct
// Categories are denoted by a ; separated list. Category names are separated by >.
// 2 Examples:
// 1. Vehicle, Caravan & Marine > Marine Products > Mounts & Accessories;Brackets & Clamps > Pole Clamps & Brackets
// 2. Masts & Towers;Masts & Towers > Free-Standing Towers
type CategoryRecord struct {
	Category string `csv:"Category"`
	// Add other fields as required
}

// Implement the Entity interface for ProductRecord
func (c CategoryRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "Category":
		return c.Category
	default:
		return ""
	}
}

// TRUNCATE TABLE `oc_category`;
// INSERT INTO `oc_category` (`category_id`, `image`, `parent_id`, `top`, `column`, `sort_order`, `status`, `date_added`, `date_modified`) VALUES
// (25, '', 0, 1, 1, 3, 1, '2009-01-31 01:04:25', '2011-05-30 12:14:55');

func GetCategoryMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_category",
		ColumnOrder: []string{"category_id", "image", "parent_id", "top", "column", "sort_order", "status", "date_added", "date_modified"},
		Fields: []FieldMapping{
			{"Category", "category_id", MapCategoryToCategoryId},
			{"", "image", EmptyString()},
			{"", "parent_id", EmptyString()},
			{"", "top", EmptyString()},
			{"", "column", EmptyString()},
			{"", "sort_order", func(entity Entity) interface{} { return "0" }},
			{"", "status", func(entity Entity) interface{} { return "1" }},
			{"", "date_added", func(entity Entity) interface{} { return "2023-09-20" }},
			{"", "date_modified", func(entity Entity) interface{} { return "2023-09-20" }},
		},
	}
}

// TRUNCATE TABLE `oc_category_description`;
// INSERT INTO `oc_category_description` (`category_id`, `language_id`, `name`, `description`, `meta_title`, `meta_description`, `meta_keyword`) VALUES
// (28, 1, 'Monitors', '', 'Monitors', '', ''),

// This will define the levels of category hierarchy
// TRUNCATE TABLE `oc_category_path`;
// INSERT INTO `oc_category_path` (`category_id`, `path_id`, `level`) VALUES
// (25, 25, 0),

// Product can be assigned to multiple categories, which are denoted by a ; separated list. Category names are separated by >.
// Vehicle, Caravan & Marine > Marine Products > Mounts & Accessories;Brackets & Clamps > Pole Clamps & Brackets
// TRUNCATE TABLE `oc_product_to_category`;
// INSERT INTO `oc_product_to_category` (`product_id`, `category_id`) VALUES ('28', '20');

func GetProductToCategoryMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_to_category",
		ColumnOrder: []string{"product_id", "category_id"},
		Fields: []FieldMapping{
			{"Model", "product_id", GetProductIdTransformation(productIdMapping)},
			{"Category", "category_id", MapCategoryToCategoryId},
		},
	}
}

// TRUNCATE TABLE `oc_category_to_store`;
// will always be 0 the default store
// INSERT INTO `oc_category_to_store` (`category_id`, `store_id`) VALUES
// (17, 0),
// Map out our actual SQL for ProductToStore
func GetCategoryToStoreMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_category_to_store",
		ColumnOrder: []string{"category_id", "store_id"},
		Fields: []FieldMapping{
			{"Category", "category_id", GetProductIdTransformation(productIdMapping)},
			{"", "store_id", func(entity Entity) interface{} { return "0" }}, // Default store value
		},
	}
}

// Create and maintain a map of Category to category_id

var (
	categoryIDCounter       = 1
	categoryToCategoryIdMap = make(map[string]int)
)

// Recursively parses the category string and assigns unique category IDs.
func ParseAndAssignCategoryIDs(categoryHierarchy string) ([]int, []string) {
	segments := strings.Split(categoryHierarchy, ">")
	var ids []int
	var parsedCategories []string

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		parsedCategories = append(parsedCategories, segment)

		var categoryId int
		if existingId, ok := categoryToCategoryIdMap[segment]; !ok {
			categoryId = categoryIDCounter
			categoryToCategoryIdMap[segment] = categoryId
			categoryIDCounter++
		} else {
			categoryId = existingId
		}

		ids = append(ids, categoryId)
	}

	return ids, parsedCategories
}

func MapCategoryToCategoryId(entity Entity) interface{} {
	c, _ := entity.GetValue("Category").(string)
	hierarchies := strings.Split(c, ";")
	var lastID int

	for _, hierarchy := range hierarchies {
		ids, _ := ParseAndAssignCategoryIDs(hierarchy)
		lastID = ids[len(ids)-1]
	}

	return lastID
}

func GenerateCategorySQLStatements(category CategoryRecord, imgPath string) []string {
	hierarchies := strings.Split(category.Category, ";")
	var statements []string

	for _, hierarchy := range hierarchies {
		ids, parsedCategories := ParseAndAssignCategoryIDs(hierarchy)

		parentID := 0
		for index, cat := range parsedCategories {
			currentID := ids[index]

			// SQL for oc_category
			statements = append(statements, fmt.Sprintf("INSERT IGNORE INTO `oc_category` (`category_id`, `image`, `parent_id`, `top`, `column`, `sort_order`, `status`, `date_added`, `date_modified`) VALUES (%d, '%s', %d, 1, 1, 0, 1, '2023-09-20', '2023-09-20');", currentID, imgPath, parentID))

			// SQL for oc_category_description
			statements = append(statements, fmt.Sprintf("INSERT IGNORE INTO `oc_category_description` (`category_id`, `language_id`, `name`, `description`, `meta_title`, `meta_description`, `meta_keyword`) VALUES (%d, 1, '%s', '', '%s', '', '');", currentID, cat, cat))

			// SQL for oc_category_path (handle hierarchy levels)
			for j := 0; j <= index; j++ {
				statements = append(statements, fmt.Sprintf("INSERT IGNORE INTO `oc_category_path` (`category_id`, `path_id`, `level`) VALUES (%d, %d, %d);", currentID, ids[j], j))
			}

			statements = append(statements, fmt.Sprintf("INSERT IGNORE INTO `oc_category_to_store` (`category_id`, `store_id`) VALUES (%d, 0);", currentID))

			parentID = currentID
		}
	}
	return statements
}
