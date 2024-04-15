package scenarios

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type ScanResult struct {
	address string         `json:"address"`
	ports   []PortResponse `json:"ports"`
}

type PortResponse struct {
	num    int
	isOpen bool
}

func GetString(result ScanResult) string {
	openPorts := []PortResponse{}
	for _, pRes := range result.ports {
		if pRes.isOpen {
			openPorts = append(openPorts, pRes)
		}
	}
	result.ports = openPorts
	res, _ := json.Marshal(result)
	return string(res)
}

func Scan(address string, portMin int, portMax int) (ScanResult, error) {
	if portMax < portMin {
		return ScanResult{}, errors.New("portRange min > max")
	}
	portResponses := []PortResponse{}

	var wg sync.WaitGroup
	resultChan := make(chan PortResponse)
	done := make(chan int)
	goMerger(&portResponses, resultChan, done)

	for i := portMin; i <= portMax; i++ {
		fmt.Println("=>", i)
		wg.Add(1)
		goScanUnit(address, i, resultChan, &wg)
	}

	wg.Wait()
	done <- 1
	return ScanResult{
		address: address,
		ports:   portResponses,
	}, nil
}

func goMerger(portResponses *[]PortResponse, resultChan chan PortResponse, done <-chan int) {
	go func() {
		for {
			select {
			case <-done:
				return
			case response := <-resultChan:
				fmt.Println("res : ", response.isOpen, response.num)
				*portResponses = append(*portResponses, response)
			}
		}
	}()
}

func goScanUnit(address string, i int, resultChan chan PortResponse, wg *sync.WaitGroup) {
	go func() {
		port := i
		defer wg.Done()
		address := fmt.Sprintf("%s:%d", address, port)
		d := net.Dialer{Timeout: time.Second * 4}
		_, err := d.Dial("tcp", address)
		portResp := PortResponse{
			num: port,
		}
		if err == nil {
			fmt.Printf("==> ", port)
			portResp.isOpen = true
		}
		resultChan <- portResp
	}()
}
