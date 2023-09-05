package TwoPoint

import (
	"fmt"
	"sort"
)

/*
27. 移除元素
提示
简单
1.9K
相关企业
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

// 使用双指针（前后指针）
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1

	for left < right {
		// 从左边找到第一个等于 val 的元素（找到第一个等于的时候，跳出）
		for left < right && nums[left] != val {
			left++
		}
		// 从右边找到第一个不等于 val 的元素（找到第一个不等于的时候，跳出）
		for left < right && nums[right] == val {
			right--
		}
		// 交换左右指针所指向的元素
		nums[left], nums[right] = nums[right], nums[left]
	}

	// 判断最后一个元素是否为 val
	if nums[left] == val {
		return left
	}
	// 返回新的数组长度
	return left + 1
}

// 快慢指针法
func RemoveElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}
	nums = nums[:result]
	fmt.Println("the nums is:", nums)
	return result
}

// 26. 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

// 283. 移动零
func MoveZeroes(nums []int) {
	// 使用双指针方法
	// left 指针用于指向当前非零元素应该存放的位置
	// right 指针用于遍历数组元素

	left := 0

	// 遍历数组
	for right := 0; right < len(nums); right++ {
		// 如果当前元素不为零，则将其移动到 left 指针所指的位置，并将 left 指针右移
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	// 将剩余的位置填充为零
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

/*
844. 比较含退格的字符串
简单
662
相关企业
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。

示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。

提示：

1 <= s.length, t.length <= 200
s 和 t 只含有小写字母以及字符 '#'
*/

// 844. 比较含退格的字符串
func BackspaceCompare(s string, t string) bool {
	return buildModifiedString(s) == buildModifiedString(t)
}

func buildModifiedString(s string) string {
	slow := 0
	fast := 0
	str := []byte(s)
	for fast < len(str) {
		if str[fast] != '#' {
			str[slow] = str[fast]
			slow++
		} else {
			if slow > 0 {
				slow--
			}
		}
		fast++
	}
	fmt.Println("the string is:", string(str[:slow]))

	return string(str[:slow])
}

func buildModifiedString1(s string) string {
	str := []byte(s)
	slow := 0

	for fast := 0; fast < len(str); fast++ {
		if str[fast] != '#' {
			str[slow] = str[fast]
			slow++
		} else if slow > 0 {
			slow--
		}
	}

	return string(str[:slow])
}

/*
977. 有序数组的平方
简单
873
相关企业
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

// 977. 有序数组的平方（非双指针）
func SortedSquares(nums []int) []int {
	for k, v := range nums {
		nums[k] = v * v
	}
	sort.Ints(nums)
	fmt.Println(nums)
	return nums
}

// 双指针实现
func SortedSquares1(nums []int) []int {
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1
	result := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			result[index] = nums[right] * nums[right]
			right--
		} else {
			result[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return result
}

/*
LCR 008. 长度最小的子数组
中等
126
相关企业
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
*/

// LCR 008. 长度最小的子数组-----使用滑动窗口
func MinSubArrayLen(target int, nums []int) int {
	// 如果数组为空，返回0
	if len(nums) == 0 {
		return 0
	}

	start := 0              // 子数组的起始索引
	sum := 0                // 子数组元素的和
	length := len(nums) + 1 // 最小子数组长度的初始值为数组长度加1

	// 遍历数组元素
	for end := 0; end <= len(nums)-1; end++ {
		sum += nums[end] // 将当前元素添加到子数组的和中

		// 如果子数组的和大于等于目标值
		for sum >= target {
			length = min(length, end-start+1) // 更新最小子数组长度

			// 从子数组的起始位置移除元素，继续查找更小的子数组
			sum -= nums[start]
			start++
		}
	}

	// 如果最小子数组长度没有被更新过，则返回0，表示找不到满足条件的子数组
	if length == len(nums)+1 {
		return 0
	}

	return length // 返回最小子数组的长度
}

func min(x, y int) int {
	// 返回两个整数中的较小值
	if x < y {
		return x
	}
	return y
}

// LeetCode59题
/*
题目：螺旋矩阵 II

题目描述：
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按照螺旋顺序排列的 n x n 正方形矩阵。

示例:
输入: 3
输出:
[
[ 1, 2, 3 ],
[ 8, 9, 4 ],
[ 7, 6, 5 ]
]
*/
/*
解法：
这道题可以使用模拟的方法来生成螺旋矩阵。我们可以定义四个变量来表示当前生成元素的位置，
分别是上边界（top）、下边界（bottom）、左边界（left）和右边界（right）。初始时，上边界和左边界都为0，下边界和右边界都为n-1。然后按照螺旋顺序依次生成矩阵元素。

具体的生成过程如下：

初始化一个 n x n 的矩阵，所有元素都设置为0。
定义一个变量 num，初始值为1，表示当前要生成的元素。
使用一个循环，循环条件为 top <= bottom 且 left <= right，表示还有元素需要生成。
在循环中，按照螺旋顺序生成元素：
从左到右：将 num 赋值给矩阵中的元素 matrix[top][i]，同时 num 自增1。
从上到下：将 num 赋值给矩阵中的元素 matrix[i][right]，同时 num 自增1。
从右到左：将 num 赋值给矩阵中的元素 matrix[bottom][i]，同时 num 自增1。
从下到上：将 num 赋值给矩阵中的元素 matrix[i][left]，同时 num 自增1。
更新边界值：top、bottom、left、right 分别加1或减1。
循环结束后，返回生成的矩阵。
*/
func GenerateMatrix(n int) [][]int {
	// 创建一个二维数组
	matrix := make([][]int, n)
	// 为三行赋值为0
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	//设置上下左右的区间
	left, right, top, bottom := 0, n-1, 0, n-1
	nums := 1
	//for nums<=n*n
	for top <= bottom && left <= right {
		//从左到右(为第一行后n-1个赋值)
		for i := left; i <= right; i++ {
			matrix[top][i] = nums
			nums++
		}
		top++ //移动到第二行(这样可以直接跳过第一行最后一个)

		//从上到下(为n-1列后n-1个赋值)
		for i := top; i <= bottom; i++ {
			matrix[i][right] = nums
			nums++
		}
		right-- //移动到第二列(这样可以直接跳过第n-1列最后一个)

		//从右到左(为n-1行前n-1个赋值)
		for i := right; i >= left; i-- {
			matrix[bottom][i] = nums
			nums++
		}
		bottom-- //移动到第一行(这样可以直接跳过第二行第一个)

		//从下到上(为第一列前n-1个赋值)
		for i := bottom; i >= top; i-- {
			matrix[i][left] = nums
			nums++
		}
		left++ //移动到第二行(这样可以直接跳过第一行最后一个)
	}
	return matrix
}

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for top <= bottom && left <= right {
		//从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		//从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		//从右到左
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		//从下到上
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++

		if left > right || top > bottom {
			break
		}
	}
	return result
}
