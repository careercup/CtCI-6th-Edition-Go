package chapter1

func ZeroMatrixB(matrix [][]int) {

	rowCnt, columnCnt := len(matrix), len(matrix[0])
	if rowCnt == 0 {
		return
	}
	var zeroFirstCol, zeroFirstRow bool
	for j := 0; j < columnCnt; j++ {
		if matrix[0][j] == 0 {
			zeroFirstRow = true
			break
		}
	}
	for i := 0; i < rowCnt; i++ {
		if matrix[i][0] == 0 {
			zeroFirstCol = true
			break
		}
	}
	for i := 1; i < rowCnt; i++ {
		for j := 1; j < columnCnt; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for j := 1; j < columnCnt; j++ {

		if matrix[0][j] == 0 {
			ZeroColumn(matrix, j)
		}

	}

	for i := 1; i < rowCnt; i++ {

		if matrix[i][0] == 0 {
			ZeroRow(matrix, i)
		}

	}
	if zeroFirstCol {
		ZeroColumn(matrix, 0)
	}

	if zeroFirstRow {
		ZeroRow(matrix, 0)
	}

}

func ZeroRow(matrix [][]int, row int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
}

func ZeroColumn(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

func ZeroMatrix(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for r := 0; r < len(rows); r++ {
		if rows[r] {
			ZeroRow(matrix, r)
		}
	}
	for c := 0; c < len(cols); c++ {
		if cols[c] {
			ZeroColumn(matrix, c)
		}
	}
}
