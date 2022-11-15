package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"sort"
	"strconv"
)

type VisitEvent struct {
	RequestID string `json:"request_id" bson:"request_id,omitempty"`
	Timestamp int64  `json:"timestamp" bson:"timestamp,omitempty"`
	Client    string `json:"client"`
}

type RiskWrapper struct {
	people []VisitEvent
	by     func(p, q *VisitEvent) bool
}

type SortBy func(p, q *VisitEvent) bool

func (pw RiskWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw RiskWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw RiskWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}

func SortRisk(people []VisitEvent, by SortBy) {
	sort.Sort(RiskWrapper{people, by})
}

//
//func main() {
//
//	//var ch chan string
//	//
//	//<-ch
//	//
//	//lock := new(sync.Mutex)
//	//lock.Lock()
//
//	//研究值传递和地址传递
//	//1. 值传递只会把参数的值复制一份放进对应的函数，两个变量的地址不同，
//	//不可相互修改。
//	//2. 地址传递(引用传递)会将变量本身传入对应的函数，在函数中可以对该变
//	//量进行值内容的修改。
//	//18、Go 语言当中数组和切片在传递的时候的区别是什么？
//	//1. 数组是值传递
//	//2. 切片是引用传递
//
//	//  创建一个长度为2的int数组
//	//arr := make([]int, 0)
//	//fmt.Println(reflect.TypeOf(arr))
//	var s []int32
//	//fmt.Println(append(s, 1))
//	if s == nil {
//		fmt.Println(1)
//	} else {
//		fmt.Println(2)
//	}
//	s0 := append(s, 1)
//	fmt.Println(len(s0))
//	fmt.Println(cap(s))
//	fmt.Println(append(s, 1))
//
//	ar2 := make([]int, 0)
//	ar2 = append(ar2, 1, 2)
//	fmt.Println(ar2)
//
//	ar3 := [2]int{}
//	ar3[0] = 1
//	ar3[1] = 2
//	fmt.Println(ar3)
//	//sort.Reverse(ar3)
//
//	events := make([]VisitEvent, 0)
//	events = append(events, VisitEvent{
//		RequestID: "1",
//		Timestamp: 1,
//		Client:    "1",
//	}, VisitEvent{
//		RequestID: "5",
//		Timestamp: 5,
//		Client:    "5",
//	}, VisitEvent{
//		RequestID: "2",
//		Timestamp: 2,
//		Client:    "2",
//	}, VisitEvent{
//		RequestID: "4",
//		Timestamp: 4,
//		Client:    "4",
//	}, VisitEvent{
//		RequestID: "3",
//		Timestamp: 3,
//		Client:    "3",
//	})
//	SortRisk(events, func(p, q *VisitEvent) bool {
//		return p.Timestamp > q.Timestamp
//	})
//	fmt.Println(events)
//}

func main() {
	//	arr := make([]int, 0, 10)
	//	for i := 0; i < 2000; i++ {
	//		if i == 512 {
	//			fmt.Println(i)
	//		}
	//		fmt.Println("len 为", len(arr), "cap 为", cap(arr))
	//		runtime.GC()
	//		arr = append(arr, i)
	//	}
	//
	//	runtime.GC()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.10.20.187:6379",
		Password: "sBL2y7Uuxqyi", // no password set
		DB:       0,              // use default DB
	})
	name, _ := rdb.HGet(context.Background(), "person", "name").Result()
	fmt.Println(name)

	fmt.Println(math.MaxInt64)

	a := make([]interface{}, 0)
	for i := 0; i <= 512; i++ {
		a = append(a, strconv.Itoa(i), strconv.Itoa(math.MaxInt64)+strconv.Itoa(math.MaxInt64)+
			strconv.Itoa(math.MaxInt64)+"1111111")
	}

	b := rdb.HSet(context.Background(), "person", a...).String()
	fmt.Println(b)
}
