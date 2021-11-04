// 1.「)」must put after each 「(」.
// 2. There are must exactly n 「(」 in each result.
// we use "count" to record the number of「(」
// and "status" to promise there are enough「(」before we append「)」to the result.
// Then we can span the tree recursively.
package main

func span(frag string, status, count int, list *[]string) {
	if status == 0 && count == 0 {
		*list = append(*list, frag)
	}
	if count > 0 {
		span(frag+"(", status+1, count-1, list)
	}
	if status > 0 {
		span(frag+")", status-1, count, list)
	}
}

func generateParenthesis(n int) []string {
	list := []string{}
	span("(", 1, n-1, &list)
	return list
}
