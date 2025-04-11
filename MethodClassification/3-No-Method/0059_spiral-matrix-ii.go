package main

import "fmt"

// 主函数，演示生成并打印不同大小的矩阵
func main() {
	// 生成并打印一个3x3的矩阵
	matrix := generateMatrix(3)
	fmt.Println(matrix)
	// 生成并打印一个4x4的矩阵
	matrix = generateMatrix(4)
	fmt.Println(matrix)
}

// generateMatrix 生成一个大小为 n x n 的矩阵，其中填充了从 1 到 n*n 的数字，以螺旋顺序排列。
// 参数 n: 矩阵的行数和列数
// 返回值: 一个填充了数字的二维切片矩阵
func generateMatrix(n int) [][]int {
	// 初始化一个 n x n 的二维切片
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 初始化计数器和边界指针
	count := 1
	left, right := 0, n-1
	top, bottom := 0, n-1

	// 以螺旋顺序填充矩阵
	for count <= n*n {
		// 从左到右填充上边界
		for i := left; i <= right; i++ {
			matrix[top][i] = count
			count++
		}
		top++
		// 从上到下填充右边界
		for i := top; i <= bottom; i++ {
			matrix[i][right] = count
			count++
		}
		right--
		// 从右到左填充下边界
		for i := right; i >= left; i-- {
			matrix[bottom][i] = count
			count++
		}
		bottom--
		// 从下到上填充左边界
		for i := bottom; i >= top; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}
	// 返回填充完成的矩阵
	return matrix
}
