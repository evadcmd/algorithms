// 1.「)」must put after each 「(」.
// 2. There are must exactly n 「(」 in each result.
// we use "count" to record the number of「(」
// and "status" to promise there are enough「(」before we append「)」to the result.
// Then we can span the tree recursively.
package main

import "bytes"

func span(buf *bytes.Buffer, status, count int, list *[]string) {
	if status == 0 && count == 0 {
		*list = append(*list, buf.String())
		return
	}
	if count > 0 {
		buf.WriteString("(")
		span(buf, status+1, count-1, list)
		buf.Truncate(buf.Len() - 1)
	}
	if status > 0 {
		buf.WriteString(")")
		span(buf, status-1, count, list)
		buf.Truncate(buf.Len() - 1)
	}
}

func generateParenthesis(n int) []string {
	list := []string{}
	// span(bytes.NewBuffer(make([]byte, 0, n*2)), 0, n, &list)
	span(new(bytes.Buffer), 0, n, &list)
	return list
}
