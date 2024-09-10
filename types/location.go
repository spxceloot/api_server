package types

type GeoLocation struct {
	Latitude  int64 `json:"lat"`
	Longitude int64 `json:"long"`
}

type Address struct {
	ID       string `json:"id"`
	City     string `json:"city"`
	Province string `json:"province"`
	Street   string `json:"street"`
	Country  string `json:"country"`
}
