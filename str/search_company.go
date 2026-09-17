package str

// SearchCompanyResult is a single company entry within a search/company response.
type SearchCompanyResult struct {
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// SearchCompanies is a paginated list of companies, as returned by GET /search/company.
type SearchCompanies struct {
	Page         int                   `json:"page"`
	Results      []SearchCompanyResult `json:"results"`
	TotalPages   int                   `json:"total_pages"`
	TotalResults int                   `json:"total_results"`
}
