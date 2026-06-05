type Solution struct{}

// ["2#ab", "2#cb"]
// 
func (s *Solution) Encode(strs []string) string {
	output := ""
	for _, str := range strs {
		count := len(str)
		output += strconv.Itoa(count)
		output += "#"
		output += str
	}
	
	return output
}

// 8#2#ab2#cd
func (s *Solution) Decode(encoded string) []string {
	index := 0
	output := make([]string, 0, len(encoded)/4)
	
	for index < len(encoded) {
		hashIndex := index
        for encoded[hashIndex] != '#' {
			hashIndex += 1
		}
        
        length, _ := strconv.Atoi(encoded[index:hashIndex])
        index = hashIndex+1
        word := encoded[index:(index+length)]
        output = append(output, word)
        index = index+length
	}
	return output
}
