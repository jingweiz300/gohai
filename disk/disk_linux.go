package disk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func getDiskInfo() (interface{}, error) {
	//从/proc/diskstats结果获取结果
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return nil, err
	}
	// 处理结束后关闭文件
	defer file.Close()
	// 使用bufio读取
	r := bufio.NewReader(file)
	for {
		// 分行读取文件  ReadLine返回单个行，不包括行尾字节(\n  或 \r\n)
		data, _, err := r.ReadLine()

		// // 以分隔符形式读取,比如此处设置的分割符是\n,则遇到\n就返回,且包括\n本身 直接返回字符串
		// data, err := r.ReadString('\n')

		// // 打印出内容
		// fmt.Printf("%v", string(data))
		// fmt.Println("-------------------------")

		// // 以分隔符形式读取,比如此处设置的分割符是\n,则遇到\n就返回,且包括\n本身 直接返回字节数数组
		// // data, err := r.ReadBytes('\n')

		// // 读取到末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}

		// 打印出内容
		line := string(data)
		//fmt.Printf("%v\n", line)
		r := regexp.MustCompile("[^\\s]+")
		res := r.FindAllString(line, -1)
		fmt.Println(res)
		if res[1] != "0" {
			continue
		}
		if strings.HasPrefix(res[2], "fd") || strings.HasPrefix(res[2], "loop") || strings.HasPrefix(res[2], "sr") {
			continue
		}
		fmt.Printf("match:%v\n", res)
	}
	return nil, nil
}
