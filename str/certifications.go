package str

// Certifications is the response of GET /certification/movie/list and
// GET /certification/tv/list - an object keyed by ISO 3166-1 country code.
type Certifications struct {
	Certifications map[string][]Certification `json:"certifications"`
}

// Certification is one certification rating within a country.
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}
