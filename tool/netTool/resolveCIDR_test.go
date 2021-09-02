package netTool

import (
	"fmt"
	"testing"
)

func TestResolveCidr(t *testing.T) {
	var a NetIP
	minIP,maxIP:=a.GetCidrIpRange("20.26.17.0/24")
	fmt.Printf("最小IP地址：%v , 最大IP地址：%v \n",minIP,maxIP)

	hostNum:=a.GetCidrHostNum(24)
	fmt.Printf("CIDR 最多有几个可以分配：%v\n",hostNum)

	maskIP:=a.GetCidrIpMask(24)
	fmt.Printf("CIDR 的掩码地址：%v\n",maskIP)

	res:=a.MatchIp("20.26.17.250","20.26.17.0/24")
	if res{
		fmt.Println("输入的ip在CIDR里面！")
	}else {
		fmt.Println("输入的ip不在CIDR里面！！")
	}
}