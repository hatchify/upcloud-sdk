package upcloud

// Plan represents UpCloud plan
type Plan struct {
	CoreNumber       int    `json:"core_number"`
	MemoryAmount     int    `json:"memory_amount"`
	Name             string `json:"name"`
	PublicTrafficOut int    `json:"public_traffic_out"`
	StorageSize      int    `json:"storage_size"`
	StorageTier      string `json:"storage_tier"`
}

// Plans represents all UpCloud plans
type Plans struct {
	Plan *[]Plan `json:"plan"`
}

// getPlansResponse is a response wrapper to match the UpCloud API payload
type getPlansResponse struct {
	Plans *Plans `json:"plans"`
}
