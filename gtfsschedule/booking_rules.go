package gtfsschedule

type BookingRule struct {
	BookingRuleID          string  `gorm:"primaryKey"` // ルールの一意の識別子
	BookingType            int     `gorm:"not null"`   // 予約タイプ
	PriorNoticeDurationMin *int    // 最小通知時間（分）
	PriorNoticeDurationMax *int    // 最大通知時間（分）
	PriorNoticeLastDay     *int    // 予約リクエストを行う便前の最終日
	PriorNoticeLastTime    *string // 便前日の予約リクエストを行う最終時間
	PriorNoticeStartDay    *int    // 予約リクエストを行う便前の最も早い日
	PriorNoticeStartTime   *string // 最も早い日の最も早い時間
	PriorNoticeServiceID   *string // サービス ID
	Message                *string // 通知メッセージ
	PickupMessage          *string // ピックアップメッセージ
	DropOffMessage         *string // ドロップオフメッセージ
	PhoneNumber            *string // 予約のための電話番号
	InfoURL                *string // 予約ルールに関する情報の URL
	BookingURL             *string // 予約リクエスト用の URL
}
