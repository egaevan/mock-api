package domain

import "time"

type TNPS struct {
	OverallTNPSFirstPeriod  TNPSScore `json:"overall_tnps_first_period"`
	OverallTNPSSecondPeriod TNPSScore `json:"overall_tnps_second_period"`
	TNPSChannel             TNPSScore `json:"tnps_channel"`
	TNPSJourney             TNPSScore `json:"tnps_journey"`
	TNPSTrends              TNPSScore `json:"tnps_trends"`
	TUPCTPChannel           TNPSScore `json:"tup_ctp_channel"`
	TUPDigitalJourney       TNPSScore `json:"tup_digital_journey"`
}

type TNPSScore struct {
	TotalRespondent int
	Promotor        int
	Passive         int
	Detractor       int
}

type TNPSMetric struct {
	ChannelName string
	Interact    int
	Choose      int
	Consume     int
	GetHelp     int
	Pay         int
	Leave       int
}

type Channel struct {
	Id          int
	ChannelName string
}

type Journey struct {
	Id          int
	JourneyName string
	Category    string
}

type CustomerType struct {
	Id       int
	TypeName string
}

type Area struct {
	Id       int
	Name     string
	Location []Location `json:"location"`
}

type Location struct {
	Id           int
	LocationName string
}

type SentimentData struct {
	SentimentTimeline  []SentimentScoreTimeline `json:"sentiment_timeline"`
	SentimentBreakdown SentimentScore           `json:"sentiment_breakdown"`
	SentimentScore     int                      `json:"sentiment_score"`
}

type SentimentScore struct {
	Total    int
	Positive int
	Neutral  int
	Mixed    int
	Negative int
}

type SentimentScoreTimeline struct {
	Total    int
	Positive int
	Neutral  int
	Mixed    int
	Negative int
	Date     time.Time
}

type Sentiment struct {
	Description   string
	ChannelName   string
	Author        string
	Msisdn        string
	Location      string
	Sentiment     string
	ARPU          int
	Los           string
	TelcoBehavior string
}

type SentimentDetailProfile struct {
	UserProfile      []SocialMedia      `json:"user_profile"`
	UserData         UserData           `json:"user_data"`
	SentimentProfile []SentimentProfile `json:"sentiment_profile"`
}

type SocialMedia struct {
	SocialMediaName string
	Username        string
	Author          string
	Followers       int
	Following       int
	Friends         int
}

type UserData struct {
	Msisdn        string
	ARPU          int
	Los           string
	TelcoBehavior string
	Location      string
	NPS           int
	CSI           string
	TNPS          int
	CTPs          string
}

type SentimentProfile struct {
	Description string
	ChannelName string
	Date        time.Time
	Author      string
	Location    string
	Sentiment   string
	ThreadUrl   string
}

type GeoMapPage struct {
	GeoMap    []GeoMap
	GeoByArea GeoByArea
}

type GeoMap struct {
	TNPSScore TNPSScore
	Location  Location
}

type GeoByArea struct {
	Area       Area
	TopicRatio []TopicRatio
}

type TopicRatio struct {
	Topic Topic
	Ratio int
}

type Source struct {
	Id          int
	Name        string
	Image       string
	Description string
}

type Topic struct {
	Id   int
	Name string
}

type SentimentCategory struct {
	Id   int
	Name string
}

type Word struct {
	Name  string
	Total int
}

type TopicData struct {
	Id             int
	Name           string
	CreatedBy      string
	SentimentScore SentimentScore
}

type CreateCustomTopic struct {
	Name             string   `json:"name" binding:"required"`
	IncludedKeyword  []string `json:"included_keyword" binding:"required"`
	ExcludedKeyword  []string `json:"excluded_keyword" binding:"required"`
	MinimumWordCount int      `json:"minimum_word_count" binding:"required"`
	Translation      bool     `json:"translation"`
}

type TopicDetail struct {
	TotalPhrase    int
	SentimentScore SentimentScore
}
