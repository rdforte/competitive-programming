package main

import "fmt"

func main() {
	t := Constructor()
	t.Set("love", "1", 1)
	t.Set("love", "3", 3)
	t.Set("love", "4", 4)
	t.Set("love", "5", 5)
	t.Set("love", "6", 6)
	fmt.Println(t.Get("love", 2))
}

type TimeMap map[string][]TimestampVal

func Constructor() TimeMap {
	keyValStore := make(map[string][]TimestampVal)
	return keyValStore
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	(*this)[key] = append((*this)[key], TimestampVal{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := (*this)[key]; !ok {
		return ""
	} else {
		if timestamp < (*this)[key][0].stamp {
			return ""
		}

		l, u := 0, len((*this)[key])-1
		for l < u {
			mid := l + (u-l)/2

			if timestamp > (*this)[key][mid].stamp {
				l = mid + 1
			} else {
				u = mid
			}
		}

		if (*this)[key][l].stamp > timestamp && l > 0 {
			return (*this)[key][l-1].val
		}

		return (*this)[key][l].val
	}
}

type TimestampVal struct {
	val   string
	stamp int
}
