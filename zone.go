package upcloud

// Zone represents UpCloud zone
type Zone struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Public      string `json:"public"`
}

// Zones represents all UpCloud zones
type Zones struct {
	Zone *[]Zone `json:"zone"`
}

// getZonesResponse is a response wrapper to match the UpCloud API payload
type getZonesResponse struct {
	Zones *Zones `json:"zones"`
}
