package scenarios

import (
	"errors"
	"fmt"
	"net"
	"server/utils"
	work_dto "shared/dto"
	"shared/shared_utils"
	"sync"
	"time"
)

type ScanResult struct {
	UserId          int32           `json:"user_id"`
	Date            time.Time       `json:"date"`
	PortAnalysisMap PortResponseMap `json:"port_analysis"`
}

type PortResponse struct {
	Num    int  `json:"num"`
	IsOpen bool `json:"is_open"`
}

type Analysis struct {
	ip           string
	portResponse PortResponse
}

type PortResponseMap map[string][]PortResponse

func Scan(payload work_dto.PortTestScenarioRequest) (ScanResult, error) {
	fmt.Println("=> scan beginning")
	utils.PrettyDisplay("SCAN : ", payload)
	if payload.PortRange.Min > payload.PortRange.Max {
		return ScanResult{}, errors.New("portRange min > max")
	}
	portResponses := PortResponseMap{}

	var wg sync.WaitGroup
	resultChan := make(chan Analysis)
	done := make(chan int)
	goMerger(portResponses, resultChan, done)
	utils.PrettyDisplay("SCAN22 : ", payload)
	addresses, err := shared_utils.ExtractIpAddressesFromRange(payload.IPRange)

	utils.PrettyDisplay("ADDRESSES : ", addresses)
	if err != nil {
		return ScanResult{}, err
	}

	for _, address := range addresses {
		for i := payload.PortRange.Min; i <= payload.PortRange.Max; i++ {
			wg.Add(1)
			goScanUnit(address, i, resultChan, &wg)
		}
	}

	wg.Wait()
	done <- 1
	return ScanResult{
		UserId:          payload.UserId,
		Date:            time.Now(),
		PortAnalysisMap: portResponses,
	}, nil
}

func goMerger(portResponses PortResponseMap, resultChan chan Analysis, done <-chan int) {
	go func() {
		for {
			select {
			case <-done:
				return
			case response := <-resultChan:

				if response.portResponse.IsOpen {
					fmt.Println("==> ", response)
					portResponses[response.ip] = append(portResponses[response.ip], response.portResponse)
				}
			}
		}
	}()
}

func goScanUnit(address string, i int, resultChan chan Analysis, wg *sync.WaitGroup) {
	go func() {
		port := i
		defer wg.Done()
		fullAddress := fmt.Sprintf("%s:%d", address, port)
		d := net.Dialer{Timeout: time.Second * 4}
		_, err := d.Dial("tcp", fullAddress)
		portResp := PortResponse{
			Num: port,
		}
		if err == nil {
			fmt.Printf("==> ", port)
			portResp.IsOpen = true
		}
		resultChan <- Analysis{
			ip:           address,
			portResponse: portResp,
		}
	}()
}
