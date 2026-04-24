// int int oper
// 4 13 5
// -> we found the third one is int not operation
// -> find next three eles, which means (ignore) the first ele, and find the next thrid one until it is operation
// store [4]
// 13 5 /
// 2

func evalRPN(tokens []string) int {
	// prev stack to store which we haven't calculate
	// prev we need to push , pop 
	// - push push at the end and pop get at the end first int last out
	// the condition of done is  both tokens and prevStack are empty
	// tokens alwasy get from the head: tokens[0]
	switch len(tokens) {
		case 0:
			return 0
		case 2:
		 	return 0
	}


	stackStruct := make([]int, 0, len(tokens))
	stack := &stackStruct
	for {
		valStr, popSuccess := popFirst(&tokens)
		if !popSuccess {
			if len(*stack) ==1 {
				return (*stack)[0]
			}
			// something wrong we can't find enough operation
			return 0
		}
		valInt, isInt := convertInt(valStr)
		if isInt {
			*stack = append(*stack, valInt)
			continue
		}

		second, isSecondPopSuccess := popLast(stack)
		first, isFirstPopSuccess := popLast(stack)
		
		if !isFirstPopSuccess || !isSecondPopSuccess {
			// somethine wend wrong we have operation but don't have numbers
			return 0
		}

		sum, evalIsSuccess := eval(first, second, valStr)
		if !evalIsSuccess {
			return 0
		}
		*stack = append(*stack, sum)
	}

}


// pop the first (index 0)
func popFirst(stack *[]string) (output string, isSuccess bool) {
	if len(*stack) == 0 {
		return "", false
	}
	
	output = (*stack)[0]
	*stack = (*stack)[1:]
	return output, true 
}


func popLast(stack *[]int) (output int, isSuccess bool) {
	if len(*stack) == 0 {
		return 0, false
	}
	output = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return output, true
}

func convertInt(val string) (output int, isSuccess bool) {
	output, convrErr := strconv.Atoi(val)
	if convrErr != nil {
		return 0, false
	}
	return output, true
}

func eval(first, second int, opera string) (output int, isSuccess bool) {
	switch opera {
		case "+":
			return first+second, true
		case "-":
			return first-second, true
		case "*":
			return first*second, true
		case "/":
			return first/second, true
		default :
			return 0, false
	}
}
