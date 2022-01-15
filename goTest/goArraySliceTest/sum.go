package goarrayslicetest

func Sum(arr []int) (sum int) {
	sum = 0
	for _, v := range arr {
		sum += v
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
