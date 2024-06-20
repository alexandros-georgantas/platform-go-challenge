package serializers

type Asset struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	RelatedType string `json:"relatedType"`
	RelatedID   int    `json:"relatedId"`
	Chart       *Chart
	Audience    *Audience
	Insight     *Insight
}

type AddToFavorites struct {
	ID int `json:"id"`
}

type UpdateDescription struct {
	Description string `json:"description"`
}
