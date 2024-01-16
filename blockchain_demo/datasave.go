package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type recode struct {
	number   int
	PrevHash string
	Data     string
	Hash     string
	nouce    string
}

var AllCode []recode

func SetInfo() {
	fmt.Println("请输入本次记录（记录员/记录时间/记录信息）：")
	var message string
	fmt.Scanln(&message)
	recodes := strings.Split(message, "/")

	if len(recodes) == 4 {
		ConfirmInfo(recodes)
	}
}

func ConfirmInfo(args []string) {
	var inputRecode recode
	if len(args) == 4 {
		inputRecode.Data, inputRecode.PrevHash, inputRecode.Hash, inputRecode.nouce = args[0], args[1], args[2], args[3]
		inputRecode.number = GetNum() + 1
	} else if len(args) == 5 {
		number, _ := strconv.Atoi(args[0])
		inputRecode.number, inputRecode.Data, inputRecode.PrevHash, inputRecode.Hash, inputRecode.nouce = number, args[1], args[2], args[3], args[4]
	} else {
		return
	}

	AllCode = append(AllCode, inputRecode)
}

func GetNum() int {
	for k, v := range AllCode {
		if v.number == 0 {
			return k
		}

		if k == len(AllCode)-1 && v.number != 0 {
			return len(AllCode)
		}
	}

	return 0
}

func FindData() {
	var num int = GetNum()
	if num == 0 {
		return
	}

	fmt.Println("请输入查询关键词：")
	var find string
	fmt.Scan(&find)
	for _, v := range AllCode {
		if v.Data == find {
			fmt.Println(v.number, ":", v.Data, "/", v.PrevHash, "/", v.Hash, "/", v.nouce)
		}
	}
}

func SaveData() {
	filePath := "D:/GO_work/data.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	write := bufio.NewWriter(file)
	for _, v := range AllCode {
		if v.number != 0 {
			write.WriteString(strconv.Itoa(v.number) + "/" + v.Data + "/" + v.PrevHash + "/" + v.Hash + "/" + v.nouce)
		}
		write.WriteString("\n")
	}
	write.Flush()
	file.Close()
}

func GetData() {
	filePath := "D:/GO_work/data.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	read := bufio.NewReader(file)
	for {
		lineBytes, err := read.ReadBytes('\n')
		line := string(lineBytes)
		ConfirmInfo(strings.Split(line, "/"))
		if err == io.EOF {
			break
		}
	}
}
