package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"randChar"
	"time"
	"trie"
)

func main() {
	var root *trie.Trie_node
	root = trie.AllocTreeNode()

	var ret bool
	var res string
	var start_time time.Time
	var end_time time.Time
	var dur_time time.Duration

	fo, err := os.OpenFile("dictionary.txt", os.O_CREATE|os.O_RDWR, os.ModeType)
	if err != nil {
		panic(err)
	}

	start_time = time.Now()

	for i := 0; i < 100000; i++ {
		res = randChar.RandChar(i)

		if _, err := fo.WriteString(res + "\n"); err != nil {
			panic(err)
		}
	}

	end_time = time.Now()
	dur_time = end_time.Sub(start_time)
	var elapsed_min string = dur_time.String()

	fmt.Println("创建词库花费", elapsed_min)

	fo.Close() //模拟当文件关闭再重新打开的状态

	//以下的代码是创建字典树
	fi, err := os.OpenFile("dictionary.txt", os.O_RDWR, os.ModeType)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	start_time = time.Now()
	//读取文件
	buf := bufio.NewReader(fi)
	for {
		line, err := buf.ReadString('\n')
		trie.AddWords(line, root)
		if err != nil {
			if err == io.EOF {
				break
			}

		}
	}

	end_time = time.Now()
	dur_time = end_time.Sub(start_time)
	elapsed_min = dur_time.String()
	fmt.Println("创建树花费", elapsed_min)

	//这里是查询的代码
	var search_string string
	fmt.Scanf("%s", &search_string)
	var test int64
	start_time = time.Now()

	//查询
	//1表示全词匹配,0表示前缀匹配
	ret = trie.SearchWords(search_string, root, 1)

	end_time = time.Now()
	dur_time = end_time.Sub(start_time)
	test = dur_time.Nanoseconds()
	fmt.Println("查找时间为", test, ret)

}
