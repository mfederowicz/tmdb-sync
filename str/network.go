package str

// Network is the response shape for the network details endpoint.
type Network struct {
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// NetworkAlternativeNames is the response shape for the network alternative-names endpoint.
type NetworkAlternativeNames struct {
	ID      int64             `json:"id"`
	Results []AlternativeName `json:"results"`
}
