package GoSlice

import (
	"strconv"
)

type Slice struct {
	slice []interface{}
}

// 初始化一个切片
func NewSlice() *Slice {
	return &Slice{ slice: make([]interface{}, 0) }
}

// 压栈一个元素
func (slice *Slice) Push(element interface{}) {
	slice.slice = append(slice.slice, element)
}

// 删除并返回数切片的第一个元素
func (slice *Slice) Shift() (element interface{}, ok bool) {
	var firstEle, isDelete = slice.Delete(0)
	if !isDelete {
		return  nil,false
	}
	return firstEle, true
}

// 向切片的开头添加一个或更多元素，并返回新的长度。(存在问题，暂不使用)
//func (slice *Slice) Unshift(args...interface{}) (length int) {
//	if len(args) == 0 {
//		return slice.Length()
//	}
//	arg := make([]interface{}, 0)
//	for _, val :=  range args {
//		arg = append(arg, val)
//	}
//	fmt.Println(arg...)
//	slice.Splice(0,0, arg...)
//	return slice.Length()
//}

// 添加或删除数组中的元素,添加的元素在下标为index的元素之前
func (slice *Slice) Splice(index int, howmany int, args...interface{}) (sliceDeleted []interface{}, ok bool) {
	Deleted := make([]interface{}, 0)
	length := slice.Length()
	if howmany < 0 || index < 0 || index >= length {
		return Deleted, false
	}
	// 若删除空切片的元素，则返回false
	if length == 0 && (howmany != 0 || index != 0) {
		return Deleted, false
	}
	// 指定删除个数超出切片元素个数
	if howmany+index > length {
		copy(Deleted, slice.slice[0:length-1])
		// 删除所有指定元素,howmany为0则会跳过删除
		for i:= 0; i<length-index; i++ {
			slice.Delete(index)
		}
	} else {
		copy(Deleted, slice.slice[index:index+howmany])
		// 删除所有指定元素,howmany为0则会跳过删除
		for i:= 0; i<howmany; i++ {
			slice.Delete(index)
		}
	}

	// 没有需要添加的元素就直接返回了
	if len(args) == 0 {
		return Deleted, true
	}

	// 若把index之后的内容删完了，连index都没有，只有index-1是切片的最后一个元素
	if slice.Length() == index {
		slice.slice = append(slice.slice, args...)
		return Deleted, true
	}

	newSlice := make([]interface{}, 0)
	newSlice = append(slice.slice[:index], append(args, slice.slice[index:]...)...)
	slice.slice = newSlice
	return Deleted, true
}

// 弹出栈顶元素（切片的最后一个元素）
func (slice *Slice) Pop() (element interface{}, ok bool) {
	if slice.Length() == 0 {
		return  nil,false
	}
	item := slice.slice[0]
	slice.slice = slice.slice[:copy(slice.slice[0:], slice.slice[1:])]
	return item, true
}

// 删除指定下标元素
func (slice *Slice) Delete(index int) (element interface{}, ok bool) {
	if index >= slice.Length() || index < 0 {
		return nil, false
	}
	var deleted = slice.slice[index]
	slice.slice = slice.slice[:index+copy(slice.slice[index:], slice.slice[index+1:])]
	//fmt.Println(slice.slice)
	return deleted, true
}

// 获得指定下标元素
func (slice *Slice) Get(index int) (element interface{}, ok bool) {
	if index >= slice.Length() || index < 0 {
		return nil, false
	}
	return slice.slice[index], true
}

// 判断元素是否在切片中
func (slice *Slice) Includes(element interface{}) bool {
	if slice.Length() == 0 {
		return false
	}
	for _, item := range slice.slice {
		if item == element {
			return true
		}
	}
	return false
}

// 返回指定元素第一次出现在切片中的位置,没有返回-1
func (slice *Slice) IndexOf(element interface{}) int {
	if slice.Length() == 0 {
		return -1
	}
	for index, item := range slice.slice {
		if item == element {
			return index
		}
	}
	return -1
}

// 将数组所有元素组成字符串(仅支持字符、字符串、整型)，不成功返回false
func (slice *Slice) Join() (strJoined string, ok bool) {
	if slice.Length() == 0 {
		return "", true
	}
	var str = ""
	for _, item := range slice.slice {
		if val1, ok1 := item.(string); ok1 != false {
			// 如果是字符串类型元素
			str = str + val1
			continue
		} else if val2, ok2 := item.(int); ok2 != false {
			// 如果是整型
			str = str + strconv.Itoa(val2)
			continue
		} else if val3, ok3 := item.(byte); ok3 != false{
			str = str + string(val3)
		} else if val4, ok4 := item.(rune); ok4 != false{
			str = str + string(val4)
		} else {
			return "", false
		}
	}
	return str, true
}

// 反转切片元素
func (slice *Slice) Reverse() {
	var length = slice.Length()
	var swap interface{}
	// 0个或1个元素，保持原切片
	if length == 0 || length == 1 {
		return
	}

	// 如果只有两个元素就直接交换
	if length == 2 {
		swap = slice.slice[0]
		slice.slice[1] = slice.slice[0]
		slice.slice[0] = swap
		return
	}

	// 判断元素个数奇偶
	var isOven = false
	if length%2 == 0 {
		isOven = true
	} else {
		isOven = false
	}

	// 反转
	var right = length - 1
	var left = 0

	if isOven {
		for right-1 != left {
			swap = slice.slice[right]
			slice.slice[right] = slice.slice[left]
			slice.slice[left] = swap
			left++
			right--
		}
	} else {
		for right != left {
			swap = slice.slice[right]
			slice.slice[right] = slice.slice[left]
			slice.slice[left] = swap
			left++
			right--
		}
	}

}

// 截取切片一部分，不会改变原切片,包含下标为start的元素，不包含下标为end的元素
func (slice *Slice) Slice(start int, end int) (sliced []interface{}, ok bool) {
	newSlice := make([]interface{}, 0)
	if start >= end || start < 0 || end < 0 || start > slice.Length() - 1 || end > slice.Length() - 1 {
		return newSlice, false
	}
	// 若end已经是最后一个元素
	if end == slice.Length()-1 {
		copy(newSlice, slice.slice[start:])
		return newSlice, true
	}
	// 否则
	copy(newSlice, slice.slice[start:end+1])
	return newSlice, true
}

// 取得长度
func (slice *Slice) Length() int {
	return len(slice.slice)
}

// 取得容量
func (slice *Slice) Cap() int {
	return cap(slice.slice)
}

