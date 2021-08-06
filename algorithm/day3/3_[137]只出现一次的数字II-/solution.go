package main

// first solution
func singleNumber(nums []int) int {
	hashTable := make(map[int]int, len(nums))
	for _, v := range nums {
		hashTable[v]++
	}
	for k, v := range hashTable {
		if v == 1 {
			return k
		}
	}
	return 0
}

// best solution
func singleNumber_best(nums []int) int {
	var rst int32
	for i := 0; i < 32; i++ {
		var total int32
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			rst |= 1 << i
		}
	}
	return int(rst)
}