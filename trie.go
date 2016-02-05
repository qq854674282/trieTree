package trie

/*
import (
	"fmt"
)*/

type Trie_node struct {
	TreePointer map[byte]*Trie_node //节点
	trie_end    bool                //判断是否为叶子节点
}

func AllocTreeNode() *Trie_node {
	return &Trie_node{
		TreePointer: make(map[byte]*Trie_node),
		trie_end:    false,
	}
}

//添加节点
func AddWords(line string, root *Trie_node) bool {
	if len(line) == 0 {
		return false
	}

	for i := 0; i < len(line); i++ {
		if root.TreePointer[line[i]] == nil {
			root.TreePointer[line[i]] = AllocTreeNode()
		}

		if i == len(line)-1 {
			root.trie_end = true
			//fmt.Println(root)
			break
		}
		//fmt.Println(root)
		root = root.TreePointer[line[i]]
	}
	return true
}

//匹配单词
func SearchWords(line string, root *Trie_node, flag int) bool {
	if len(line) == 0 {
		return false
	}
	//1表示全词匹配,0表示前缀匹配
	for i := 0; i < len(line); i++ {
		//fmt.Println(root)
		if flag == 0 {
			//表示前缀匹配<只要路没有走完,但是已经找到词了,就可以return true>
			//fmt.Println(root.TreePointer[line[i]], root.TreePointer[line[i]].trie_end, i)
			if root.TreePointer[line[i]] != nil && root.TreePointer[line[i]].trie_end == true && i <= len(line)-1 {
				//fmt.Println("111111", root)
				return true
			}
		} else if flag == 1 {
			//表示全词匹配<只要路走完了,并且词也走完了,就可以return true>
			//fmt.Println(root)
			//fmt.Println(root.TreePointer[line[i]], root.TreePointer[line[i]].trie_end, i)
			if root.TreePointer[line[i]] != nil && root.TreePointer[line[i]].trie_end == true && i == len(line)-1 {
				//fmt.Println("2222222", root)
				return true
			}
		} else {
			//扩展用
		}
		if root.TreePointer[line[i]] != nil {
			root = root.TreePointer[line[i]]
		} else {
			return false
		}

		/*		fmt.Println(line[i])
				fmt.Println(i, root.TreePointer[line[i]])
				if flag == 1 {
					if root.trie_end == true && i == len(line)-1 { //全词匹配
						return true
					}
				}
				if flag == 0 {
					if root != nil && root.trie_end == true && i < len(line) { //前缀匹配
						return true
					}
				}
				if root.TreePointer[line[i]] != nil {
					root = root.TreePointer[line[i]]
				} else {
					return false
				}
		*/
	}

	return false //这里表示前面的字串全部匹配,但是长度不够,所以,为false

}

func DelWords(line string, root *Trie_node) {
	if len(line) == 0 {
		return
	}
	var res bool
	res = SearchWords(line, root, 1) //使用全词匹配
	if res == false {
		return
	}

	DelWord(line, root, 0)

	/*var temp *Trie_node
	temp = AllocTreeNode()*/

	/*for i := 0; i < len(line); i++ {
		if root.TreePointer[line[i]] != nil {
			temp = root.TreePointer[line[i]]
			root = temp
			//delete(temp)
		}
	}*/

}

func DelWord(line string, root *Trie_node, i int) {
	if i == len(line)-1 {
		return
	}
	DelWord(line, root, i+1)
	root.TreePointer[line[i]] = nil
}
