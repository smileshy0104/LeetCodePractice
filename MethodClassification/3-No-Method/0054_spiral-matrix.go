package main

import "fmt"

// 主函数，演示生成并打印不同大小的矩阵
func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(arr))
}

// spiralOrder 以螺旋顺序遍历二维矩阵并返回一个一维数组
// 参数 matrix: 二维整数数组
// 返回值: 一维整数数组，包含以螺旋顺序遍历的结果
func spiralOrder(matrix [][]int) []int {
	nums := []int{} // 创建一个空的一维数组
	// 计算矩阵的行数和列数
	height := len(matrix)
	width := len(matrix[0])
	count := 0
	left, right := 0, width-1
	top, bottom := 0, height-1
	for count < height*width {
		// 从左到右遍历上边界
		for i := left; i <= right && count < height*width; i++ {
			nums = append(nums, matrix[top][i])
			count++
		}
		top++
		// 从上到下遍历右边界
		for i := top; i <= bottom && count < height*width; i++ {
			nums = append(nums, matrix[i][right])
			count++
		}
		right--
		// 从右到左遍历下边界
		for i := right; i >= left && count < height*width; i-- {
			nums = append(nums, matrix[bottom][i])
			count++
		}
		bottom--
		// 从下到上遍历左边界
		for i := bottom; i >= top && count < height*width; i-- {
			nums = append(nums, matrix[i][left])
			count++
		}
		left++
	}
	return nums
}

func spiralArray(array [][]int) []int {
	nums := []int{} // 创建一个空的一维数组
	if len(array) == 0 {
		return nums
	}
	// 计算矩阵的行数和列数
	height := len(array)
	width := len(array[0])
	count := 0
	left, right := 0, width-1
	top, bottom := 0, height-1
	for count < height*width {
		// 从左到右遍历上边界
		for i := left; i <= right && count < height*width; i++ {
			nums = append(nums, array[top][i])
			count++
		}
		top++
		// 从上到下遍历右边界
		for i := top; i <= bottom && count < height*width; i++ {
			nums = append(nums, array[i][right])
			count++
		}
		right--
		// 从右到左遍历下边界
		for i := right; i >= left && count < height*width; i-- {
			nums = append(nums, array[bottom][i])
			count++
		}
		bottom--
		// 从下到上遍历左边界
		for i := bottom; i >= top && count < height*width; i-- {
			nums = append(nums, array[i][left])
			count++
		}
		left++
	}
	return nums
}
