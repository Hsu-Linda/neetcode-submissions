//   /: [/] [//] [///]   slase as only one slash
//   .: current
//   .. : parent  
//    more than two dot is valid directory

// /neetcode/practice//...///../course
// neetcode  practice  "" ... "" "" ..  course
// neetcode practice course
func simplifyPath(path string) string {
	pathArray := strings.Split(path, "/")
	pathStack := make([]string, 0 , len(pathArray))
	for _, f := range pathArray {
		if f == "" || f == "." {
			continue
		} else if f == ".." {
			if len(pathStack) > 0 {
				pathStack = pathStack[:len(pathStack)-1]
			}
		} else {
			pathStack = append(pathStack, f)
		}
	}
	return "/"+ strings.Join(pathStack,"/")
}