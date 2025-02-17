package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type BookingRule struct {
	BookingRuleId                 string      `gorm:"primaryKey"` // ルールの一意の識別子
	BookingType                   int         `gorm:"not null"`   // 予約タイプ
	PriorNoticeDurationMin        *int        // 最小通知時間（分）
	PriorNoticeDurationMax        *int        // 最大通知時間（分）
	PriorNoticeLastDay            *int        // 予約リクエストを行う便前の最終日
	PriorNoticeLastTime           *string     // 便前日の予約リクエストを行う最終時間
	PriorNoticeStartDay           *int        // 予約リクエストを行う便前の最も早い日
	PriorNoticeStartTime          *string     // 最も早い日の最も早い時間
	PriorNoticeServiceId          *string     // サービス ID
	Message                       *string     // 通知メッセージ
	PickupMessage                 *string     // ピックアップメッセージ
	DropOffMessage                *string     // ドロップオフメッセージ
	PhoneNumber                   *string     // 予約のための電話番号
	InfoURL                       *string     // 予約ルールに関する情報の URL
	BookingURL                    *string     // 予約リクエスト用の URL
	StopTimesPickupBookingRuleId  []StopTimes `gorm:"foreignKey:PickupBookingRuleId;references:BookingRuleId"`
	StopTimesDropOffBookingRuleId []StopTimes `gorm:"foreignKey:DropOffBookingRuleid;references:BookingRuleId "`
}

func ParseBookingRule(path string) ([]BookingRule, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open BookingRule CSV: %w", err)
	}

	// データを解析して BookingRule 構造体のスライスを作成
	var bookingRules []BookingRule
	for i := 0; i < len(df.Records); i++ {
		bookingRuleID, err := df.GetString(i, "booking_rule_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'booking_rule_id' at row %d: %w", i, err)
		}

		bookingType, err := df.GetInt(i, "booking_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'booking_type' at row %d: %w", i, err)
		}

		priorNoticeDurationMin, err := df.GetIntPtr(i, "prior_notice_duration_min")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_duration_min' at row %d: %w", i, err)
		}

		priorNoticeDurationMax, err := df.GetIntPtr(i, "prior_notice_duration_max")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_duration_max' at row %d: %w", i, err)
		}

		priorNoticeLastDay, err := df.GetIntPtr(i, "prior_notice_last_day")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_last_day' at row %d: %w", i, err)
		}

		priorNoticeLastTime, err := df.GetStringPtr(i, "prior_notice_last_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_last_time' at row %d: %w", i, err)
		}

		priorNoticeStartDay, err := df.GetIntPtr(i, "prior_notice_start_day")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_start_day' at row %d: %w", i, err)
		}

		priorNoticeStartTime, err := df.GetStringPtr(i, "prior_notice_start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_start_time' at row %d: %w", i, err)
		}

		priorNoticeServiceID, err := df.GetStringPtr(i, "prior_notice_service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'prior_notice_service_id' at row %d: %w", i, err)
		}

		message, err := df.GetStringPtr(i, "message")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'message' at row %d: %w", i, err)
		}

		pickupMessage, err := df.GetStringPtr(i, "pickup_message")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'pickup_message' at row %d: %w", i, err)
		}

		dropOffMessage, err := df.GetStringPtr(i, "drop_off_message")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'drop_off_message' at row %d: %w", i, err)
		}

		phoneNumber, err := df.GetStringPtr(i, "phone_number")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'phone_number' at row %d: %w", i, err)
		}

		infoURL, err := df.GetStringPtr(i, "info_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'info_url' at row %d: %w", i, err)
		}

		bookingURL, err := df.GetStringPtr(i, "booking_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'booking_url' at row %d: %w", i, err)
		}

		// BookingRule 構造体を作成しリストに追加
		bookingRules = append(bookingRules, BookingRule{
			BookingRuleId:          bookingRuleID,
			BookingType:            bookingType,
			PriorNoticeDurationMin: priorNoticeDurationMin,
			PriorNoticeDurationMax: priorNoticeDurationMax,
			PriorNoticeLastDay:     priorNoticeLastDay,
			PriorNoticeLastTime:    priorNoticeLastTime,
			PriorNoticeStartDay:    priorNoticeStartDay,
			PriorNoticeStartTime:   priorNoticeStartTime,
			PriorNoticeServiceId:   priorNoticeServiceID,
			Message:                message,
			PickupMessage:          pickupMessage,
			DropOffMessage:         dropOffMessage,
			PhoneNumber:            phoneNumber,
			InfoURL:                infoURL,
			BookingURL:             bookingURL,
		})
	}

	return bookingRules, nil
}
