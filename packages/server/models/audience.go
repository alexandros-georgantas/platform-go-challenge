package models

import "database/sql/driver"

type genderType string
type ageGroup string

const (
	MALE   genderType = "MALE"
	FEMALE genderType = "FEMALE"
)

const (
	CHILDREN ageGroup = "0-14"
	YOUTH    ageGroup = "15-24"
	ADULTS   ageGroup = "25-64"
	SENIORS  ageGroup = "65-94"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}

func (gt genderType) Value() (driver.Value, error) {
	return string(gt), nil
}

func (ag *ageGroup) Scan(value interface{}) error {
	*ag = ageGroup(value.([]byte))
	return nil
}

func (ag ageGroup) Value() (driver.Value, error) {
	return string(ag), nil
}

type Audience struct {
	Base                       Base       `gorm:"embedded"`
	Gender                     genderType `gorm:"type:gender_type"`
	CountryOfBirth             string
	AgeGroupe                  ageGroup `gorm:"type:age_group"`
	DailyHoursOnSM             int
	LastMonthNumberOfPurchases int
}
