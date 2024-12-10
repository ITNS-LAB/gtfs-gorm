package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
	"gorm.io/datatypes"
)

type FeedInfo struct {
	FeedPublisherName string `gorm:"primaryKey;not null"`
	FeedPublisherUrl  string `gorm:"not null"`
	FeedLang          string `gorm:"not null;default:ja"`
	FeedStartDate     *datatypes.Date
	FeedEndDate       *datatypes.Date
	FeedVersion       *string
}

func (FeedInfo) TableName() string {
	return "feed_info"
}

func ParseFeedInfo(path string) ([]FeedInfo, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open feed_info CSV: %w", err)
	}

	// データを解析して FeedInfo 構造体のスライスを作成
	var feedInfos []FeedInfo
	for i := 0; i < len(df.Records); i++ {
		feedPublisherName, err := df.GetString(i, "feed_publisher_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_publisher_name' at row %d: %w", i, err)
		}

		feedPublisherUrl, err := df.GetString(i, "feed_publisher_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_publisher_url' at row %d: %w", i, err)
		}

		feedLang, err := df.GetString(i, "feed_lang")
		if err != nil || feedLang == "" {
			feedLang = "ja" // デフォルト値
		}

		feedStartDate, err := df.GetDatePtr(i, "feed_start_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_start_date' at row %d: %w", i, err)
		}

		feedEndDate, err := df.GetDatePtr(i, "feed_end_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_end_date' at row %d: %w", i, err)
		}

		feedVersion, err := df.GetStringPtr(i, "feed_version")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_version' at row %d: %w", i, err)
		}

		// FeedInfo 構造体を作成しリストに追加
		feedInfos = append(feedInfos, FeedInfo{
			FeedPublisherName: feedPublisherName,
			FeedPublisherUrl:  feedPublisherUrl,
			FeedLang:          feedLang,
			FeedStartDate:     feedStartDate,
			FeedEndDate:       feedEndDate,
			FeedVersion:       feedVersion,
		})
	}

	return feedInfos, nil
}
