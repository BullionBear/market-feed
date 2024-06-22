package pgdb

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgDatabase struct {
	DB *gorm.DB
}

func NewPgDatabase(host string, port int, user, password, dbName, sslMode, timeZone string) (*PgDatabase, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host, port, user, password, dbName, sslMode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PgDatabase{DB: db}, nil
}

func (pg *PgDatabase) QueryKlines(startTime, endTime int64, offset, limit int) ([]PlaybackKline, error) {
	var records []PlaybackKline
	result := pg.DB.Where("open_time BETWEEN ? AND ?", startTime, endTime).
		Offset(offset).Limit(limit).Find(&records)
	return records, result.Error
}
