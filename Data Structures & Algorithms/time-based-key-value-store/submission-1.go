// string: timestamp 
// earlier  
// alice 3
// return alice 2
// SORTED 


//   0 1 2 3 [4] 5 6
//   l     m       r
//         l       r
//            m   

type TimeMapInfo struct {
	timestamp int
	value string
}


type TimeMap struct {
	store map[string]*[]TimeMapInfo
}

func Constructor() TimeMap {
	return TimeMap {
		store: make(map[string]*[]TimeMapInfo),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int) {
	storeInfo, ok := this.store[key]
	if !ok {
		storeInfoImp := make([]TimeMapInfo, 0, 3)
		storeInfo = &storeInfoImp
		this.store[key] = storeInfo
	}

	(*storeInfo) = append((*storeInfo), TimeMapInfo{timestamp: timestamp, value: value})
}

// alice [{1;happy}, {3:sad}]  

func (this *TimeMap) Get(key string, timestamp int) string {
	info, ok := this.store[key]
	if !ok {
		return ""
	}

	left := 0
	right := len((*info))-1 
	for left <= right{
		mid := left + (right-left)/2
		if (*info)[mid].timestamp == timestamp {
			return (*info)[mid].value
		} else if (*info)[mid].timestamp > timestamp {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	if right < 0 {
		return ""
	}
	return (*info)[right].value
}
