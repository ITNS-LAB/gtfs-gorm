package gtfsschedule

type FeedInfo struct {
	FeedPublisherName string `gorm:"not null"`
	FeedPublisherURL  string `gorm:"not null"`
	FeedLang          string `gorm:"not null"`
	DefaultLang       *string
	FeedStartDate     *string
	FeedEndDate       *string
	FeedVersion       *string
	FeedContactEmail  *string
	FeedContactURL    *string
}
