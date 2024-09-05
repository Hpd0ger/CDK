package dns

import (
	"github.com/cdk-team/CDK/pkg/util"
	"log"
	"net"
)

// DnsReverseLookup is a function to get PTR records for a given IP address
func dnsPTRLookup(ipAddr string) (ptrs []string, err error) {
	// Perform the reverse DNS lookup using the IP address
	ptrs, err = net.LookupAddr(ipAddr)
	return
}

// DnsDiscovery is a function to get Cluster IP address and the domain name for the IP address
func DnsDiscovery(ipCidr string) {
	// Perform the reverse DNS lookup using the IP address
	ips, err := util.ParseCidrToIPs(ipCidr)
	if err != nil {
		log.Fatal("Error parsing CIDR to IPNet:", err)
	}

	for _, ip := range ips {
		ptrRecords, _ := dnsPTRLookup(ip.String())
		// Print the domain name for the IP address
		if len(ptrRecords) > 0 {
			log.Printf("Domain name %v found for IP address %s:\n", ptrRecords, ip.String())
		}
	}
}

func RunDnsReverseLookup(args []string) {
	if len(args) != 2 {
		log.Fatal("Invalid input args, Example: ./cdk dns discovery 10.16.0.0/16")
	}

	switch args[0] {
	case "discovery":
		cidrIp := args[1]
		DnsDiscovery(cidrIp)
	}
}
