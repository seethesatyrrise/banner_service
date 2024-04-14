package tests

var (
	filterCreateBanner = []string{`{
  		"tag_ids": [1, 3, 4],
  		"feature_id": 11,
  		"content":  {"data": "test_data1"},
  		"is_active": true
	}`,
		`{
  		"tag_ids": [5, 6, 7],
  		"feature_id": 11,
  		"content":  {"data": "test_data2"},
  		"is_active": true
	}`,
		`{
  		"tag_ids": [1, 6, 7],
  		"feature_id": 12,
  		"content":  {"data": "test_data3"},
  		"is_active": true
	}`,
	}
)
