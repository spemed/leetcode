package arr

/**
59. 螺旋矩阵 II
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20
*/

type point struct {
	x int
	y int
}

func (p *point) less(other point) bool {
	return p.x < other.x && p.y < other.y
}

func (p *point) equal(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p *point) greater(other point) bool {
	return p.x > other.x && p.y > other.y
}

func generateMatrix(n int) [][]int {
	start, end := point{0, 0}, point{n - 1, n - 1}
	number := 1
	data := make([][]int, n)
	for i := range data {
		data[i] = make([]int, n)
	}
	for start.equal(end) || start.less(end) {

		if start.equal(end) {
			data[start.x][start.y] = number
		}

		for i := start.x; i < end.x; i++ {
			data[start.y][i] = number
			number++
		}

		for i := start.y; i < end.y; i++ {
			data[i][end.x] = number
			number++
		}

		for i := end.x; i > start.x; i-- {
			data[end.y][i] = number
			number++
		}

		for i := end.y; i > start.y; i-- {
			data[i][start.x] = number
			number++
		}

		start.x++
		start.y++
		end.x--
		end.y--
	}

	return data
}
