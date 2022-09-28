package api

type Node struct {
	Name      string
	Addresses Addresses
	Cuda      interface{} `json:"cuda"`
}

type Addresses struct {
	InternalIP string
	Hostname   string
}
