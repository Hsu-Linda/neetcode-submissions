// ["1","2","+","C","5","D"]
// sum 6
// 1 2 

// 3
// 1+2 5 10
func calPoints(operations []string) int {
    sum := 0
    newOpe := make([]int, 0, len(operations))
    
    for i:=0; i<len(operations);i++ {
        switch operations[i] {
        case "+":
            add := newOpe[len(newOpe)-1]+ newOpe[len(newOpe)-2]
            sum += add
            newOpe = append(newOpe, add)
        case "D":
            sum += 2*newOpe[len(newOpe)-1]
            newOpe = append(newOpe, 2*newOpe[len(newOpe)-1])
        case "C":
            num1 := newOpe[len(newOpe)-1]
            sum -= num1
            newOpe = newOpe[:len(newOpe)-1]
        default:
            num1, _ := strconv.Atoi(operations[i])
            sum += num1
            newOpe = append(newOpe, num1)
        }
    }
    return sum
    
}
