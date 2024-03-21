package util

import (
	"encoding/binary"
	"net"
)

// ParseCidrToIPs 解析CIDR类型的ip地址为ip列表
func ParseCidrToIPs(cidr string) (ips []net.IP, err error) {
	var ipv4Net *net.IPNet
	_, ipv4Net, err = net.ParseCIDR(cidr)

	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips = append(ips, ip)
	}
	return
}
