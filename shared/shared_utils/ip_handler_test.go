package shared_utils

import (
	"fmt"
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

func BenchmarkIsIpValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipCase := range ipCases {
			IsIpValid(ipCase.ip)
		}
	}
}
