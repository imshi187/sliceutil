# sliceutil
golang sliceutil, including methods like TakeWhile, Filter, ForEach and so on, which were inpired from js



# code
	var intS SliceUtil
	// SliceUtil结构体的field仅仅包括[]interface{}类型的elements
	intS.elements = []interface{}{1, 2, 3, 100, 200}

	fmt.Println("------------------------------------")
	fmt.Println(intS.TakeWhile(func(ele interface{}) bool {
		// 一般都是同一类型的数据，这里是int类型
		return ele.(int) > 2
	}))

	// filter会对slice本身进行操作，takewhile方法只是返回满足条件的元素
	intS.Filter(func(element interface{}) bool {
		return element.(int) >= 100
	}).Foreach(func(index int, item interface{}) {
		fmt.Println(item)
	})
