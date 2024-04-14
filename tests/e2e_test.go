package tests

import (
	"bannerService/internal/cache"
	"bannerService/internal/config"
	"bannerService/internal/database"
	"bannerService/internal/deletion"
	"bannerService/internal/entity"
	"bannerService/internal/handler"
	"bannerService/internal/repo"
	"bannerService/internal/service"
	"bannerService/internal/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func setupTest() (*config.Config, *sqlx.DB, *echo.Echo) {
	utils.CreateLogger()
	cfg := getTestConfig()

	time.Sleep(5 * time.Second)

	db, err := database.New(&cfg.DB)
	if err != nil {
		utils.Logger.Error("Error: failed to init db connection " + err.Error())
		os.Exit(1)
	}

	cache := cache.New(&cfg.Cache)
	deletionQueue := deletion.CreateQueue()
	router := echo.New()

	repos := repo.New(db.DB)
	services := service.New(repos, cache, deletionQueue)
	handlers := handler.New(services, &cfg.Tokens)

	handlers.Route(router)

	return cfg, db.DB, router
}

func finishTest(db *sqlx.DB) {
	dropAllTables(db)
	_ = db.DB.Close()
}

func Test_VersionsHistory(t *testing.T) {
	cfg, db, router := setupTest()
	defer finishTest(db)

	createBannersTable(db)
	createHistoryTable(db)

	// Create banner
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/banner/", bytes.NewBufferString(historyCreateBanner))
	req.Header.Set(echo.HeaderAuthorization, cfg.Tokens.Admin)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	router.ServeHTTP(w, req)

	data, err := io.ReadAll(w.Result().Body)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}

	var bannerId handler.ResponseId
	err = json.Unmarshal(data, &bannerId)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}

	// Patch created banner 3 times
	for _, patchBody := range historyPatches {
		w = httptest.NewRecorder()
		req = httptest.NewRequest("PATCH", fmt.Sprintf("/banner/%d", bannerId.BannerId),
			bytes.NewBufferString(patchBody))
		req.Header.Set(echo.HeaderAuthorization, cfg.Tokens.Admin)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		router.ServeHTTP(w, req)
	}

	// Get history of patched banner
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", fmt.Sprintf("/banner_history/%d", bannerId.BannerId), nil)
	req.Header.Set(echo.HeaderAuthorization, cfg.Tokens.Admin)
	router.ServeHTTP(w, req)

	data, err = io.ReadAll(w.Result().Body)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}

	var history entity.BannerHistory
	err = json.Unmarshal(data, &history)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}

	// Set version num historySetVersion
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST",
		fmt.Sprintf("/banner_history/%d?set_version=%d", bannerId.BannerId, historySetVersion),
		nil)
	req.Header.Set(echo.HeaderAuthorization, cfg.Tokens.Admin)
	router.ServeHTTP(w, req)

	// Get updated banner
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", fmt.Sprintf("/banner/?banner_id=%d", bannerId.BannerId), nil)
	req.Header.Set(echo.HeaderAuthorization, cfg.Tokens.Admin)
	router.ServeHTTP(w, req)

	data, err = io.ReadAll(w.Result().Body)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}

	var result []entity.BannerInfo
	err = json.Unmarshal(data, &result)
	if err != nil {
		utils.Logger.Error("VersionsHistory test error: " + err.Error())
		os.Exit(1)
	}
	if len(result) < 1 {
		utils.Logger.Error(fmt.Sprintf("VersionsHistory test error: can't find banner with id %d", bannerId.BannerId))
		os.Exit(1)
	}

	// Compare banner from history and banner after setting that version
	var bannerFromHistory entity.BannerVersion
	for _, bannerVer := range history.Versions {
		if bannerVer.Version == historySetVersion {
			bannerFromHistory = bannerVer
			break
		}
	}

	resultedBanner := result[0]

	assert.Equal(t, bannerFromHistory.Content, resultedBanner.Content)
	assert.Equal(t, bannerFromHistory.FeatureId, resultedBanner.FeatureId)
	assert.Equal(t, bannerFromHistory.TagIds, resultedBanner.TagIds)
	assert.Equal(t, bannerFromHistory.IsActive, resultedBanner.IsActive)
}
