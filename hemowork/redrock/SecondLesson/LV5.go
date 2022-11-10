package main

import "fmt"

type coder struct {
	name      string
	age       int
	number    int64
	code      int
	telephone int64
	birthday  string
}

var (
	i            int
	choose       string
	searchName   string
	searchNumber int64
	revise       string
	rename       string
	reage        int
	renumber     int64
	recode       int
	retelephone  int64
	rebirthday   string
	count        int
	operate      string
)

func main() {

	c := new([100]coder)
	fmt.Println("输入同学的数据")
	fmt.Println("按照名字，年龄，学号，代码行，电话，生日（数据之间用空格）输入")
	for i = 0; i < 100; i++ {
		fmt.Scan(&c[i].name, &c[i].age, &c[i].number, &c[i].code, &c[i].telephone, &c[i].birthday)
		fmt.Println(c[i])
		fmt.Println("输入yes或no表示你是否继续录入同学的数据")
		fmt.Scan(&choose)
		if choose == "no" {
			break
		}
	}
	fmt.Println("请输入你要进行的操作：排序or查找or修改数据")
	fmt.Scan(&operate)
	switch operate {
	case "排序":
		sort1(c)
	case "查找":
		search(c)
	case "修改数据":
		remake(c)
	}
}

// 查询
func search(c *[100]coder) {
	fmt.Println("请输入同学的名字和学号进行查找")
	fmt.Scan(&searchName, &searchNumber)
	for i = 0; i < 100; i++ {
		if c[i].name == searchName && c[i].number == searchNumber {
			fmt.Println(c[i])
		}
	}
}

// 信息修改
func remake(c *[100]coder) {
	fmt.Println("请输入同学的名字和学号进行查找并修改")
	fmt.Scan(&searchName, &searchNumber)
	for i = 0; i < 100; i++ {
		if c[i].name == searchName && c[i].number == searchNumber {
			fmt.Println(c[i])
			break
		}
	}
	fmt.Println("请输入你要修改什么？例如：姓名")
	fmt.Scan(&revise)
	switch revise {
	case "姓名":
		fmt.Println("请输入新的姓名")
		fmt.Scan(&rename)
		c[i].name = rename
	case "年龄":
		fmt.Println("请输入新的年龄")
		fmt.Scan(&reage)
		c[i].age = reage
	case "学号":
		fmt.Println("请输入新的学号")
		fmt.Scan(&renumber)
		c[i].number = renumber
	case "代码行":
		fmt.Println("请输入新的代码行")
		fmt.Scan(&recode)
		c[i].code = recode
	case "电话":
		fmt.Println("请输入新的电话")
		fmt.Scan(&retelephone)
		c[i].telephone = retelephone
	case "生日":
		fmt.Println("请输入新的生日")
		fmt.Scan(&rebirthday)
		c[i].birthday = rebirthday
	}
	fmt.Println("修改成功：", c[i])
}

// 代码行大于150的同学输出完整信息 并进行排序
func sort1(c *[100]coder) {
	fmt.Println("代码行超过150的同学(已排序)")
	//fmt.Println(c)
	for i = 0; i < 100; i++ {
		if c[i].code > 150 {
			count++
			//fmt.Println(c[i])
			for a := 1; a <= count; a++ {
				for j := 0; j < count-i; j++ {
					if c[j].code >= c[j+1].code {
						c[j], c[j+1] = c[j+1], c[j]
					}
				}
			}
			fmt.Println(c[i])
		} else if c[i].code <= 150 {
			fmt.Println("没有写代码超过150行的同学")
			break
		}
	}
}
