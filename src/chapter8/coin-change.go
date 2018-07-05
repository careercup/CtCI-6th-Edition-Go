package main
  
import "fmt"

func count(coin []int, sum int) int {
	numCoins := len(coin)
	C := make([][]int, sum+1)

	for i := 0; i <= sum; i++ {
		C[i] = make([]int, numCoins)
	}

	for i := 0; i < numCoins; i++ {
		C[0][i] = 1
	}

	for i := 1; i <= sum; i++ {
		for j := 0; j < numCoins; j++ {
			if i - coin[j] >= 0 {
				C[i][j] += C[i-coin[j]][j]
			}

			if j >= 1 {
				C[i][j] += C[i][j-1]
			}
		}
	}

	return C[sum][numCoins-1]
}

func main() {
	coin := []int{1, 5, 10, 25}
	fmt.Println(count(coin, 100))
}
