package tests

import "bannerService/internal/config"

func getTestConfig() *config.Config {
	return &config.Config{
		HTTP: config.HTTP{
			Port: "8090",
		},
		DB: config.DB{
			Name:     "postgres-test",
			Host:     "localhost",
			Port:     "5438",
			User:     "postgres",
			Password: "postgres",
		},
		Cache: config.Cache{
			Host: "localhost",
			Port: "6380",
		},
		Tokens: config.Tokens{
			User:  "user_test_token",
			Admin: "admin_test_token",
		},
	}
}
