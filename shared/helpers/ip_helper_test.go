package helpers

import (
	"fmt"
	"log"
	shared_utils "shared/utils"
	"testing"

	"github.com/stretchr/testify/assert"
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
	for _, ipRangeCase := range ipRanges {
		res, err := ExtractIpAddressesFromRange(ipRangeCase.in)
		if err != nil {
			log.Fatal(err.Error())
		}
		shared_utils.PrettyDisplay("ips : ", res)
		assert.ElementsMatch(t, ipRangeCase.out, res)
	}
}

func BenchmarkIsIpValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipCase := range ipCases {
			IsIpValid(ipCase.ip)
		}
	}
}
