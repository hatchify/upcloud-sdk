package upcloud

import "time"

type getStoragesResponse struct {
	Storages *Storages `json:"storages"`
}
type Storage struct {
	Access     string    `json:"access"`
	License    float64   `json:"license"`
	Size       int       `json:"size"`
	State      string    `json:"state"`
	Tier       string    `json:"tier,omitempty"`
	Title      string    `json:"title"`
	Type       string    `json:"type"`
	UUID       string    `json:"uuid"`
	Zone       string    `json:"zone"`
	Created    time.Time `json:"created,omitempty"`
	Origin     string    `json:"origin,omitempty"`
	PartOfPlan string    `json:"part_of_plan,omitempty"`
}
type Storages struct {
	Storage *[]Storage `json:"storage"`
}
