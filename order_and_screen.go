package main


import (
	"bufio"
	"github.com/360EntSecGroup-Skylar/excelize"

	// "bytes"
	"encoding/csv"
	"flag"
	"fmt"
	//"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	// "math"
)

type Param struct {
	c string
	si int
	sm string
	st string
	so string
	oi int
	ot string
}
var params Param
func main() {
	var xlsxfileName    string
	var csvfileName    string
	var txtfileName    string
	var searchIndex int
	var searchName string
	var searchType string
	var searchOperator string
	var orderIndex int
	var orderType string
	var column string
	// var key int
	// var f int
	flag.StringVar(&xlsxfileName, "xlsx", "", "xlsx文件地址不能为空")
	flag.StringVar(&csvfileName, "csv", "", "csv文件地址不能为空")
	flag.StringVar(&txtfileName, "txt", "", "txt文件地址不能为空")
	flag.IntVar(&searchIndex, "si", -1, "筛选列")
	flag.StringVar(&searchName, "sm", "", "筛选关键字")
	flag.StringVar(&searchType, "st", "", "筛选关类型") //1:int 2:sting
	flag.StringVar(&searchOperator, "so", "", "筛选条件") //1-1:等于，1-2:不等于, 1-3:大于，1-4:大于等于, 1-5:小于,1-6:小于等于,1-7:介于//2-1:等于，2-2:不等于, 2-3:包含，2-4:不包含
	flag.IntVar(&orderIndex, "oi", -1,"排序列")
	flag.StringVar(&orderType, "ot", "", "排序类型")
	flag.StringVar(&column, "c", "", "列索引")
	flag.Parse()
	
	params = Param{column, searchIndex, searchName, searchType, searchOperator, orderIndex, orderType}

	if xlsxfileName != "" && params.c != ""  {
		filterXlsx(xlsxfileName)
	} else if csvfileName != "" && params.c != ""  {
		filterCsv(csvfileName)
	} else if txtfileName != "" && params.c != ""  {
		filterTxt(txtfileName)
	} else {
		fmt.Printf("%s\n", "error")
	}
}

func filterXlsx(filename string) {
    _, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, _ := excelize.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	sheet_name := f.GetSheetName(1)
	rows := f.GetRows(sheet_name)
	order(rows)

}

func filterCsv(filename string) {
	fs, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer fs.Close()

	r := csv.NewReader(fs)
	// i := 0
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
		// fmt.Println(row)
		newContent = append(newContent, row)
	}
	order(newContent)
}

func filterTxt(filename string) {
	f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    buf := bufio.NewReader(f)
	var result  []string
	var newContent [][]string
	// i := 0
    for {
        a, _, c := buf.ReadLine()
        if c == io.EOF {
            break
		}
		result = strings.Split(string(a), ",")
		newContent = append(newContent, result)
	}
	order(newContent)
	
}

func order(rows[][]string) {
	col := strings.Split(string(params.c), ",")
	// lie := [...]int{0, 1, 2, 3}a
	search_index := params.si //筛选的列名下标
	serch_name := params.sm //筛选关键字
	serch_type := params.st //筛选类型
	serch_op := params.so //筛选关条件
	order_index := params.oi //排序的列名下标
	if params.oi == -1 {
		order_index, _ = strconv.Atoi(col[0]) //排序的列名下标
	}

	order_type := params.ot
	str := make(map[int]string)
	data := make(map[int]string)
    for k, row := range rows {
		flag := 0
		value := ""
		index := true
		for _, val := range col {
			col_num, _ := strconv.Atoi(val)
			if(serch_name != "") {
				if(serch_type == "1") {
					index_data, _ := strconv.Atoi(row[search_index])
					serchname, _ := strconv.Atoi(serch_name)
					if serch_op == "1" {
						index = (index_data == serchname)
					} else if serch_op == "2" {
						index = (index_data != serchname)
					} else if serch_op == "3" {
						index = (index_data > serchname)
					} else if serch_op == "4" {
						index = (index_data >= serchname)
					} else if serch_op == "5" {
						index = (index_data < serchname)
					} else if serch_op == "6" {
						index = (index_data <= serchname)
					} else if serch_op == "7" {
						serchnum := strings.Split(string(serch_name), ",")
						serchname_left, _ := strconv.Atoi(serchnum[0])
						serchname_right, _ := strconv.Atoi(serchnum[1])
						index_left := (serchname_left <= index_data)
						index_right := (index_data <= serchname_right)
						index = (index_left == index_right)
					}
				} else {
					if serch_op == "1" {
						index = (row[search_index] == serch_name)
					} else if serch_op == "2" {
						index = (row[search_index] != serch_name)
					} else if serch_op == "3" {
						index = strings.Contains(row[search_index], serch_name)
					} else if serch_op == "4" {
						is_includ := strings.Contains(row[search_index], serch_name)
						index = !is_includ
					}
				}
			}
			if(index != true && k > 0) {
				continue
			}
			flag = 1
			if(order_index == col_num) {
				str[k] = row[col_num]
			}
			value = value + row[col_num] + "\t"
		}
		if(flag == 1) {
			data[k] = value
		}
	}

	type List struct {
		name string
		num int
	}
	var aa []List
	for k, v := range str {
		aa = append(aa, List{v, k})	
	}
	// if serch_type == "1" {
	// 	for k, v := range str {
	// 		v1, _ := strconv.Atoi(v)
	// 		aa = append(aa, List{v1, k})	
	// 	}
	// } else {
	// 	for k, v := range str {
	// 		aa = append(aa, List{v, k})	
	// 	}

	// }
	if(order_type != "") {
		sort.Slice(aa, func(i, j int) bool {
			if(order_type == "asc") {
				return aa[i].name < aa[j].name  // 升序
			} else {
				return aa[i].name > aa[j].name  // 降序
			}
		})
	}
	res := ""
	max := 100
	for k1, v1 := range aa {
		if(k1 >= max) {
			break
		}
		if(v1.num == 0) {
			continue
		} else {
			res = res + data[v1.num] + "\n"
		}
	}
	res = data[0] + "\n" + res 
	fmt.Println(res)
}