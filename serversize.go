package upcloud

// ServerSize represents UpCloud server size
type ServerSize struct {
	CoreNumber   string `json:"core_number"`
	MemoryAmount string `json:"memory_amount"`
}

// ServerSizes represents all UpCloud server sizes
type ServerSizes struct {
	ServerSize *[]ServerSize `json:"server_size"`
}

// getServerSizesResponse is a response wrapper to match the UpCloud API payload
type getServerSizesResponse struct {
	ServerSizes *ServerSizes `json:"server_sizes"`
}
