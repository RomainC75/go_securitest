package shared_utils

import "errors"

var ipCases = []struct {
	ip       string
	expected bool
}{
	{"123.123.123.123", true},
	{"123.123.0.123", true},
	{"123.255.0.123", true},
	{"123.123.523.123", false},
	{"123.123.123", false},
	{"123.123.123.123.2", false},
	{"123.123..123", false},
}

var strIpCases = []struct {
	strIp         string
	resIp         [4]int
	errorExpected bool
}{
	{"123.123.123.123", [4]int{123, 123, 123, 123}, false},
	{"123.123.0.123", [4]int{123, 123, 0, 123}, false},
	{"123.255.0.123", [4]int{123, 255, 0, 123}, false},
	{"123.123.523.123", [4]int{}, true},
	{"123.123.123", [4]int{}, true},
	{"123.123.123.123.2", [4]int{}, true},
	{"123.123..123", [4]int{}, true},
}

var incrementIpCases = []struct {
	ip       [4]int
	expected [4]int
}{
	{[4]int{123, 123, 123, 123}, [4]int{123, 123, 123, 124}},
	{[4]int{0, 0, 0, 0}, [4]int{0, 0, 0, 1}},
	{[4]int{123, 123, 255, 123}, [4]int{123, 123, 255, 124}},
	{[4]int{123, 123, 255, 255}, [4]int{123, 124, 0, 0}},
	{[4]int{123, 255, 255, 255}, [4]int{124, 0, 0, 0}},
	{[4]int{255, 255, 255, 255}, [4]int{0, 0, 0, 0}},
}

var isIpsEqualsCases = []struct {
	ip1    [4]int
	ip2    [4]int
	expect bool
	err    error
}{
	{[4]int{123, 123, 123, 123}, [4]int{123, 123, 123, 123}, true, nil},
	{[4]int{123, 123, 123, 123}, [4]int{123, 123, 123, 124}, false, nil},
	{[4]int{0, 0, 0, 0}, [4]int{0, 0, 0, 0}, true, nil},
	{[4]int{0, 0, 0, 1}, [4]int{0, 0, 0, 0}, false, nil},
	{[4]int{2, 5677, 0, 1}, [4]int{0, 0, 0, 0}, false, errors.New("ip1 is not valid")},
	{[4]int{2, 3, 0, 1}, [4]int{0, 0, 12345, 0}, false, errors.New("ip2 is not valid")},
}
