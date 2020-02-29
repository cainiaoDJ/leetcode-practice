package main

import (
	"fmt"
	"strconv"
	"time"
)

type CalNode struct {
	v  int
	op int
}
type CalList []CalNode

type Stack []int

func (s *Stack) Push(n int) { *s = append(*s, n) }
func (s *Stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func calculate2(s string) int {
	tmp := 0
	stack := make(Stack, 0)
	var lastOp byte = '+'
	s += "+0"
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp = 10*tmp + int(s[i]-'0')
			continue
		} else if s[i] == ' ' {
			continue
		}
		switch lastOp {
		case '+':
			stack.Push(tmp)
			break
		case '-':
			stack.Push(tmp * -1)
			break
		case '*':
			tmp = stack.Pop() * tmp
			stack.Push(tmp)
			break
		case '/':
			tmp = stack.Pop() / tmp
			stack.Push(tmp)
			break

		}
		lastOp, tmp = s[i], 0
	}

	ret := 0
	for len(stack) > 0 {
		ret = ret + stack.Pop()
	}
	return ret
}

func calculate(s string) int {
	tmp := -1
	list := make(CalList, 0)
	for {
		if len(s) == 0 {
			tmp = calHighLevel(&list, tmp)
			tmp = calNormalLevel(&list, tmp)
			//list = append(list, CalNode{tmp, 0})
			//fmt.Println("最终结果:",tmp)
			return tmp
		}
		k := s[:1]
		s = s[1:]

		if k == " " {
			continue
		} else if k == "+" {
			tmp = calHighLevel(&list, tmp)
			tmp = calNormalLevel(&list, tmp)
			list = append(list, CalNode{tmp, 1})
			tmp = -1

		} else if k == "-" {
			tmp = calHighLevel(&list, tmp)
			tmp = calNormalLevel(&list, tmp)
			list = append(list, CalNode{tmp, 2})
			tmp = -1
		} else if k == "*" {
			tmp = calHighLevel(&list, tmp)
			//tmp = calNormalLevel(&list, tmp)
			list = append(list, CalNode{tmp, 3})
			tmp = -1
		} else if k == "/" {
			tmp = calHighLevel(&list, tmp)
			//tmp = calNormalLevel(&list, tmp)
			list = append(list, CalNode{tmp, 4})
			tmp = -1
		} else {
			n, _ := strconv.Atoi(k)
			if tmp >= 0 {
				n = 10*tmp + n
			}
			tmp = n
		}

	}
}

func calHighLevel(list *CalList, num int) int {
	cnt := len(*list)
	if cnt <= 0 {
		return num
	}
	node := (*list)[cnt-1:]
	if node[0].op == 3 {
		*list = (*list)[:cnt-1]
		num = node[0].v * num
	} else if node[0].op == 4 {
		*list = (*list)[:cnt-1]
		num = node[0].v / num
	}
	return num
}

func calNormalLevel(list *CalList, num int) int {
	cnt := len(*list)
	if cnt <= 0 {
		return num
	}
	node := (*list)[:1]
	if node[0].op == 1 {
		*list = (*list)[1:]
		num = node[0].v + num
	} else if node[0].op == 2 {
		*list = (*list)[1:]
		num = node[0].v - num
	}
	return num
}

func main() {
	start := time.Now()
	ret := calculate("100-234+234*234-234*123+34/412-123556*234+234/3")
	cost := time.Since(start)
	fmt.Println("计算结果:", ret, ",耗时:", cost)
	start = time.Now()
	ret = calculate2("100-234+234*234-234*123+34/412-123556*234+234/3")
	cost = time.Since(start)
	fmt.Println("计算结果:", ret, ",耗时:", cost)
	start = time.Now()
	ret = 100 - 234 + 234*234 - 234*123 + 34/412 - 123556*234 + 234/3
	cost = time.Since(start)
	fmt.Println("计算结果:", ret, ",耗时:", cost)
}
