package leetcode

func AddBinary(a string, b string) string {
	i, j := len(a), len(b)
	var x []rune
	var sum, carry rune
	for i > 0 && j > 0 {
		first, second := rune(a[i-1]), rune(b[j-1])
		sum, carry = sumAndCarry(first, second, carry)
		x = append(x, sum)
		i--
		j--
	}
	for i > 0 {
		sum, carry = sumAndCarry('0', rune(a[i-1]), carry)
		x = append(x, sum)
		i--
	}
	for j > 0 {
		sum, carry = sumAndCarry('0', rune(b[j-1]), carry)
		x = append(x, sum)
		j--
	}
	if carry == '1' {
		x = append(x, carry)
	}
	for k, l := 0, len(x)-1; k < l; k, l = k+1, l-1 {
		x[k], x[l] = x[l], x[k]
	}
	return string(x)
}

func sumAndCarry(first, second, carry rune) (sum, carry2 rune) {
	if first == '0' && second == '0' {
		sum, carry2 = '0', '0'
	} else if first == '1' && second == '1' {
		sum, carry2 = '0', '1'
	} else {
		sum, carry2 = '1', '0'
	}

	if carry == '1' {
		if sum == '0' {
			sum = '1'
		} else {
			sum, carry2 = '0', '1'
		}
	}
	return
}

func AddBinary_okhomin(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	indexB := len(b) - 1
	result := make([]byte, len(a))

	var shifter, sum byte
	for i := len(a) - 1; i >= 0; i-- {
		sum = shifter + a[i]
		if indexB >= 0 {
			sum += b[indexB]
			indexB--
		}

		result[i] = sum%2 + '0'
		shifter = sum >> 1 % 2
	}
	if shifter == 0 {
		return string(result)
	}

	return "1" + string(result)
}

func AddBinary_SamuelMuloki(a string, b string) string {
    ans, carry := "", 0
    for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
        sum := carry
        if i >= 0 {
            sum += int(a[i]-'0')
        }

        if j >= 0 {
            sum += int(b[j]-'0')
        }
    
        ans = string(byte(sum%2)+'0') + ans
        if sum <= 1 {
            carry = 0
        } else {
            carry = 1
        }
    }

    if carry == 1 {
        ans = string(byte(carry+'0')) + ans
    }

    return ans
}
