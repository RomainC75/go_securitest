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
	Address string         `json:"address"`
	Date    time.Time      `json:"date"`
	Ports   []PortResponse `json:"ports"`
}

type PortResponse struct {
	Num    int  `json:"num"`
	IsOpen bool `json:"is_open"`
}

func GetString(result ScanResult) ([]byte, error) {
	openPorts := []PortResponse{}
	for _, pRes := range result.Ports {
		if pRes.IsOpen {
			openPorts = append(openPorts, pRes)
		}
	}
	result.Ports = openPorts
	res, err := json.Marshal(result)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
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
		Address: address,
		Date:    time.Now(),
		Ports:   portResponses,
	}, nil
}

func goMerger(portResponses *[]PortResponse, resultChan chan PortResponse, done <-chan int) {
	go func() {
		for {
			select {
			case <-done:
				return
			case response := <-resultChan:
				// fmt.Println("res : ", response.IsOpen, response.Num)
				if response.IsOpen {
					*portResponses = append(*portResponses, response)
				}
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
			Num: port,
		}
		if err == nil {
			fmt.Printf("==> ", port)
			portResp.IsOpen = true
		}
		resultChan <- portResp
	}()
}
