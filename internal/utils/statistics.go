package utils

func CalculateAverage(numbers []float64, chan4 chan float64) {
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}

	chan4 <- sum / float64(len(numbers))
}

func CalculateSum(numbers []float64, chan3 chan float64) {
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	chan3 <- sum
}

func CalculateMinimum(numbers []float64, chan1 chan float64) {
	min := numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	chan1 <- min
}

func CalculateMaximum(numbers []float64, chan2 chan float64) {
	max := numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	chan2 <- max
}
