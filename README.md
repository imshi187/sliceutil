# sliceutil

golang sliceutil, including methods like TakeWhile, Filter, ForEach and so on, which were inpired by js

# examples

```go
package main

import (
	"fmt"
	"github.com/imshi187/sliceutil"
)

func main() {
	var s sliceutil.SliceUtil
	s.Elements = []interface{}{}
	// append an element
	s.Append(1, 100, 200, 3000) //1 100 200 3000

	// 遍历每一个元素
	s.Foreach(func(index int, item interface{}) {
		fmt.Print(item, " ")
	})

	fmt.Println()
	s.MapTo(func(item interface{}) interface{} {
		return item.(int) - 1
	}).Foreach(func(index int, item interface{}) {
		fmt.Print(item, " ")
	}) //0 99 199 2999

	fmt.Println()
	fmt.Println(s.Where(99)) //1

	s.Modify(1, 666).Foreach(func(index int, item interface{}) {
		fmt.Print(item, " ")
	}) //0 666 199 2999

	allGreater := s.AllMatch(func(item interface{}) bool {
		// if all elements are greater than 1
		return item.(int) >= 1
	})
	fmt.Println(allGreater) //false

	s.Filter(func(item interface{}) bool {
		return item.(int) <= 0
	}).Foreach(func(index int, item interface{}) {
		fmt.Println(item)
	}) //only 0

}

```

# install

import github.com/imshi187/sliceutil
