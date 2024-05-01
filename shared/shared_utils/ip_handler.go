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
	return err == nil
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

func ExtractIpAddressesFromRange(ipRange IpRange) ([]string, error) {
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
	ips = append(ips, convertIntIpToString(currentIp))
	for {
		isEqual, err := IsIpsEquals(currentIp, targetIp)
		if err != nil {
			return []string{}, err
		}
		if isEqualStartIp, _ := IsIpsEquals(currentIp, [4]int{255, 255, 255, 255}); isEqualStartIp || isEqual {
			break
		}
		fmt.Println("=> currentIpd : ", currentIp)

		IncrementIp(&currentIp)
		ips = append(ips, convertIntIpToString(currentIp))

	}
	return ips, nil
}

func IncrementIp(ip *[4]int) {
	index := 3

	for {
		if ip[index] == 255 {
			ip[index] = 0
			if index == 0 {
				break
			}
			index--
			continue
		} else {
			ip[index]++
		}
		break
	}
	fmt.Println("inside : ", ip)
}

func IsIpsEquals(ip1 [4]int, ip2 [4]int) (bool, error) {
	isValid := IsIpValid(fmt.Sprintf("%d.%d.%d.%d", ip1[0], ip1[1], ip1[2], ip1[3]))
	if !isValid {
		return false, errors.New("ip1 is not valid")
	}
	isValid = IsIpValid(fmt.Sprintf("%d.%d.%d.%d", ip2[0], ip2[1], ip2[2], ip2[3]))
	if !isValid {
		return false, errors.New("ip2 is not valid")
	}
	for i := 0; i < 4; i++ {
		if ip1[i] != ip2[i] {
			return false, nil
		}
	}
	return true, nil
}
