package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

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

		feedPublisherURL, err := df.GetString(i, "feed_publisher_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_publisher_url' at row %d: %w", i, err)
		}

		feedLang, err := df.GetString(i, "feed_lang")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_lang' at row %d: %w", i, err)
		}

		defaultLang, err := df.GetStringPtr(i, "default_lang")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'default_lang' at row %d: %w", i, err)
		}

		feedStartDate, err := df.GetStringPtr(i, "feed_start_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_start_date' at row %d: %w", i, err)
		}

		feedEndDate, err := df.GetStringPtr(i, "feed_end_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_end_date' at row %d: %w", i, err)
		}

		feedVersion, err := df.GetStringPtr(i, "feed_version")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_version' at row %d: %w", i, err)
		}

		feedContactEmail, err := df.GetStringPtr(i, "feed_contact_email")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_contact_email' at row %d: %w", i, err)
		}

		feedContactURL, err := df.GetStringPtr(i, "feed_contact_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_contact_url' at row %d: %w", i, err)
		}

		// FeedInfo 構造体を作成しリストに追加
		feedInfos = append(feedInfos, FeedInfo{
			FeedPublisherName: feedPublisherName,
			FeedPublisherURL:  feedPublisherURL,
			FeedLang:          feedLang,
			DefaultLang:       defaultLang,
			FeedStartDate:     feedStartDate,
			FeedEndDate:       feedEndDate,
			FeedVersion:       feedVersion,
			FeedContactEmail:  feedContactEmail,
			FeedContactURL:    feedContactURL,
		})
	}

	return feedInfos, nil
}

type FeedInfoGeom struct {
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

func ParseFeedInfoGeom(path string) ([]FeedInfoGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open feed_info CSV: %w", err)
	}

	// データを解析して FeedInfo 構造体のスライスを作成
	var feedInfos []FeedInfoGeom
	for i := 0; i < len(df.Records); i++ {
		feedPublisherName, err := df.GetString(i, "feed_publisher_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_publisher_name' at row %d: %w", i, err)
		}

		feedPublisherURL, err := df.GetString(i, "feed_publisher_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_publisher_url' at row %d: %w", i, err)
		}

		feedLang, err := df.GetString(i, "feed_lang")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_lang' at row %d: %w", i, err)
		}

		defaultLang, err := df.GetStringPtr(i, "default_lang")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'default_lang' at row %d: %w", i, err)
		}

		feedStartDate, err := df.GetStringPtr(i, "feed_start_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_start_date' at row %d: %w", i, err)
		}

		feedEndDate, err := df.GetStringPtr(i, "feed_end_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_end_date' at row %d: %w", i, err)
		}

		feedVersion, err := df.GetStringPtr(i, "feed_version")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_version' at row %d: %w", i, err)
		}

		feedContactEmail, err := df.GetStringPtr(i, "feed_contact_email")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_contact_email' at row %d: %w", i, err)
		}

		feedContactURL, err := df.GetStringPtr(i, "feed_contact_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'feed_contact_url' at row %d: %w", i, err)
		}

		// FeedInfo 構造体を作成しリストに追加
		feedInfos = append(feedInfos, FeedInfoGeom{
			FeedPublisherName: feedPublisherName,
			FeedPublisherURL:  feedPublisherURL,
			FeedLang:          feedLang,
			DefaultLang:       defaultLang,
			FeedStartDate:     feedStartDate,
			FeedEndDate:       feedEndDate,
			FeedVersion:       feedVersion,
			FeedContactEmail:  feedContactEmail,
			FeedContactURL:    feedContactURL,
		})
	}

	return feedInfos, nil
}
