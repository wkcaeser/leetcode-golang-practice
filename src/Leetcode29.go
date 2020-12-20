package main

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}

	if dividend == ^int(^uint32(0)>>1) && divisor == -1 {
		return int(^uint32(0) >> 1)
	}

	underZero := dividend^divisor < 0

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	rs := quick(dividend, divisor)

	if underZero {
		return -rs
	} else {
		return rs
	}
}

func quick(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}

	n2 := 31
	for (dividend >> n2) < divisor {
		n2--
	}

	rs := 1 << n2

	rs += divide(dividend-divisor<<n2, divisor)

	return rs
}
