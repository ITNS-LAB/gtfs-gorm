package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFeedInfo(path string) ([]ormstatic.FeedInfo, error) {
	var feedInfos []ormstatic.FeedInfo
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return feedInfos, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.FeedInfo{}, err
		}

		feedPublisherName, err := dataframe.ParseString(df.GetElement("feed_publisher_name"))
		if err != nil {
			return []ormstatic.FeedInfo{}, err
		}

		feedPublisherUrl, err := dataframe.ParseString(df.GetElement("feed_publisher_url"))
		if err != nil {
			return []ormstatic.FeedInfo{}, err
		}

		feedLang, err := dataframe.ParseString(df.GetElement("feed_lang"))
		if err != nil {
			return []ormstatic.FeedInfo{}, err
		}

		feedStartDate, err := dataframe.ParseNullDataTypesDate(df.GetElement("feed_start_date"))
		feedEndDate, err := dataframe.ParseNullDataTypesDate(df.GetElement("feed_end_date"))

		feedInfos = append(feedInfos, ormstatic.FeedInfo{
			FeedPublisherName: feedPublisherName,
			FeedPublisherUrl:  feedPublisherUrl,
			FeedLang:          feedLang,
			DefaultLang:       df.GetElement("default_lang"),
			FeedStartDate:     feedStartDate,
			FeedEndDate:       feedEndDate,
			FeedVersion:       df.GetElement("feed_version"),
			FeedContactEmail:  df.GetElement("feed_contact_email"),
			FeedContactUrl:    df.GetElement("feed_contact_url"),
		})
	}
	return feedInfos, nil
}
