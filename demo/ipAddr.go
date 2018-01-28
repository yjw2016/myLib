//自动执行String()方法
package main

import "fmt"

type IPAddr [4]byte

//当执行所有print时，String()将自动执行
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])

}

func main() {
	//var ips IPAddr = [4]byte{127,0,0,1}
	map0 := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for k, v := range map0 {
		fmt.Println(k, ":", v)

	}

	//resl := ips.format2str()
	//fmt.Println(resl)
}
