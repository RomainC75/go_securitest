package scenarios

import (
	"fmt"
	"net"
	"os"
	shared_dto "shared/dto"
	"syscall"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type SFingerPrint struct {
	payload shared_dto.IpRange
}

func NewFingerPrint(payload shared_dto.IpRange) *SFingerPrint {
	return &SFingerPrint{
		payload: payload,
	}
}

func (fp *SFingerPrint) Run() (int, error) {

	destination := fp.payload.IpMin

	conn, err := net.Dial("ip4:icmp", destination)
	if err != nil {
		return 0, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Get the underlying file descriptor
	rawConn, err := conn.(*net.IPConn).SyscallConn()
	if err != nil {
		return 0, fmt.Errorf("failed to get raw connection: %v", err)
	}

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
		return 0, fmt.Errorf("failed to marshal message: %v", err)
	}

	var controlErr error
	err = rawConn.Control(func(fd uintptr) {
		controlErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TTL, 64)
	})
	if err != nil || controlErr != nil {
		return 0, fmt.Errorf("failed to set TTL: %v, %v", err, controlErr)
	}

	_, err = conn.Write(wb)
	if err != nil {
		return 0, fmt.Errorf("failed to send packet: %v", err)
	}

	rb := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	n, err := conn.Read(rb)
	if err != nil {
		return 0, fmt.Errorf("failed to read response: %v", err)
	}

	var ttl int
	rawConn.Control(func(fd uintptr) {
		ttl, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TTL)
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get TTL: %v", err)
	}

	// Parse the ICMP response
	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		return ttl, fmt.Errorf("failed to parse response: %v", err)
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("Received Echo Reply from %s\n", destination)
		return ttl, nil
	default:
		return ttl, fmt.Errorf("unexpected ICMP message type: %v", rm.Type)
	}
}

func (fp *SFingerPrint) Check() error {
	return nil
}
