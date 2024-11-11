package tx_types

import db "server/db/sqlc"

type CreatedScanTx struct {
	Scan      db.Scan      `json:"scan"`
	IpRange   db.IpRange   `json:"ip_range"`
	PortRange db.PortRange `json:"port_range"`
}
