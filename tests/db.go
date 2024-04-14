package tests

import "github.com/jmoiron/sqlx"

var (
	createBannerTableQuery = `create table if not exists banners
(
    banner_id  serial primary key,
    content    jsonb     not null,
    tag_ids    bigint[]  not null,
    feature_id bigint    not null,
    is_active  boolean,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);`

	createHistoryTableQuery = `create table if not exists banners_history
(
    id         serial primary key,
    banner_id  int      not null,
    content    jsonb    not null,
    tag_ids    bigint[] not null,
    feature_id bigint   not null,
    is_active  boolean
);`

	dropBannerTableQuery = `DROP TABLE IF EXISTS banners`

	dropHistoryTableQuery = `DROP TABLE IF EXISTS banners_history`
)

func createBannersTable(db *sqlx.DB) {
	db.Exec(createBannerTableQuery)
}

func createHistoryTable(db *sqlx.DB) {
	db.Exec(createHistoryTableQuery)
}

func dropAllTables(db *sqlx.DB) {
	db.Exec(dropHistoryTableQuery)
	db.Exec(dropBannerTableQuery)
}
