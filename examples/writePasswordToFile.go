package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func savePasswordFile() {
	path := "D:\\work\\project\\NBData\\qqAnd163Password.txt"
	passfile, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	br := bufio.NewReader(passfile)
	savefile, err := os.Create("D:\\work\\project\\NBData\\password.txt")
	defer savefile.Close()

	save := bufio.NewWriter(savefile)

	for {
		line, _, end := br.ReadLine()
		linestr := string(line)
		if end == io.EOF {
			break //跳出循环
		}
		lines := strings.Split(linestr, "#")
		if len(lines) == 2 {
			password := lines[1]
			fmt.Fprintln(save, password) //写入
		}
	}
	save.Flush() //刷新
}
