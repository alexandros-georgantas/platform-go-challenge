package serializers

type Chart struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	HorizontalAxisLabel string    `json:"xAxisLabel"`
	VerticalAxisLabel   string    `json:"yAxisLabel"`
	Data                []float64 `json:"data"`
}
