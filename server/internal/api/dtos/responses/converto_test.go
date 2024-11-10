package dto_res

import (
	"database/sql"
	db "server/db/sqlc"
	shared_dto "shared/dto"
	"shared/types"
	shared_utils "shared/utils"
	"testing"
	"time"
)

var t1, err1 = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

var data = db.ListScansByUserRow{
	ID:        2,
	UserID:    3,
	Scenario:  1,
	CreatedAt: t1,
	UpdatedAt: t1,
	ID_2: sql.NullInt64{
		Int64: 1,
		Valid: false,
	},
	ScanID: sql.NullInt64{
		Int64: 2,
		Valid: false,
	},
	RangeMin: sql.NullInt32{
		Int32: 3,
		Valid: false,
	},
	RangeMax: sql.NullInt32{
		Int32: 4,
		Valid: false,
	},
	IsUnique: sql.NullBool{
		Bool:  false,
		Valid: false,
	},
	ID_3: sql.NullInt64{
		Int64: 5,
		Valid: false,
	},
	ScanID_2: sql.NullInt64{
		Int64: 6,
		Valid: false,
	},
	IpMin: sql.NullString{
		String: "sdfrrrrrrr",
		Valid:  false,
	},
	IpMax: sql.NullString{
		String: "aaaaaaaaaar",
		Valid:  false,
	},
	IsUnique_2: sql.NullBool{
		Bool:  false,
		Valid: false,
	},
}

var data2 = &shared_dto.FullPortTestScenarioReq{
	ScenarioBase: shared_dto.ScenarioBase{
		BasicData: "scenario1",
	},
	PortTestScenario: shared_dto.PortTestScenario{
		IPRange: shared_dto.IpRange{
			IpMin: "127.0.0.1",
			IpMax: types.SNullString{
				NullString: sql.NullString{
					String: "127.0.0.2",
					Valid:  true,
				},
			},
			Unique: false,
		},
		PortRange: shared_dto.PortRange{
			Min: 23,
			Max: types.SNullInt32{
				NullInt32: sql.NullInt32{
					Int32: 45,
					Valid: true,
				},
			},
		},
	},
}

func TestToDbNull(t *testing.T) {
	res := convertStruct(data2)
	shared_utils.PrettyDisplay("res : ", res)
}

// func TestToFullPortTestScenarioReq(t *testing.T) {
// 	fmt.Println("errr : ", err1)
// 	res := ToFullPortTestScenarioReq(data)
// 	shared_utils.PrettyDisplay(" res : ", res)
// }
