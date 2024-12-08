package scenarios

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type TargetInfos struct {
	ip   string
	name string
}

type STraceroute struct {
	targetAddress string
	routes        []TargetInfos
}

func NewTraceroute(targetAddress string) *STraceroute {
	return &STraceroute{
		targetAddress: targetAddress,
		routes:        []TargetInfos{},
	}
}

func (fp *STraceroute) Run() ([]TargetInfos, error) {
	ttl := 1
	for {
		addr, ok, pass, err := fp.send(ttl)
		fmt.Println("--> addr : ", addr, ok, pass, ttl)
		if err != nil && !ok && !pass {
			return []TargetInfos{}, err
		}
		if !ok && !pass {
			fp.routes = append(fp.routes, addr)
		} else if !ok && pass {
			fp.routes = append(fp.routes, TargetInfos{
				ip:   "",
				name: "",
			})
		} else {
			return fp.routes, nil
		}
		ttl++
	}

}

func (fp *STraceroute) send(ttl int) (TargetInfos, bool, bool, error) {
	address := fp.targetAddress

	conn, err := net.Dial("ip4:icmp", address)
	if err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Create ICMP connection
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("error listening on ICMP: %v", err)
	}
	defer c.Close()

	// Set TTL
	ipConn := c.IPv4PacketConn()
	if err := ipConn.SetTTL(ttl); err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to set TTL: %v", err)
	}

	// Prepare ICMP message
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	// Marshal the message
	wb, err := wm.Marshal(nil)
	if err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to marshal message: %v", err)
	}

	// Send the packet
	dst, err := net.ResolveIPAddr("ip4", address)
	if err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to resolve address: %v", err)
	}

	// Write to the destination
	if _, err := c.WriteTo(wb, dst); err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to write to destination: %v", err)
	}

	// Read response
	rb := make([]byte, 1500)
	err = c.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return TargetInfos{}, false, false, fmt.Errorf("failed to set read deadline: %v", err)
	}

	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Printf("failed to read response: %v", err)
		return TargetInfos{}, false, true, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse the response
	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		log.Printf("failed to parse response: %v", err)
		return TargetInfos{}, false, false, fmt.Errorf("failed to parse response: %v", err)
	}

	log.Println("TTYPE : ", rm.Type, rm.Code, peer)

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got PONG from %v", peer)
		addr := peer.String()
		dnsName, err := net.LookupAddr(addr)
		if err != nil {
			log.Printf("cannot get address : %s\n", err.Error())
			dnsName = append(dnsName, "")
		} else {
			log.Printf("+++ DNS : %s\n", dnsName)
		}

		return TargetInfos{
			ip:   peer.String(),
			name: dnsName[0],
		}, true, false, nil
	case ipv4.ICMPTypeTimeExceeded:
		log.Printf("TTL Exceeded - Intermediate hop found at %v", peer)
		addr := peer.String()
		dnsName, err := net.LookupAddr(addr)

		if err != nil {
			log.Printf("cannot get address : %s\n", err.Error())
			dnsName = append(dnsName, "")
		} else {
			log.Printf("+++ DNS : %s\n", dnsName)
		}
		return TargetInfos{
			ip:   peer.String(),
			name: dnsName[0],
		}, false, false, nil
	default:
		return TargetInfos{}, false, false, fmt.Errorf("wrong case")
	}
}
