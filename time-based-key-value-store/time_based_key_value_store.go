package time_based_key_value_store

type TimeMap struct {
	timeMap map[string][]entry
}

type entry struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMap[key] = append(this.timeMap[key], entry{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vals, ok := this.timeMap[key]
	if !ok {
		return ""
	}

	var value string
	l, r := 0, len(vals)-1
	for l <= r {
		mid := (l + r) / 2

		if vals[mid].timestamp <= timestamp {
			value = vals[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return value
}
