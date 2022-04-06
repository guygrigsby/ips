package ips

import (
	"math/rand"
	"strings"
	"testing"
)

var (
	ipList []string
)

func init() {
	ipList = makeIPs(2000000)
}

//func BenchmarkManyRequestHandled(b *testing.B) {
//	for _, ip := range ipList {
//		request_handled(ip)
//	}
//}
//func BenchmarkOneRequestHandled(b *testing.B) {
//	request_handled(ipList[0])
//}
func BenchmarkTop100(b *testing.B) {
	for _, ip := range ipList {
		request_handled(ip)
	}
	top100()
}
func makeIPs(count int) []string {
	var all []string
	for i := 0; i < count; i++ {
		all = append(all, randIP())
	}
	return all
}
func randIP() string {
	ip := make([]byte, 4)
	rand.Read(ip)
	var s strings.Builder

	for i, entry := range ip {
		if i != 0 {
			s.Write([]byte("."))
		}
		s.WriteByte(entry)
	}
	return s.String()
}
