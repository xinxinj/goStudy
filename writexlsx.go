package main


import (
	"bytes"
	"encoding/csv"
	"bufio"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"io"
	"strings"
)

func main() {
	var xlsxfileName    string
	//var xlsfileName    string
	var csvfileName    string
	var txtfileName    string
	var value string
	var key int
	flag.StringVar(&xlsxfileName, "xlsx", "", "xlsx文件地址不能为空")
	flag.StringVar(&csvfileName, "csv", "", "csv文件地址不能为空")
	flag.StringVar(&txtfileName, "txt", "", "txt文件地址不能为空")
	flag.StringVar(&value, "v", "", "value不能为空")
	flag.IntVar(&key, "k", -1, "key必须为整数")
	flag.Parse()
	if xlsxfileName != "" && value != "" && key >= 0 {
		writeXlsx(xlsxfileName, key, value)
	} else if csvfileName != "" && value != "" && key >= 0 {
		writeCsv(csvfileName, key, value)
	} else if txtfileName != "" && value != "" && key >= 0 {
		writeTxt(txtfileName, key, value)
	} else {
		fmt.Printf("%s\n", "error")
	}

}

//  编辑xlsx单元格内容
func writeXlsx(filename string, key int, value string)  {
	_, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, _ := excelize.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	arr := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	inde := ""
	if(key > 26) {
		v := key/26
		m := key%26
		inde = arr[v-1] + arr[m-1]
	} else {
		inde = arr[key]
	}
	
	index := inde + "1"
	sheet_name := f.GetSheetName(1)
	f.SetCellValue(sheet_name, index, value)
	if err := f.SaveAs(filename); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", "success")
	}
}

//编辑csv
func writeCsv(filename string, key int, value string)  {
	fs, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer fs.Close()

	r := csv.NewReader(fs)
	i := 0
	//针对大文件，一行一行的读取文件
	var newContent [][]string
    for {
        row, err := r.Read()
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if err == io.EOF {
            break
		}
		if i == 0 {
			row[key] = value
		}
		i++
		newContent = append(newContent, row)
		doCsv(filename, newContent)
	}
	fmt.Printf("%s", "success")
}

func doCsv(filename string, row [][]string) {
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	for i:=0;i<len(row);i++ {
		r2.Write(row[i])
	}
	
	r2.Flush()
	fout,err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
	defer fout.Close()
	if err != nil {
		fmt.Println(filename,err)
		return
	}
	_, err = fout.WriteString(buf.String())
	if err != nil {
		fmt.Println(filename,err)
		return
	}
}

// 编辑txt
func writeTxt(filename string, key int, value string) {
	f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    buf := bufio.NewReader(f)
	var result  []string
	line := ""
	i := 0
    for {
        a, _, c := buf.ReadLine()
        if c == io.EOF {
            break
		}
		if i == 0 {
			result = strings.Split(string(a), ",")
			result[key] = value
			line += strings.Join(result, ",") + "\n"
		} else {
			line += string(a) + "\n"
		}
		i++
	}
	doTxt(filename, strings.Trim(line, "\n"))
	//fmt.Println(strings.Trim(line, "\n"))
}

func doTxt(filename string, line string) {
	f, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = f.WriteString(line)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Printf("%s", "success")
}
