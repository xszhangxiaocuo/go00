package main

import (
	"bufio"
	"dsLab/lab7/myMap"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
简易搜索引擎
2023年5月20日10:45:44
*/

var docs = make([]string, 0)

func main() {
	//读取文本文件
	//倒排索引构造查找表
	table := buildTable()
	var query string
	fmt.Print("请输入关键词：")
	fmt.Scan(&query)
	//查找query并返回查找结果
	result := find(table, query)
	if len(result) == 0 {
		fmt.Println("结果不存在！")
		return
	}
	fmt.Println(result)
	for i, doc := range docs {
		for _, s := range result {
			if s == fmt.Sprintf("doc%s.txt", strconv.Itoa(i+1)) {
				fmt.Println(s, ":", doc)
			}
		}

	}
	generateHTML(result, docs)
}

func buildTable() *myMap.MyMap {
	table := myMap.CreatMap()
	for i := 1; i <= 10; i++ {
		file := "doc" + strconv.Itoa(i) + ".txt"
		inputFile, err := os.Open("lab7/testdoc/" + file)
		if err != nil {
			fmt.Println("文件打开失败！")
			return nil
		}
		defer inputFile.Close()
		reader := bufio.NewReader(inputFile)
		docs = append(docs, "")
		tmpstr := ""
		for {
			inputString, _, readerError := reader.ReadLine()
			tmpstr += string(inputString)
			words := strings.FieldsFunc(string(inputString), func(r rune) bool {
				return (r < 'A' || r > 'Z') && (r < 'a' || r > 'z')
			})
			//更新查询表
			if len(words) != 0 {
				for _, word := range words {
					tmp := table.Get(word)
					if tmp == nil {
						v := myMap.CreatMap()
						v.Put(file, true) //标记当前文件中已经包含该单词
						table.Put(word, v)
					} else {
						v := tmp.(*myMap.MyMap)
						if v.Get(file) == nil {
							v.Put(file, true)
							table.Put(word, v)
						}
					}
				}
			}
			if readerError == io.EOF {
				break
			}
		}
		docs[i-1] = tmpstr
	}
	return table
}

func find(table *myMap.MyMap, query string) []string {
	result := make([]string, 0)
	v := table.Get(query)
	if v != nil {
		tmpv := v.(*myMap.MyMap)
		keys, _ := tmpv.GetKey()
		for _, key := range keys {
			result = append(result, key)
		}
	}
	return result
}

func generateHTML(result []string, docs []string) {
	// 创建一个HTML模板
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Search Results</title>
	</head>
	<body>
		<h1>Search Results</h1>
		<ul>
		{{range $index, $doc := .}}
			<li><strong>doc{{add $index 1}}.txt</strong>: {{$doc}}</li>
		{{end}}
		</ul>
	</body>
	</html>`

	// 创建一个模板，并添加一个函数
	t, err := template.New("searchResult").Funcs(template.FuncMap{"add": func(i int) int { return i + 1 }}).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// 创建一个HTML文件
	f, err := os.Create("result.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 需要传递给模板的数据
	data := make([]string, 0)
	for i, doc := range docs {
		for _, s := range result {
			if s == fmt.Sprintf("doc%s.txt", strconv.Itoa(i+1)) {
				data = append(data, doc)
			}
		}
	}

	// 将数据应用到模板并写入文件
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("HTML file has been generated successfully!")
}
