package Open_the_Lock

func openLock(deadends []string, target string) int {
	visited := make(map[string]struct{}, len(deadends))
	for _, v := range deadends {
		visited[v] = struct{}{}
	}

	var depth int
	var n = "0000"
	queue := []string{n}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			n, queue = queue[0], queue[1:]
			// 已经访问过的节点直接continue，防止走回头路
			if _, has := visited[n]; has {
				continue
			}
			if n == target {
				return depth
			}
			// 将该节点加入已访问的节点
			visited[n] = struct{}{}
			// 将其他8个相邻组合加入queue
			for j := 0; j < 4; j++ {
				plus, minus := plusOne(n, j), minusOne(n, j)
				if _, has := visited[plus]; !has {
					queue = append(queue, plus)
				}
				if _, has := visited[minus]; !has {
					queue = append(queue, minus)
				}
			}
		}
		depth++
	}
	return -1
}

// 最优解：双向BFS
func openLock_double_bfs(deadends []string, target string) int {
	visited := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		visited[v] = true
	}

	// 使用集合不使用队列
	q1 := map[string]bool{"0000": true}
	q2 := map[string]bool{target: true}

	var depth int
	for len(q1) > 0 && len(q2) > 0 {
		// 优化：总是扩散空间小的集合
		if len(q1) > len(q2) {
			temp := q1
			q1 = q2
			q2 = temp
		}

		// hashset遍历时不能修改，用temp暂存结果
		temp := map[string]bool{}

		// 将q1中的所有结果向外扩散
		for key := range q1 {
			if visited[key] {
				continue
			}
			// q2中拥有q1的元素，说明有交集，就是最终答案
			if q2[key] == true {
				return depth
			}
			visited[key] = true

			for i := 0; i < 4; i++ {
				plus, minus := plusOne(key, i), minusOne(key, i)
				if !visited[plus] {
					temp[plus] = true
				}
				if !visited[minus] {
					temp[minus] = true
				}
			}
		}
		depth++

		// 下一轮扩散q2
		q1 = q2
		q2 = temp
	}
	return -1
}

func plusOne(s string, bit int) string {
	bt := []byte(s)
	if bt[bit] == '9' {
		bt[bit] = '0'
	} else {
		bt[bit]++
	}
	return string(bt)
}

func minusOne(s string, bit int) string {
	bt := []byte(s)
	if bt[bit] == '0' {
		bt[bit] = '9'
	} else {
		bt[bit]--
	}
	return string(bt)
}
