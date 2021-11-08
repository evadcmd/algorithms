package main

func maxScoreSightseeingPair(values []int) int {
	max := 0
	for i, prev := 0, 0; i < len(values); i, prev = i+1, prev-1 {
		// According to conditions: values[i] >= 1,
		// it is save to set max's (and prev's) default value to 0.
		// First iteration: tmp := values[0] + 0 which is > 0
		// => max would be set to value[0] as default
		if tmp := values[i] + prev; tmp > max {
			max = tmp
		}
		if values[i] > prev {
			prev = values[i]
		}
	}
	return max
}
