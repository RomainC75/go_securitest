package shared_utils

import (
	"fmt"
	"log"
	"server/utils"
	work_dto "shared/dto"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIpValid(t *testing.T) {
	for _, ipCase := range ipCases {
		result := IsIpValid(ipCase.ip)
		require.Equal(t, ipCase.expected, result)
	}
}

func TestConvertStringIpToInts(t *testing.T) {
	for _, strIpCase := range strIpCases {
		intIp, err := ConvertStringIpToInts(strIpCase.strIp)
		fmt.Println(strIpCase.resIp, intIp, err)
		if strIpCase.errorExpected {
			require.Error(t, err)
		} else {
			fmt.Println(strIpCase.resIp, intIp)
			require.NoError(t, err)
			require.Equal(t, strIpCase.resIp, intIp)

		}
	}
}

func TestIncrementIp(t *testing.T) {
	for _, testCase := range incrementIpCases {
		var buffer [4]int
		copy(buffer[:], testCase.ip[:])
		IncrementIp(&testCase.ip)
		require.ElementsMatch(t, testCase.ip, testCase.expected)
	}
}

func TestIsIpsEquals(t *testing.T) {
	for i, isIpsEqualsCase := range isIpsEqualsCases {
		fmt.Println("=> ", i)
		isEqual, err := IsIpsEquals(isIpsEqualsCase.ip1, isIpsEqualsCase.ip2)
		if isIpsEqualsCase.err != nil {
			require.Error(t, err)
			require.Equal(t, err.Error(), isIpsEqualsCase.err.Error())
			require.Equal(t, isEqual, false)
		} else {
			require.Equal(t, isIpsEqualsCase.expect, isEqual)
			require.NoError(t, err)
		}
	}
}

func TestExtractAddressesFromRange(t *testing.T) {
	ipRange := work_dto.IpRange{
		IpMin: "255.255.255.250",
		IpMax: "0.0.0.2",
	}
	res, err := ExtractIpAddressesFromRange(ipRange)
	if err != nil {
		log.Fatal("xx")
	}
	utils.PrettyDisplay("ips : ", res)

}

func BenchmarkIsIpValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipCase := range ipCases {
			IsIpValid(ipCase.ip)
		}
	}
}
