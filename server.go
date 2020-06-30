package upcloud

// Tags represents UpCloud server tags
type Tags struct {
	Tag []string `json:"tag"`
}

// Server represents UpCloud server
type Server struct {
	CoreNumber    string `json:"core_number"`
	Hostname      string `json:"hostname"`
	License       int    `json:"license"`
	MemoryAmount  string `json:"memory_amount"`
	Plan          string `json:"plan"`
	PlanIvp4Bytes string `json:"plan_ivp4_bytes,omitempty"`
	PlanIpv6Bytes string `json:"plan_ipv6_bytes,omitempty"`
	State         string `json:"state"`
	Tags          *Tags  `json:"tags"`
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
	Zone          string `json:"zone"`
}

// Servers represents all UpCloud servers
type Servers struct {
	Server *[]Server `json:"server"`
}

// getServersResponse is a response wrapper to match the UpCloud API payload
type getServersResponse struct {
	Servers *Servers `json:"servers"`
}
