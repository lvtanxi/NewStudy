// map
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//简单初始化
	m := make(map[int]string)
	fmt.Println(m)
	m[1] = "OK"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)
	m[2] = "test"
	m[3] = "jay"
	m[4] = "good"
	m[5] = "hh"

	//双重map
	m1 := make(map[int]map[int]string)
	m1[1] = m
	fmt.Println(m1)

	//多返回值

	value, flag := m1[2][1]
	if !flag {
		m1[2] = make(map[int]string)
		m1[2][1] = "Good"
	} else {
		fmt.Println(value)
	}
	x := m1[2][1]
	fmt.Println(x)

	//map集合的遍历
	for index, value := range m {
		fmt.Println(index, value)
	}

	for index := range m {
		fmt.Println(index)
	}

	for _, value := range m {
		fmt.Println(value)
	}

	sm := make([]map[int]string, 5)
	for index := range sm {
		sm[index] = make(map[int]string)
		sm[index][index] = "ok"
	}
	fmt.Println(sm)

	//间接排序

	smp := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(smp)
	sa := make([]int, len(smp))
	i := 0
	for k := range smp {
		sa[i] = k
		i++
	}
	fmt.Println(sa)
	sort.Ints(sa)
	fmt.Println(sa)

	ssm :=make(map[string]int,len(smp))

	for index,value :=range smp{
		ssm[value]=index
	}
	fmt.Println(ssm)

	fmt.Println(strconv.Itoa(1)+"12")

}
