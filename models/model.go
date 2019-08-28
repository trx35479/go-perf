package models

type Products struct {
	Data struct {
		Products []struct {
			AdditionalInformation struct {
				EligibilityUri    string `json:"eligibilityUri"`
				FeesAndPricingUri string `json:"feesAndPricingUri"`
				OverviewUri       string `json:"overviewUri"`
				TermsUri          string `json:"termsUri"`
			} `json:"additionalInformation"`
			ApplicationUri  string      `json:"applicationUri"`
			Brand           string      `json:"brand"`
			Description     string      `json:"description"`
			EffectiveFrom   string      `json:"effectiveFrom"`
			IsTailored      interface{} `json:"isTailored"`
			LastUpdated     string      `json:"lastUpdated"`
			Name            string      `json:"name"`
			ProductCategory string      `json:"productCategory"`
			ProductId       string      `json:"productId"`
		} `json:"products"`
	} `json:"data"`
}
