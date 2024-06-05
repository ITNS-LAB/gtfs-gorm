package ormstatic

import (
	"database/sql"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/nulldatatypes"
)

type FeedInfo struct {
	FeedPublisherName string `gorm:"primaryKey;not null"`
	FeedPublisherUrl  string `gorm:"not null"`
	FeedLang          string `gorm:"not null"`
	DefaultLang       sql.NullString
	FeedStartDate     nulldatatypes.NullDate
	FeedEndDate       nulldatatypes.NullDate
	FeedVersion       sql.NullString
	FeedContactEmail  sql.NullString
	FeedContactUrl    sql.NullString
}

func (FeedInfo) TableName() string {
	return "feed_info"
}

func (f FeedInfo) GetFeedPublisherName() any {
	return f.FeedPublisherName
}

func (f FeedInfo) GetFeedPublisherUrl() any {
	return f.FeedPublisherUrl
}

func (f FeedInfo) GetFeedLang() any {
	return f.FeedLang
}

func (f FeedInfo) GetDefaultLang() any {
	if f.DefaultLang.Valid {
		return f.DefaultLang.String
	}
	return nil
}

func (f FeedInfo) GetFeedStartDate() any {
	if f.FeedStartDate.Valid {
		return f.FeedStartDate.Date
	}
	return nil
}

func (f FeedInfo) GetFeedEndDate() any {
	if f.FeedEndDate.Valid {
		return f.FeedEndDate.Date
	}
	return nil
}

func (f FeedInfo) GetFeedVersion() any {
	if f.FeedVersion.Valid {
		return f.FeedVersion.String
	}
	return nil
}

func (f FeedInfo) GetFeedContactEmail() any {
	if f.FeedContactEmail.Valid {
		return f.FeedContactEmail.String
	}
	return nil
}

func (f FeedInfo) GetFeedContactUrl() any {
	if f.FeedContactUrl.Valid {
		return f.FeedContactUrl.String
	}
	return nil
}
