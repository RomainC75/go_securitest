package scenarios

import (
	"fmt"
	"testing"
)

var testCasesTR = []struct {
	targetAddress string
	out           string
}{
	{
		targetAddress: "google.com",
		out:           "xxxx",
	},
}

func TestTraceroute(t *testing.T) {
	fp := NewTraceroute(testCasesTR[0].targetAddress)
	// res, _ := fp.Run()
	// log.Println("--> res : ", res)

	res, err := fp.Run()
	if err != nil {
		fmt.Println("--> err : ", err.Error())
	}
	fmt.Println("--> res : ", res)
}
