package models

type Insight struct {
	Base Base `gorm:"embedded"`
	Text string
}
