package unifi

type unifiLogin struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type unifiAuthorize struct {
	Cmd string `json:"cmd"` // must be "authorize-guest" or "unauthorize-guest"
	Mac string `json:"mac"`
	Minutes int `json:"minutes"`
	AuthorizedBy string `json:"authorized_by"`
}

type UnifiClientResult struct {
	Data []UnifiClient `json:"data"`
}
type UnifiClient struct {
	Expired bool `json:"expired"`
	Mac string `json"mac"`
	ApMac string `json"ap_mac"`
	Start int `json:"start"`
	End int `json:"end"`
	Duration int `json:"duration"`
	Channel int `json:"channel"`
	Hostname string `json:"hostname"`
	AuthorizedBy string `json:"authorized_by"` // voucher/api/none
	UnauthorizedBy string `json:"unauthorized_by"` // voucher/api/none
	Radio string `json:"radio"`
	RoamCount int `json:"roam_count"`
	RxBytes int  `json:"rx_bytes"`
	TxBytes int `json:"tx_bytes"`
}
