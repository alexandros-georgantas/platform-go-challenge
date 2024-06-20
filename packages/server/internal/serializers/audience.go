package serializers

type Audience struct {
	ID                         int    `json:"id"`
	Gender                     string `json:"gender"`
	CountryOfBirth             string `json:"countryIfBirth"`
	AgeGroupe                  string `json:"ageGroup"`
	DailyHoursOnSocialMedia    int    `json:"dailyHoursOnSocialMedia"`
	LastMonthNumberOfPurchases int    `json:"lastMonthNumberOfPurchases"`
}
