package sliceutil

// SliceUtil slice结构体
type SliceUtil struct {
	// interface{}代表任何类型
	Elements []interface{}
}

// HasElement 判断是否有某个元素,成功返回true,否则返回false
func (s SliceUtil) HasElement(element interface{}) bool {
	for _, v := range s.Elements {
		// 找到元素
		if v == element {
			return true
		}
	}
	return false
}

// Length 返回元素个数
func (s SliceUtil) Length() int {
	return len(s.Elements)
}

// Add 增加元素, 可以一次性增加多个
func (s *SliceUtil) Append(element interface{}, args ...interface{}) SliceUtil {

	s.Elements = append(s.Elements, element)
	if len(args) > 0 {
		s.Elements = append(s.Elements, args...)
	}
	return (*s)
}

// Delete 按照元素本身来删除元素
func (s *SliceUtil) Delete(i int) SliceUtil {

	if s.Elements == nil || len(s.Elements) == 0 {
		return (*s)
	}
	if i < 0 || i >= len(s.Elements) {
		panic("index out of bounds")
	}

	copy(s.Elements[i:], s.Elements[i+1:])
	s.Elements = s.Elements[:len(s.Elements)-1]

	return (*s)
}

// 将指定位置的元素修改为新元素
func (s *SliceUtil) Modify(index int, newElement interface{}) SliceUtil {
	s.Elements[index] = newElement

	return (*s)
}

// Where 查找到element所在的index,如果找不到返回-1
func (s SliceUtil) Where(toFindElement interface{}) int {
	for i, element := range s.Elements {
		//找到所在元素
		if element == toFindElement {
			return i
		}
	}
	return -1
}

// 通过index找到元素;，其中索引由用户自己判断
func (s SliceUtil) FindByIndex(index int) interface{} {
	return s.Elements[index]
}

// Foreach 遍历每个元素，在回调函数内对元素进行操作, 可以参考js的forEach
func (s SliceUtil) Foreach(callback func(index int, item interface{})) {
	for index, element := range s.Elements {
		callback(index, element)
	}
}

// Filter filter: 进行操作时需要类型断言，比如ele1.(string)
// filter是对元素本身就行了操作
func (s SliceUtil) Filter(condition func(item interface{}) bool) SliceUtil {
	// 这样可以chain invoke
	var s1 SliceUtil
	s1.Elements = []interface{}{}
	s.Foreach(func(currentIndex int, eachElement interface{}) {
		// 如果条件成立: 使用用户传入的回调函数作为判断条件
		if condition(eachElement) {
			// 添加元素
			s1.Elements = append(s1.Elements, eachElement)
		}
	})
	return s1
}

// MapTo 将每一个元素映进行射，
//
//	slice1.MapTo(func(item interface{}){
//			return item.(int)*10;
//	})

func (s SliceUtil) MapTo(callback func(item interface{}) interface{}) SliceUtil {
	s.Foreach(func(currentIndex int, eachElement interface{}) {
		// 对每一个元素就行重新赋值
		s.Elements[currentIndex] = callback(eachElement)
	})
	return s
}

// AllSatisfied 判断每个元素是否都满足条件, 是，返回true; 否则返回false
func (s SliceUtil) AllMatch(judge func(item interface{}) bool) bool {
	for _, element := range s.Elements {
		// 如果有元素不满足条件，立刻返回
		if !judge(element) {
			return false
		}
	}
	return true
}

// -1表示找不到，否则返回index
func (s *SliceUtil) findFirst(toFindElement interface{}) int {

	for i, element := range s.Elements {
		// 判断是否符合条件
		if toFindElement == element {
			return i
		}
	}
	return -1
}

// -1表示找不到，否则返回index
func (s *SliceUtil) findLast(element interface{}) int {
	for i := len(s.Elements) - 1; i >= 0; i-- {
		if s.Elements[i] == element {
			return i
		}
	}
	return -1
}

// 返回满足条件的所有元素,不对原来的切片进行操作
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
