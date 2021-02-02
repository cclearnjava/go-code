package examples

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
)

func checkPassword(password string) bool {
	db, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/mydata?charset=uft-8")
	fmt.Println(db, err)

	stmt, err := db.Prepare("select * from meimei")
	fmt.Println(stmt, err)

	if err != nil {
		return false
	} else {
		return true
	}
}

func RunAllData() {
	path := "D:\\work\\project\\NBData\\qqAnd163Password.txt"
	passfile, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	br := bufio.NewReader(passfile) //读取文件对象
	for {
		line, _, end := br.ReadLine() //每次读取一行
		if end == io.EOF {
			break //文件结束，跳出死循环
		}
		linestr := string(line) //读取每一行
		lines := strings.Split(linestr, "#")
		if len(lines) == 2 {
			password := lines[0]
			fmt.Println(password)

			if checkPassword(password) {
				fmt.Println("成功")
				break //跳出循环
			} else {
				fmt.Println("失败")
			}
		}

	}
}

func main() {
	RunAllData()
}
