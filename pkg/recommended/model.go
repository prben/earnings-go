package recommended


type RecommendedSymbols struct {
	Finance struct {
		Result []struct {
			RecommendedSymbols []Recommended `json:"recommendedSymbols"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"finance"`
}

type Recommended struct {
	Symbol string `json:"symbol"`
}
