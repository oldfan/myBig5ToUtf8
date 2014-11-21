// php mysql big-5 to utf-8  project main.go
package main

import (
	"bufio"
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io"
	"os"
	"strings"
)

func main() {
	str := [89]string{"TYPE=MyISAM", "么\\", "功\\", " 吒\\", "吭\\", "沔\\", "坼\\", "歿\\", "俞\\", "枯\\", "苒\\", "娉\\", "珮\\", "豹\\", "崤\\", "淚\\", "許\\", "廄\\", "琵\\", "跚\\", "愧\\", "稞\\", "鈾\\", "暝\\", "蓋\\", "墦\\", "穀\\", "閱\\", "璞\\", "餐\\", "縷\\", "擺\\", "黠\\", "孀\\", "踊\\", "髏\\", "躡\\", "尐\\", "佢\\", "汻\\", "岤\\", "垥\\", "柦\\", "胐\\", "娖\\", "涂\\", "罡\\", "偅\\", "惝\\", "牾\\", "莍\\", "傜\\", "揊\\", "焮\\", "茻\\", "鄃\\", "幋\\", "滜\\", "綅\\", "赨\\", "塿\\", "縷\\", "槙\\", "擺\\", "箤\\", "踊\\", "嫹\\", "髏\\", "潿\\", "蔌\\", "醆\\", "嬞\\", "獦\\", "佢\\", "螏\\", "餤\\", "燡\\", "螰\\", "駹\\", "礒\\", "鎪\\", "瀙\\", "酀\\", "瀵\\", "騱\\", "酅\\", "贕\\", "鱋\\", " 鱭\\"}
	str1 := [89]string{"ENGINE=MyISAM", "么", "功", " 吒", "吭", "沔", "坼", "歿", "俞", "枯", "苒", "娉", "珮", "豹", "崤", "淚", "許", "廄", "琵", "跚", "愧", "稞", "鈾", "暝", "蓋", "墦", "穀", "閱", "璞", "餐", "縷", "擺", "黠", "孀", "踊", "髏", "躡", "尐", "佢", "汻", "岤", "垥", "柦", "胐", "娖", "涂", "罡", "偅", "惝", "牾", "莍", "傜", "揊", "焮", "茻", "鄃", "幋", "滜", "綅", "赨", "塿", "縷", "槙", "擺", "箤", "踊", "嫹", "髏", "潿", "蔌", "醆", "嬞", "獦", "佢", "螏", "餤", "燡", "螰", "駹", "礒", "鎪", "瀙", "酀", "瀵", "騱", "酅", "贕", "鱋", " 鱭"}
	sqlFile := os.Args[1]

	//開啟檔案
	sqlopen, err := os.Open(sqlFile)
	if err != nil {
		fmt.Println(sqlopen, err)
		return
	}

	inputReader := bufio.NewReader(sqlopen)
	defer sqlopen.Close()
	outfile, err := os.Create("outfile.sql")
	if err != nil {
		fmt.Println("檔案建立錯誤")
		return
	}
	defer outfile.Close()

	for {
		inputString, err := inputReader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("轉換完成")
			//fmt.Println(newstr)
			return
		}
		strout, err := iconv.ConvertString(inputString, "big-5", "utf-8")
		if err != nil {
			fmt.Println("轉碼錯誤%s", err)
		}
		for i := 0; i < 89; i++ {
			if strings.Contains(strout, str[i]) {
				strout = strings.Replace(strout, str[i], str1[i], -1)

			}
		}
		outfile.WriteString(strout)

	}

}
