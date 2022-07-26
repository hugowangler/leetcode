package leetcode

func parse(expr string, index int) (bool, int) {
	var values []bool
	op := expr[index]
	var continueIndex int
	for i := index + 2; i < len(expr); {
		if expr[i] == ',' {
			i++
		}
		if expr[i] == ')' {
			continueIndex = i + 1
			break
		}
		if expr[i] == '&' || expr[i] == '|' || expr[i] == '!' {
			value, j := parse(expr, i)
			values = append(values, value)
			i = j
		} else {
			values = append(values, expr[i] == 't')
			i++
		}
	}
	var res bool
	switch op {
	case '&':
		res = true
		for _, value := range values {
			res = res && value
		}
	case '|':
		res = false
		for _, value := range values {
			res = res || value
		}
	case '!':
		res = !values[0]
	}
	return res, continueIndex
}

func parseBoolExpr(expression string) bool {
	res, _ := parse(expression, 0)
	return res
}
