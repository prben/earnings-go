package recommended

import "github.com/jinzhu/gorm"

type RecommendedSymbols struct {
	Finance struct {
		Result []struct {
			RecommendedSymbols []Recommended `json:"recommendedSymbols"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"finance"`
}

type Recommended struct {
	gorm.Model

	Symbol string `json:"symbol"`
}
