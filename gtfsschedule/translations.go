package gtfsschedule

type Translation struct {
	TableName   string  `gorm:"not null"`
	FieldName   string  `gorm:"not null"`
	Language    string  `gorm:"not null"`
	Translation string  `gorm:"not null"`
	RecordID    *string `json:"record_id,omitempty"`     // 翻訳するフィールドに対応するレコード ID
	RecordSubID *string `json:"record_sub_id,omitempty"` // サブ ID（必要な場合）
	FieldValue  *string `json:"field_value,omitempty"`   // 翻訳を適用する特定の値
}
