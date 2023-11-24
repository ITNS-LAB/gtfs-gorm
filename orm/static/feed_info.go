package static

import "gorm.io/datatypes"

type FeedInfo struct {
	FeedPublisherName *string `gorm:"primaryKey;not null"`
	FeedPublisherUrl  *string `gorm:"not null"`
	FeedLang          *string `gorm:"not null"`
	DefaultLang       *string
	FeedStartDate     *datatypes.Date
	FeedEndDate       *datatypes.Date
	FeedVersion       *string
	FeedContactEmail  *string
	FeedContactUrl    *string
}

func (FeedInfo) TableName() string {
	return "feed_info"
}
