package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var r *bufio.Reader

func init() {
	r = bufio.NewReader(os.Stdin)
}

func ReadConsoleString() string {
	data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1)
	return data
}

func ReadConsoleInt64() int64 {
	data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	dataInt, _ := strconv.ParseInt(data, 10, 64)
	return dataInt
}

func ReadConsoleFloat64() float64 {
	data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	dataFloat, _ := strconv.ParseFloat(data, 64)
	return dataFloat
}
