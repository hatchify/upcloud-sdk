package upcloud

// Tags represents UpCloud server tags
type Tags struct {
	Tag *[]string `json:"tag,omitempty"`
}
type IPAddress struct {
	Access  string `json:"access,omitempty"`
	Address string `json:"address,omitempty"`
	Family  string `json:"family,omitempty"`
}
type IPAddresses struct {
	IPAddress *[]IPAddress `json:"ip_address,omitempty"`
}
type Interface struct {
	Index       int          `json:"index,omitempty"`
	IPAddresses *IPAddresses `json:"ip_addresses,omitempty"`
	Mac         string       `json:"mac,omitempty"`
	Network     string       `json:"network,omitempty"`
	Type        string       `json:"type,omitempty"`
	Bootable    string       `json:"bootable,omitempty"`
}
type Interfaces struct {
	Interface *[]Interface `json:"interface,omitempty"`
}
type Networking struct {
	Interfaces *Interfaces `json:"interfaces,omitempty"`
}
type StorageDevice struct {
	Action       string `json:"action,omitempty"`
	Address      string `json:"address,omitempty"`
	PartOfPlan   string `json:"part_of_plan,omitempty"`
	Storage      string `json:"storage,omitempty"`
	StorageSize  int    `json:"storage_size,omitempty"`
	StorageTitle string `json:"storage_title,omitempty"`
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	BootDisk     string `json:"boot_disk,omitempty"`
}
type StorageDevices struct {
	StorageDevice *[]StorageDevice `json:"storage_device,omitempty"`
}

// Server represents UpCloud server
type Server struct {
	CoreNumber    string `json:"core_number,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	License       int    `json:"license,omitempty"`
	MemoryAmount  string `json:"memory_amount,omitempty"`
	Plan          string `json:"plan,omitempty"`
	PlanIvp4Bytes string `json:"plan_ivp4_bytes,omitempty,omitempty"`
	PlanIpv6Bytes string `json:"plan_ipv6_bytes,omitempty,omitempty"`
	State         string `json:"state,omitempty"`
	Tags          *Tags  `json:"tags,omitempty"`
	Title         string `json:"title,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Zone          string `json:"zone,omitempty"`
}

// Servers represents all UpCloud servers
type Servers struct {
	Server *[]Server `json:"server,omitempty"`
}

// ServerDetails represents all UpCloud detailed server objects
type ServerDetails struct {
	BootOrder            string          `json:"boot_order,omitempty"`
	CoreNumber           string          `json:"core_number,omitempty"`
	Firewall             string          `json:"firewall,omitempty"`
	Host                 int64           `json:"host,omitempty"`
	Hostname             string          `json:"hostname,omitempty"`
	IPAddresses          *IPAddresses    `json:"ip_addresses,omitempty"`
	License              int             `json:"license,omitempty"`
	MemoryAmount         string          `json:"memory_amount,omitempty"`
	Networking           *Networking     `json:"networking,omitempty"`
	NicModel             string          `json:"nic_model,omitempty"`
	Plan                 string          `json:"plan,omitempty"`
	PlanIpv4Bytes        string          `json:"plan_ipv4_bytes,omitempty"`
	PlanIpv6Bytes        string          `json:"plan_ipv6_bytes,omitempty"`
	SimpleBackup         string          `json:"simple_backup,omitempty"`
	State                string          `json:"state,omitempty"`
	StorageDevices       *StorageDevices `json:"storage_devices,omitempty"`
	Tags                 *Tags           `json:"tags,omitempty"`
	Timezone             string          `json:"timezone,omitempty"`
	Title                string          `json:"title,omitempty"`
	UUID                 string          `json:"uuid,omitempty"`
	VideoModel           string          `json:"video_model,omitempty"`
	RemoteAccessEnabled  string          `json:"remote_access_enabled,omitempty"`
	RemoteAccessType     string          `json:"remote_access_type,omitempty"`
	RemoteAccessHost     string          `json:"remote_access_host,omitempty"`
	RemoteAccessPassword string          `json:"remote_access_password,omitempty"`
	RemoteAccessPort     string          `json:"remote_access_port,omitempty"`
	Zone                 string          `json:"zone,omitempty"`
}

// getServersResponse is a response wrapper to match the UpCloud API payload
type getServersResponse struct {
	Servers *Servers `json:"servers,omitempty"`
}

// serverDetailsWrapper is a response wrapper to match the UpCloud API payload
type serverDetailsWrapper struct {
	ServerDetails *ServerDetails `json:"server,omitempty"`
}

// StartServer optional parameters for starting servers
type StartServer struct {
	Host      int64 `json:"host,omitempty"`
	AvoidHost int64 `json:"avoid_host,omitempty"`
}
type startServerRequest struct {
	StartServer StartServer `json:"server"`
}

// StopServer optional parameters for stopping servers
type StopServer struct {
	StopType string `json:"stop_type,omitempty"`
	Timeout  string `json:"timeout,omitempty"` //1-600 range
}
type stopServerRequest struct {
	StopServer StopServer `json:"stop_server"`
}