package event_logic

import (
	"fmt"
	db "server/db/sqlc"
	"shared/scenarios"
)

func HandleScanResponse(payload scenarios.ScanResult) {
	db := db.GetConnection()
	fmt.Println("=> db : ", db)
	fmt.Println("HandleScanResponse : ", payload)
}
