package types

type OCRTemplate struct {
	ID          uint        `gorm:"primaryKey"`
	Nationality string      `gorm:"uniqueIndex:idx_nationality" json:"nationality"`
	Bounds      []Rectangle `gorm:"foreignKey:TemplateID" json:"bounds"`
}

type Rectangle struct {
	ID          uint  `gorm:"primaryKey"`
	TemplateID  uint  `json:"template_id"`
	TopLeft     Point `gorm:"embedded;embeddedPrefix:top_left_" json:"-"`
	TopRight    Point `gorm:"embedded;embeddedPrefix:top_right_" json:"-"`
	BottomLeft  Point `gorm:"embedded;embeddedPrefix:bottom_left_" json:"-"`
	BottomRight Point `gorm:"embedded;embeddedPrefix:bottom_right_" json:"-"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
