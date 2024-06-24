package types

type OCRTemplate struct {
	ID          uint        `gorm:"primaryKey"`
	Nationality string      `gorm:"uniqueIndex:idx_nationality" json:"nationality"`
	Bounds      []Rectangle `gorm:"foreignKey:TemplateID" json:"bounds"`
}

type Rectangle struct {
	ID          uint    `gorm:"primaryKey"`
	TemplateID  uint    `json:"template_id"`
	TopLeft     float64 `json:"top_left"`
	TopRight    float64 `json:"top_right"`
	BottomLeft  float64 `json:"bottom_left"`
	BottomRight float64 `json:"bottom_right"`
}
