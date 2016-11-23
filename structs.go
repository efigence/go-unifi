package unifi

type unifiLogin struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type unifiAuthorize struct {
	Cmd string `json:"cmd"` // must be "authorize-guest" or "unauthorize-guest"
	Mac string `json:"mac"`
	Minutes string `json:"minutes"`
	AuthorizedBy string `json:"authorized_by"`
}

type UnifiClientResult struct {
	Data []UnifiClient `json:"data"`
}
type UnifiClient struct {
	Expired bool `json:"expired"`
	Mac string `json"mac"`
	Start int `json:"start"`
	End int `json:"end"`
	Hostname string `json:"hostname"`
	AuthorizedBy string `json:"authorized_by"`
	RxBytes int `json:"rx_bytes"`
	TxBytes int `json:"tx_bytes"`
}
