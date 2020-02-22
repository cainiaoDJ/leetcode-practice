package main

import "fmt"

func quickSortCustom(nums *[]int, start int, end int) {
	// 起点和终点重合的时候，退出
	if start >= end {
		return
	}
	i, j := start, end
	//左节点的坑挖出来，备用
	mid := (*nums)[start]
	for {
		// 左节点位移到跟右节点重合时退出for循环
		if i >= j {
			break
		}

		// 右节点开始向左查找，直到比基准值小
		for {

			if j > i && (*nums)[j] >= mid {
				j--
			} else {
				break
			}
		}
		// 把右节点的坑，填入左节点。现在右节点j空出来了
		if i < j {
			(*nums)[i] = (*nums)[j]
			i++
		}

		// 左节点开始向右寻找，直到找到比基准值大的
		for {
			if j > i && (*nums)[i] < mid {
				i++
			} else {
				break
			}
		}
		//把左节点的值填入右节点的坑，（上面的右节点j是空的没有变过）。此时左节点i空出来了
		if i < j {
			(*nums)[j] = (*nums)[i]
			j--
		}

	}
	// 把最开始挖出来的坑放到新坑里面
	(*nums)[i] = mid
	quickSortCustom(nums, start, i-1)
	quickSortCustom(nums, i+1, end)
}

type IntSlice []int

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	//startTime := time.Now()
	//for i := 1000000; i <= 10000000; i += 1000000 {
	//    numsIn := make(IntSlice, 0)
	//    numsMy := make([]int, 0)
	//    for n := 0; n < i; n++ {
	//        //rand.Seed(int64(n))
	//        x := rand.Intn(10000)
	//        numsIn = append(numsIn, x)
	//        numsMy = append(numsMy, x)
	//
	//    }
	//
	//    l := len(numsMy)
	//    //fmt.Println("原始:", numsMy, "排序个数：", l)
	//    fmt.Println("排序个数：", l)
	//
	//    // ---------------- 内置排序 开始----------------
	//    startTime = time.Now()
	//    sort.Ints(numsIn)
	//    cost := time.Since(startTime)
	//    fmt.Println("内置排序耗时：", cost)
	//    // ---------------- 内置排序 结束----------------
	//
	//    // ---------------- 自研排序 开始----------------
	//    startTime = time.Now()
	//    quickSortCustom(&numsMy, 0, l-1)
	//    cost = time.Since(startTime)
	//    fmt.Println("自研排序耗时：", cost)
	//    // ---------------- 自研排序 结束----------------
	//
	//    fmt.Println("结果抽查:", numsMy[l/2] == numsIn[l/2])
	//
	//    //fmt.Println("内置:", numsIn)
	//    //fmt.Println("自研:", numsMy)
	//
	//    fmt.Println(" ")
	//
	//}

	nums := []int{8081, 7887, 1847, 4059, 2081, 1318, 4425, 2540, 456, 3300}
	fmt.Println("原始的数据:", nums)
	quickSortCustom(&nums, 0, len(nums)-1)
	fmt.Println("自己排序后:", nums)
	//
	//nums2 := IntSlice{8081, 7887, 1847, 4059, 2081, 1318, 4425, 2540, 456, 3300}
	////fmt.Println("排序前:",nums2)
	//sort.Stable(nums2)
	//fmt.Println("内置排序后:", nums2)
}
