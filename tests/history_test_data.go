package tests

var (
	historyCreateBanner = `{
  		"tag_ids": [1, 3, 4],
  		"feature_id": 11,
  		"content":  {"data": "test_data"},
  		"is_active": true
	}`

	historyPatches = []string{`{
  		"tag_ids": [1],
  		"feature_id": 11,
  		"content":  {"data": "test_data_patch1"},
  		"is_active": true
	}`,
		`{
  		"tag_ids": [3],
  		"feature_id": 11,
  		"content":  {"data": "test_data_patch2"},
  		"is_active": true
	}`,
		`{
  		"tag_ids": [1, 3, 4],
  		"feature_id": 12,
  		"content":  {"data": "test_data_patch3"},
  		"is_active": true
	}`,
	}

	historySetVersion = 2
)
