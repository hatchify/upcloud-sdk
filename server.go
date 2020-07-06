package upcloud

// Tags represents UpCloud server tags
type Tags struct {
	Tag *[]string `json:"tag"`
}
type IPAddress struct {
	Access  string `json:"access"`
	Address string `json:"address"`
	Family  string `json:"family"`
}
type IPAddresses struct {
	IPAddress *[]IPAddress `json:"ip_address"`
}
type Interface struct {
	Index       int          `json:"index"`
	IPAddresses *IPAddresses `json:"ip_addresses"`
	Mac         string       `json:"mac"`
	Network     string       `json:"network"`
	Type        string       `json:"type"`
	Bootable    string       `json:"bootable"`
}
type Interfaces struct {
	Interface *[]Interface `json:"interface"`
}
type Networking struct {
	Interfaces *Interfaces `json:"interfaces"`
}
type StorageDevice struct {
	Action       string `json:"action"`
	Address      string `json:"address"`
	PartOfPlan   string `json:"part_of_plan"`
	Storage      string `json:"storage"`
	StorageSize  int    `json:"storage_size"`
	StorageTitle string `json:"storage_title"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	BootDisk     string `json:"boot_disk"`
}
type StorageDevices struct {
	StorageDevice *[]StorageDevice `json:"storage_device"`
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

// ServerDetails represents all UpCloud detailed server objects
type ServerDetails struct {
	BootOrder            string          `json:"boot_order"`
	CoreNumber           string          `json:"core_number"`
	Firewall             string          `json:"firewall"`
	Host                 int64           `json:"host"`
	Hostname             string          `json:"hostname"`
	IPAddresses          *IPAddresses    `json:"ip_addresses"`
	License              int             `json:"license"`
	MemoryAmount         string          `json:"memory_amount"`
	Networking           *Networking     `json:"networking"`
	NicModel             string          `json:"nic_model"`
	Plan                 string          `json:"plan"`
	PlanIpv4Bytes        string          `json:"plan_ipv4_bytes"`
	PlanIpv6Bytes        string          `json:"plan_ipv6_bytes"`
	SimpleBackup         string          `json:"simple_backup"`
	State                string          `json:"state"`
	StorageDevices       *StorageDevices `json:"storage_devices"`
	Tags                 *Tags           `json:"tags"`
	Timezone             string          `json:"timezone"`
	Title                string          `json:"title"`
	UUID                 string          `json:"uuid"`
	VideoModel           string          `json:"video_model"`
	RemoteAccessEnabled  string          `json:"remote_access_enabled"`
	RemoteAccessType     string          `json:"remote_access_type"`
	RemoteAccessHost     string          `json:"remote_access_host"`
	RemoteAccessPassword string          `json:"remote_access_password"`
	RemoteAccessPort     string          `json:"remote_access_port"`
	Zone                 string          `json:"zone"`
}

// getServersResponse is a response wrapper to match the UpCloud API payload
type getServersResponse struct {
	Servers *Servers `json:"servers"`
}

// serverDetailsWrapper is a response wrapper to match the UpCloud API payload
type serverDetailsWrapper struct {
	ServerDetails *ServerDetails `json:"server"`
}
