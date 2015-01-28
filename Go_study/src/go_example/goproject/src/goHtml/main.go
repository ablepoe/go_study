// goHtml project main.go
package main

import (
	"fmt"
	"goHtml/handlers"
)

const (
	ConfFile = "server.xml"
)

type DasSource struct {
	ConfFile string //配置文件
}

func getDasSource(ConfFile string) (d *DasSource) {
	d = new(DasSource)
	d.ConfFile = ConfFile
	return
}

func main() {
	d := getDasSource(ConfFile)
	fmt.Println(*d)
	handler.Test()
}
