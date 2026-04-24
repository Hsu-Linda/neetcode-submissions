func groupAnagrams(strs []string) [][]string {
	// key  = [26]int is for frequency
	// value is string list

	anagramMap := make(map[[26]int][]string)
	for _, str := range strs {
		frequency := [26]int{}
		for _, char := range str {
			frequency[char - 'a'] ++
		}
		anagramMap[frequency] = append(anagramMap[frequency], str)
	}

	output := make([][] string, 0, len(anagramMap))
	for _, group := range anagramMap {
		output = append(output, group)
	}

	return output
}
