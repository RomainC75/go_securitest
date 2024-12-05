package dto_res

import (
	"database/sql"
	"fmt"
	db "server/db/sqlc"
	tx_types "server/db/sqlc/txTypes"
	shared_utils "shared/utils"
	"testing"
	"time"
)

var testCase = tx_types.CreatedScanTx{
	Scan: db.Scan{
		ID:        1,
		UserID:    1,
		Scenario:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	IpRange: db.IpRange{
		ID:     1,
		ScanID: 1,
		IpMin:  "127.0.0.1",
		IpMax: sql.NullString{
			String: "127.0.0.2",
			Valid:  true,
		},
	},
	PortRange: db.PortRange{
		ID:       1,
		ScanID:   1,
		RangeMin: 1,
		RangeMax: sql.NullInt32{
			Int32: 2,
			Valid: true,
		},
	},
}

func TestTransformToNullString(t *testing.T) {
	res := TransformNullStrings(testCase)
	if res != nil {
		t.Error("res should not be nil")
	}
	t.Log("res", res)
	fmt.Println("lskjdflkjsd")
	shared_utils.PrettyDisplay("res", res)
}
