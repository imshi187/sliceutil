package main

// SliceUtil slice结构体
type SliceUtil struct {
	// interface{}代表任何类型
	elements []interface{}
}

// HasElement 是否有某个元素
func (s SliceUtil) HasElement(element interface{}) bool {
	for _, v := range s.elements {
		// 找到元素
		if v == element {
			return true
		}
	}
	return false
}

// Length 返回长度
func (s SliceUtil) Length() int {
	return len(s.elements)
}

// Add 增加元素
func (s *SliceUtil) Append(element interface{}, args ...interface{}) SliceUtil {

	s.elements = append(s.elements, element)
	if len(args) > 0 {
		s.elements = append(s.elements, args...)
	}
	return (*s)
}

// Delete 按照元素本身来删除元素 (尚未成功)
func (s *SliceUtil) Delete(i int) SliceUtil {

	if s.elements == nil || len(s.elements) == 0 {
		return (*s)
	}
	if i < 0 || i >= len(s.elements) {
		panic("index out of bounds")
	}

	copy(s.elements[i:], s.elements[i+1:])
	s.elements = s.elements[:len(s.elements)-1]

	return (*s)
}

// modify
func (s *SliceUtil) Modify(index int, element interface{}) SliceUtil {
	s.elements[index] = element

	return (*s)
}

// Where 查找到element所在的index
func (s SliceUtil) Where(toFindElement interface{}) int {
	for i, element := range s.elements {
		//找到所在元素
		if element == toFindElement {
			return i
		}
	}
	return -1
}

// 找到元素 by index
// 索引由用户自己判断
func (s SliceUtil) FindByIndex(index int) interface{} {
	return s.elements[index]
}

// Foreach foreach成功实现
func (s SliceUtil) Foreach(callback func(index int, item interface{})) {
	for index, element := range s.elements {
		callback(index, element)
	}
}

// Filter filter: 就行操作时需要类型断言，比如ele1.(string)
// filter是对元素本身就行了操作
func (s SliceUtil) Filter(condition func(item interface{}) bool) SliceUtil {
	// 这样可以chain invoke
	var s1 SliceUtil
	s1.elements = []interface{}{}
	s.Foreach(func(currentIndex int, eachElement interface{}) {
		// 如果条件成立: 使用用户传入的回调函数作为判断条件
		if condition(eachElement) {
			// 添加元素
			s1.elements = append(s1.elements, eachElement)
		}
	})
	return s1
}

// MapTo 将每一个元素映射
func (s SliceUtil) MapTo(callback func(item interface{}) interface{}) SliceUtil {
	s.Foreach(func(currentIndex int, eachElement interface{}) {
		// 对每一个元素就行重新赋值
		s.elements[currentIndex] = callback(eachElement)
	})
	return s
}

// AllSatisfied 判断每个元素是否都满足条件
func (s SliceUtil) AllMatch(judge func(item interface{}) bool) bool {
	for _, element := range s.elements {
		// 如果有元素不满足条件，立刻返回
		if !judge(element) {
			return false
		}
	}
	return true
}

// -1表示找不到，否则返回index
func (s *SliceUtil) findFirst(toFindElement interface{}) int {

	for i, element := range s.elements {
		// 判断是否符合条件
		if toFindElement == element {
			return i
		}
	}
	return -1
}

// -1表示找不到，否则返回index
func (s *SliceUtil) findLast(element interface{}) int {
	for i := len(s.elements) - 1; i >= 0; i-- {
		if s.elements[i] == element {
			return i
		}
	}
	return -1
}

// 返回满足条件的所有元素,不对原来的切片就行操作
func (s SliceUtil) TakeWhile(callback func(ele interface{}) bool) []interface{} {
	result := make([]interface{}, 0)
	s.Foreach(func(index int, item interface{}) {
		// 如果满足条件
		if callback(item) {
			result = append(result, item)
		}
	})

	return result

}

//func main() {
//	var intS SliceUtil
//	// SliceUtil结构体的field仅仅包括[]interface{}类型的elements
//	intS.elements = []interface{}{1, 2, 3, 100, 200}
//
//	fmt.Println("------------------------------------")
//	fmt.Println(intS.TakeWhile(func(ele interface{}) bool {
//		// 一般都是同一类型的数据，这里是int类型
//		return ele.(int) > 2
//	}))
//
//	// filter会对s本身就行操作，takewhile方法只是返回满足条件的元素
//	intS.Filter(func(element interface{}) bool {
//		return element.(int) >= 100
//	}).Foreach(func(index int, item interface{}) {
//		fmt.Println(item)
//	})
//
//}
