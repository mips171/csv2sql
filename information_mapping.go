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

// IINSERT INTO `oc_journal3_blog_post` (`post_id`, `author_id`, `image`, `comments`, `status`, `sort_order`, `date_created`, `date_updated`, `views`, `post_data`) VALUES ('13', '0', ”, '0', '1', '0', '2023-11-09 10:44:26', '2023-11-09 10:44:26', '1', 'null');
func GetInformationMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_journal3_blog_post",
		ColumnOrder: []string{"status", "comments", "date_created", "date_updated"},
		Fields: []FieldMapping{
			{"", "status", func(entity Entity) interface{} { return "1" }},
			{"", "comments", func(entity Entity) interface{} { return "0" }},
			{"", "date_created", func(entity Entity) interface{} { return "2023-11-09 10:44:26" }},
			{"", "date_updated", func(entity Entity) interface{} { return "2023-11-09 11:44:26" }},
		},
	}
}

// INSERT INTO `oc_journal3_blog_post_description` (`post_id`, `language_id`, `name`, `description`, `meta_title`, `meta_keywords`, `meta_robots`, `meta_description`, `keyword`, `tags`) VALUES ('13', '1', 'New Blog Post', '<p>Blog Content</p>', ”, ”, ”, ”, 'SEO Text', ”);
func GetInformationDescriptionMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_journal3_blog_post_description",
		ColumnOrder: []string{"post_id", "language_id", "name", "description", "meta_title", "meta_description", "keyword"},
		Fields: []FieldMapping{
			{"Name", "post_id", GetInfodTransformation(productIdMapping)},
			{"", "language_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for English
			{"Name", "name", JustUse("Name")},
			{"Content", "description", JustUse("Content")},
			{"Name", "meta_title", JustUse("Name")},
			{"Meta", "meta_description", JustUse("Meta")},
			{"Keywords", "keyword", JustUse("Keywords")},
		},
	}
}

func GetInfoToStoreMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_journal3_blog_post_to_store",
		ColumnOrder: []string{"post_id", "store_id"},
		Fields: []FieldMapping{
			{"Name", "post_id", GetInfodTransformation(productIdMapping)},
			{"", "store_id", func(entity Entity) interface{} { return "0" }},
		},
	}
}

// INSERT INTO `oc_journal3_blog_post_to_layout` (`post_id`, `store_id`, `layout_id`) VALUES ('13', '0', '16');
func GetInfoToLayoutMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_journal3_blog_post_to_layout",
		ColumnOrder: []string{"post_id", "store_id", "layout_id"},
		Fields: []FieldMapping{
			{"Name", "post_id", GetInfodTransformation(productIdMapping)},
			{"", "store_id", func(entity Entity) interface{} { return "0" }},
			{"", "layout_id", func(entity Entity) interface{} { return "16" }},
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
