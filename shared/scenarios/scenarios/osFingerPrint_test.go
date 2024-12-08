package scenarios

import (
	"log"
	shared_dto "shared/dto"
	"testing"
)

var str = "192.168.0.30"

var testCases = []struct {
	in  shared_dto.IpRange
	out string
}{
	{
		in: shared_dto.IpRange{
			IpMin: "192.168.20",
			IpMax: &str,
		},
		out: "xxxx",
	},
}

func TestOsFingerPrint(t *testing.T) {
	fp := NewFingerPrint(testCases[0].in)
	res, _ := fp.Run()
	log.Println("--> res : ", res)

}
