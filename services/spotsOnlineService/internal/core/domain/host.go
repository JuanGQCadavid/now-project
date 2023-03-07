package domain

type HostInfo struct {
	HostId   string `json:"-"`
	HostName string `json:"host_name"`
}
