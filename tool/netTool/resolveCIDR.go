/*
@Time : 2021/9/2 15:06
@Author : TMNT1983
*/

package netTool

import (
	"fmt"
	"strconv"
	"strings"
)

type NetIP struct{}

//获取CIDR的范围值
func (n NetIP) GetCidrIpRange(cidr string) (string, string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)
	ipPrefix := ipSegs[0] + "." + ipSegs[1] + "."

	return ipPrefix + strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp),
		ipPrefix + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)
}

//获取CIDR地址范围内可拥有的主机数量
func (n NetIP) GetCidrHostNum(maskLen int) uint {
	cidrIpNum := uint(0)
	var i uint = uint(32 - maskLen - 1)
	for ; i >= 1; i-- {
		cidrIpNum += 1 << i
	}
	return cidrIpNum
}

//获取Cidr的掩码
func (n NetIP) GetCidrIpMask(maskLen int) string {
	// ^uint32(0)二进制为32个比特1，通过向左位移，得到CIDR掩码的二进制
	cidrMask := ^uint32(0) << uint(32-maskLen)
	fmt.Println(fmt.Sprintf("%b \n", cidrMask))
	//计算CIDR掩码的四个片段，将想要得到的片段移动到内存最低8位后，将其强转为8位整型，从而得到
	cidrMaskSeg1 := uint8(cidrMask >> 24)
	cidrMaskSeg2 := uint8(cidrMask >> 16)
	cidrMaskSeg3 := uint8(cidrMask >> 8)
	cidrMaskSeg4 := uint8(cidrMask & uint32(255))

	return fmt.Sprint(cidrMaskSeg1) + "." + fmt.Sprint(cidrMaskSeg2) + "." + fmt.Sprint(cidrMaskSeg3) + "." + fmt.Sprint(cidrMaskSeg4)
}

//判读ip是否在这个CIDR范围内
func (n NetIP) MatchIp(ipAdress, cidr string) bool {
	minIP, maxIP := n.GetCidrIpRange(cidr)
	ipNum := ipToInt(ipAdress)
	minIpNum := ipToInt(minIP)
	maxIpNum := ipToInt(maxIP)

	if minIpNum <= ipNum && ipNum <= maxIpNum {
		return true
	}
	return false
}

//获取第三段IP的区间
func getIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return getIpSegRange(uint8(ipSeg), uint8(24-maskLen))
}

//获取第四段IP的区间
func getIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	return segMinIp + 1, segMaxIp
}

//根据用户输入的基础IP地址和CIDR掩码计算一个IP片段的区间
func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}

//将ip地址转变为int64
func ipToInt(ipAdress string) int64 {
	ipList := strings.Split(ipAdress, ".")

	bin := make([]string, len(ipList))
	for i, v := range ipList {
		bin[i] = fmt.Sprintf("%03v", v)
	}

	num, err := strconv.ParseInt(strings.Join(bin, ""), 10, 64)
	if err != nil {
		fmt.Printf("ip转换为Int64失败: %v\n", err)
		return 0
	}
	return num
}
