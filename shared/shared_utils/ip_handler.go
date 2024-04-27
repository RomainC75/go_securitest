package shared_utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type IpRange struct {
	IpMin string `json:"ip_min" validate:"required"`
	IpMax string `json:"ip_max" validate:"required,ip"`
}

func IsIpValid(ip string) bool {
	re := regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
	if !re.Match([]byte(ip)) {
		return false
	}
	_, err := ConvertStringIpToInts(ip)
	if err != nil {
		return false
	}
	return true
}

func ConvertStringIpToInts(ip string) ([4]int, error) {
	fmt.Println("=> ", ip)
	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		fmt.Println("=> error 1 ", ip)
		return [4]int{}, errors.New("invalid Ip")
	}
	res := [4]int{}
	for index, number := range numbers {
		intNum, err := strconv.Atoi(number)
		if err != nil || intNum < 0 || intNum > 255 {
			fmt.Println("=> error 2 ", ip)
			return [4]int{}, errors.New("invalid Ip")
		}
		res[index] = intNum
	}
	fmt.Println("=> OOK ", ip)
	return res, nil
}

func convertIntIpToString(ip [4]int) string {
	separatedIp := [4]string{}
	for i := 0; i < 4; i++ {
		separatedIp[i] = strconv.Itoa(ip[i])
	}
	return strings.Join(separatedIp[:], ".")
}

func ExtractBytesFromIpString(ipRange IpRange) ([]string, error) {
	if !IsIpValid(ipRange.IpMin) || !IsIpValid(ipRange.IpMax) {
		return []string{}, errors.New("invalid Ip")
	}
	currentIp, err := ConvertStringIpToInts(ipRange.IpMin)
	if err != nil {
		return []string{}, err
	}
	targetIp, err := ConvertStringIpToInts(ipRange.IpMax)
	if err != nil {
		return []string{}, err
	}
	ips := []string{}

	for !IsIpsEquals(currentIp, targetIp) {

		if currentIp[3] == 255 {
			IncrementIp(currentIp)
			continue
		}

		currentIp[3]++
		ips = append(ips, convertIntIpToString(currentIp))
	}
	return ips, nil
}

func IncrementIp(ip [4]int) {
	index := 3
	for {
		ip[index] = 0
		if ip[index-1] == 255 && index > 0 {
			index--
			continue
		}
		break
	}
}

func IsIpsEquals(ip1 [4]int, ip2 [4]int) bool {
	for i := 0; i < 4; i++ {
		if ip1[i] != ip2[i] {
			return false
		}
	}
	return true
}

// func IpRangeExtractor(ipRange IpRange) []string {

// }
