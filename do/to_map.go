package do

// ToMap: convert array to map
// this way it looks like an array and we can delete
func ToMap(arr ...interface{}) (m map[int]interface{}) {

	m = make(map[int]interface{})

	for id, e := range arr {
		m[id] = e
	}

	return
}
