package model

type EventDetail struct {
	EventID uint  `gorm:"column:event_id"`
	Rate    uint8 `gorm:"column:rate"`
	Event   Event `gorm:"foreignKey:EventID"`
}
