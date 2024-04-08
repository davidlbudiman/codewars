package katas

var subPartitions = map[int]int{
	0: 1, 1: 1, 2: 2, 3: 3, 4: 5, 5: 7, 6: 11, 7: 15, 8: 22, 9: 30, 10: 42, 11: 56, 12: 77,
	13: 101, 14: 135, 15: 176, 16: 231, 17: 297, 18: 385, 19: 490, 20: 627, 21: 792,
	22: 1002, 23: 1255, 24: 1575, 25: 1958, 26: 2436, 27: 3010, 28: 3718, 29: 4565,
	30: 5604, 31: 6842, 32: 8349, 33: 10143, 34: 12310, 35: 14883, 36: 17977,
	37: 21637, 38: 26015, 39: 31185, 40: 37338, 41: 44583, 42: 53174, 43: 63261,
	44: 75175, 45: 89134, 46: 105558, 47: 124754, 48: 147273, 49: 173525,
}

func Partitions(n int) int {
	if n >= len(subPartitions) {
		calculateSubPartitions(n)
	}
	return subPartitions[n]
}

func calculateSubPartitions(n int) {
	x, sum := n, 0
	for sign, nextTurn, i, j := 1, 1, 1, 1; x >= 0; nextTurn++ {
		if nextTurn%2 == 1 {
			x -= i
		} else {
			x -= j
		}
		if subPartitions[x] == 0 {
			calculateSubPartitions(x)
		}
		sum += sign * subPartitions[x]
		if nextTurn%2 == 1 {
			i += 2
		} else {
			j++
			sign = -1 * sign
		}
	}
	subPartitions[n] = sum
}
func Partitions_vaskiny(n int) int {
	table := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int, n+1)
	}

	for i := 0; i <= n-1; i++ {
		table[0][i] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-1; j++ {
			if i-j < 0 {
				table[i][j] = table[i][j-1]
				continue
			}
			table[i][j] = table[i][j-1] + table[i-j][j]
		}

	}

	return table[n][n-1] + 1
}

func Partitions_lancatlin(n int) int {
	return p_lancatlin(n, n)
}

func p_lancatlin(n int, d int) int {
	if d > n {
		d = n
	}
	if d <= 1 {
		return 1
	}
	count := 0
	for i := d; i > 0; i-- {
		count += p_lancatlin(n-i, i)
	}
	return count
}
