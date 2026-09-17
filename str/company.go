package str

// ParentCompany is the parent-company summary nested inside a Company, when present.
type ParentCompany struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// Company is the response shape for the company details endpoint.
type Company struct {
	Description   string        `json:"description"`
	Headquarters  string        `json:"headquarters"`
	Homepage      string        `json:"homepage"`
	ID            int64         `json:"id"`
	LogoPath      string        `json:"logo_path"`
	Name          string        `json:"name"`
	OriginCountry string        `json:"origin_country"`
	ParentCompany ParentCompany `json:"parent_company"`
}

// AlternativeName is a single alternative name for a company.
type AlternativeName struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// AlternativeNames is the response shape for the company alternative-names endpoint.
type AlternativeNames struct {
	ID      int64             `json:"id"`
	Results []AlternativeName `json:"results"`
}
