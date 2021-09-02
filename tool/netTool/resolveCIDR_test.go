package netTool

import (
	"fmt"
	"testing"
)

func TestResolveCidr(t *testing.T) {
	var a netIP
	minIP,maxIP:=a.GetCidrIpRange("20.26.17.0/24")
	fmt.Printf("最小IP地址：%v , 最大IP地址：%v \n",minIP,maxIP)

	hostNum:=a.GetCidrHostNum(24)
	fmt.Printf("CIDR 最多有几个可以分配：%v\n",hostNum)

	maskIP:=a.GetCidrIpMask(24)
	fmt.Printf("CIDR 的掩码地址：%v\n",maskIP)
}