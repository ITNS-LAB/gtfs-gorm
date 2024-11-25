package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseFeedInfo(path string) ([]gtfsjp.FeedInfo, error) {
	var feedInfos []gtfsjp.FeedInfo
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return feedInfos, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		feedInfos = append(feedInfos, gtfsjp.FeedInfo{
			FeedPublisherName: util.IsBlank(df.GetElement("feed_publisher_name")),
			FeedPublisherUrl:  util.IsBlank(df.GetElement("feed_publisher_url")),
			FeedLang:          util.IsBlank(df.GetElement("feed_lang")),
			DefaultLang:       util.IsBlank(df.GetElement("default_lang")),
			FeedStartDate:     util.ParseDate(df.GetElement("feed_start_date")),
			FeedEndDate:       util.ParseDate(df.GetElement("feed_end_date")),
			FeedVersion:       util.IsBlank(df.GetElement("feed_version")),
			FeedContactEmail:  util.IsBlank(df.GetElement("feed_contact_email")),
			FeedContactUrl:    util.IsBlank(df.GetElement("feed_contact_url")),
		})
	}
	return feedInfos, nil
}
