package scenarios

import (
	"errors"
	"fmt"
	"net"
	shared_dto "shared/dto"
	"shared/helpers"
	shared_utils "shared/utils"
	"sync"
	"time"
)

type Scan struct {
	payload shared_dto.PortTestScenario
}

func NewScan(payload shared_dto.PortTestScenario) *Scan {
	return &Scan{
		payload: payload,
	}
}

type ScanResult struct {
	Date            time.Time       `json:"date"`
	PortAnalysisMap PortResponseMap `json:"port_analysis"`
}

type PortResponse struct {
	Num    int32 `json:"num"`
	IsOpen bool  `json:"is_open"`
}

type Analysis struct {
	ip           string
	portResponse PortResponse
}

type PortResponseMap map[string][]PortResponse

func (s *Scan) Check() error {
	if s.payload.PortRange.Max != nil && s.payload.PortRange.Min > *s.payload.PortRange.Max {
		return errors.New("portRange min > max")
	}
	return nil
}

func (s *Scan) Run() (interface{}, error) {
	portResponses := PortResponseMap{}

	var wg sync.WaitGroup
	resultChan := make(chan Analysis)
	done := make(chan int)
	goMerger(portResponses, resultChan, done)

	shared_utils.PrettyDisplay("SCAN222 : ", s.payload)

	addresses, err := helpers.ExtractIpAddressesFromRange(s.payload.IPRange)

	fmt.Println(addresses)

	if err != nil {
		return ScanResult{}, err
	}

	var portMax int32
	if s.payload.PortRange.Max != nil {
		portMax = *s.payload.PortRange.Max
	} else {
		portMax = s.payload.PortRange.Min
	}

	for _, address := range addresses {
		for i := s.payload.PortRange.Min; i <= portMax; i++ {
			wg.Add(1)
			goScanUnit(address, i, resultChan, &wg)
		}
	}

	wg.Wait()
	done <- 1
	return ScanResult{
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
					portResponses[response.ip] = append(portResponses[response.ip], response.portResponse)
				}
			}
		}
	}()
}

func goScanUnit(address string, i int32, resultChan chan Analysis, wg *sync.WaitGroup) {
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
			portResp.IsOpen = true
		}
		resultChan <- Analysis{
			ip:           address,
			portResponse: portResp,
		}
	}()
}
