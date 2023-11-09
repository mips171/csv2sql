package main

import "strconv"

//INSERT INTO `oc_information_description` (`information_id`, `language_id`, `title`, `description`, `meta_title`, `meta_description`, `meta_keyword`) VALUES
// (4, 1, 'About Us', '&lt;p&gt;\r\n	About Us&lt;/p&gt;\r\n', 'About Us', '', ''),
// (5, 1, 'Terms &amp; Conditions', '&lt;p&gt;\r\n	Terms &amp;amp; Conditions&lt;/p&gt;\r\n', 'Terms &amp; Conditions', '', ''),
// (3, 1, 'Privacy Policy', '&lt;p&gt;\r\n	Privacy Policy&lt;/p&gt;\r\n', 'Privacy Policy', '', ''),
// (6, 1, 'Delivery Information', '&lt;p&gt;\r\n	Delivery Information&lt;/p&gt;\r\n', 'Delivery Information', '', ''),
// (7, 1, 'Subscriptions', 'Nowadays, our stores are introducing a new subscription system. Customers can now have the ability to handle customer payments with better features!', 'Subscriptions', '', '');

// CSV
//"Content Type","Content Path","Name","Description 1","Description 2","SEO Meta Description","Sort Order","SEO Page Title","SEO Meta Keywords"

type InformationRecord struct {
	Name     string `csv:"Name"`
	Content  string `csv:"Description 1"`
	Meta     string `csv:"SEO Meta Description"`
	Title    string `csv:"SEO Page Title"`
	Keywords string `csv:"SEO Meta Keywords"`
}

func (r InformationRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "Name":
		return r.Name
	case "Content":
		return r.Content
	case "Meta":
		return r.Meta
	case "Title":
		return r.Title
	case "Keywords":
		return r.Keywords
	default:
		return ""
	}
}

// INSERT INTO `oc_information` (`information_id`, `bottom`, `sort_order`, `status`) VALUES
func GetInformationMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_information",
		ColumnOrder: []string{"bottom", "sort_order", "status"},
		Fields: []FieldMapping{
			{"", "bottom", func(entity Entity) interface{} { return "1" }},
			{"", "sort_order", func(entity Entity) interface{} { return "0" }},
			{"", "status", func(entity Entity) interface{} { return "1" }},
		},
	}
}

// INSERT INTO `oc_information_description` (`information_id`, `language_id`, `title`, `description`, `meta_title`, `meta_description`, `meta_keyword`) VALUES
func GetInformationDescriptionMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_information_description",
		ColumnOrder: []string{"information_id", "language_id", "title", "description", "tag", "meta_title", "meta_description", "meta_keyword"},
		Fields: []FieldMapping{
			{"Name", "information_id", GetInfodTransformation(productIdMapping)},
			{"", "language_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for English
			{"Name", "title", JustUse("Name")},
			{"Content", "description", JustUse("Content")},
			{"Meta", "tag", JustUse("")},
			{"Name", "meta_title", JustUse("Title")},
			{"Meta", "meta_description", JustUse("Meta")},
			{"Keywords", "meta_keyword", JustUse("Keywords")},
		},
	}
}

func GetInfoToStoreMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_information_to_store",
		ColumnOrder: []string{"information_id", "store_id"},
		Fields: []FieldMapping{
			{"Name", "information_id", GetInfodTransformation(productIdMapping)},
			{"", "store_id", func(entity Entity) interface{} { return "0" }},
		},
	}
}

func GetInfodTransformation(productIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		model := entity.GetValue("Name").(string)
		if id, exists := productIdMapping[model]; exists {
			return strconv.Itoa(id)
		}
		return nil
	}
}
