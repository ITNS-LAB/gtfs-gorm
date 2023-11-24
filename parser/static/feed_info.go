package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFeedInfo(path string) []static.FeedInfo {
	var feedInfos []static.FeedInfo
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		feedInfos = append(feedInfos, static.FeedInfo{
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
