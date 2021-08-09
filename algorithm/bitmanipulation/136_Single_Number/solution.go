package main

// first solution
func singleNumber(nums []int) (rst int) {
	table := make(map[int]struct{}, len(nums)/2)
	for _, v := range nums {
		// 如果数字存在就删除
		if _, has := table[v]; has {
			delete(table, v)
		} else {
			table[v] = struct{}{}
		}
	}
	// table中剩下的唯一一个是未重复的数字
	for i := range table {
		rst = i
		break
	}
	return
}

// best solution
func singleNumber_XOR(nums []int) (rst int) {
	for _, v := range nums {
		rst ^= v
	}
	return
}
