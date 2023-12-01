package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseFeedInfo(path string) []orm.FeedInfo {
	var feedInfos []orm.FeedInfo
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		feedInfos = append(feedInfos, orm.FeedInfo{
			FeedPublisherName: dataframe.IsBlank(df.GetElement("feed_publisher_name")),
			FeedPublisherUrl:  dataframe.IsBlank(df.GetElement("feed_publisher_url")),
			FeedLang:          dataframe.IsBlank(df.GetElement("feed_lang")),
			DefaultLang:       dataframe.IsBlank(df.GetElement("default_lang")),
			FeedStartDate:     dataframe.ParseDate(df.GetElement("feed_start_date")),
			FeedEndDate:       dataframe.ParseDate(df.GetElement("feed_end_date")),
			FeedVersion:       dataframe.IsBlank(df.GetElement("feed_version")),
			FeedContactEmail:  dataframe.IsBlank(df.GetElement("feed_contact_email")),
			FeedContactUrl:    dataframe.IsBlank(df.GetElement("feed_contact_url")),
		})
	}
	return feedInfos
}
