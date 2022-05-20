package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
)

// Node is a sample type to be used as singly linked list
type Node struct {
	Val   int
	label string // not visible outside of package (starts with lowercase)
	Next  *Node
}

func main() {
	// Variables
	fmt.Println("*** Intro into variables ***")
	var name string // declaration of variable(will have default value)
	fmt.Printf("Default value of string %q\n", name)
	nickname := ""
	if name == nickname {
		fmt.Println("Different declarations of string yield same result")
	}
	a, b := "apple", "banana"                      // multiple declarations
	a, b = b, a                                    // swap without use of temp variable
	c := &b                                        // will have type of *string (pointer to string)
	_ = *c                                         // access value by dereferencing pointer
	n := &Node{Val: math.MaxInt, label: "example"} // type of *Node, math has MaxInt, MaxInt32, MaxInt64, as well as MinInt versions
	n.Next = new(Node)                             // dot access acts as (*n).Next=...

	var dfs func(num int) bool // variable can be a function, note: can be a func with no return value
	dfs = func(num int) bool {
		fmt.Printf("dfs(%d)...", num)
		if num == 0 {
			fmt.Println("Reached zero in dfs() call. Will return!")
			return true
		} else if num < 0 {
			fmt.Println("Only values >=0 are allowed")
			return false
		}
		return dfs(num - 1)
	}
	dfs(3)

	fmt.Printf("\nnext up:\n")
	// One dimensional lists
	fmt.Println("*** Lists demo ***")
	listOne := make([]int, 2)    // empty slice with len=2, init to default [0 0]
	listTwo := []int{3, 1, 2}    // definition with values
	listTwo = append(listTwo, 4) // adding a value
	sort.Ints(listTwo)           // sort in ascending order
	// swap first and last (or any valid index)
	listTwo[0], listTwo[len(listTwo)-1] = listTwo[len(listTwo)-1], listTwo[0]
	for i := 0; i < len(listOne); i++ { // iterating using index
		fmt.Printf("List One index: %d has value of %d\n", i, listOne[i])
	}
	for i, v := range listTwo { // for-range style, providing index and value
		fmt.Printf("List Two index: %d has value of %d\n", i, v)
	}
	for i := range listTwo { // only index
		_ = listTwo[i]
	}
	for _, v := range listTwo { // only values
		_ = v
	}
	listThree := listTwo
	listTwo[0] = 5
	fmt.Println("Slices reference underlying arrays:", listThree, listTwo)

	destList := make([]int, len(listTwo)) // destination should have enough capacity
	copy(destList, listTwo)               // returns num of elements copied
	listTwo[0] = 7
	fmt.Println("Copy helps to fill isolated slice:", destList, listTwo)

	fmt.Printf("\nnext up:\n")
	// Multi dimensional lists (matrix)
	fmt.Println("*** Multidimensional lists demo ***")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	for r := range matrix {
		for c := range matrix[0] {
			fmt.Printf("%d is the value as row=%d and column=%d\n", matrix[r][c], r, c)
		}
	}
	dp := make([][][]int, 2) // i,j,k
	for i := range dp {
		dp[i] = make([][]int, 3)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}
	fmt.Println(dp) // 2x3x4 matrix of zeros (default values)
	intervals := [][]int{{3, 5}, {2, 6}}
	// sort by ascending start
	// second param is less() func,
	// where intervals[i][0] is a start and intervals[i][1] is an end of i-th interval
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Println(intervals)

	fmt.Printf("\nnext up:\n")
	// Maps
	fmt.Println("*** Maps demo ***")
	mapOne := make(map[string]int, 4) // new map with element capacity hint (second arg is optional)
	mapTwo := map[string]int{"Alice": 3, "Bob": 5}
	for k, v := range mapTwo { // iterate using keys and values
		fmt.Printf("%s is %d\n", k, v)
	}
	if _, ok := mapOne["golang"]; !ok {
		mapOne["golang"] = 1 // element not present
	}
	if v, ok := mapOne["golang"]; ok {
		fmt.Printf("%s is found with value of %d\n", "golang", v)
	}

	// We can use comparable types such as fixed-size arrays as keys (slices cannot be used, they are not comparable)
	// Note struct{} is inline type definition of an empty struct,
	// which takes zero memory (in case you consider map[...]bool). Bool takes one byte, struct{} takes zero
	mapThree := make(map[[2]int]struct{})
	mapThree[[2]int{1, 3}] = struct{}{} // define (1,3) to be added to the map
	k := [2]int{1, 3}
	if _, ok := mapThree[k]; ok {
		fmt.Printf("Map entry with complex key %v is present!\n", k)
	}
	delete(mapThree, k) // delete key from map, also can delete(mapThree, [2]int{1, 3})
	if _, ok := mapThree[k]; !ok {
		fmt.Printf("Map entry with complex key %v is gone!\n", k)
	}

	fmt.Printf("\nnext up:\n")
	fmt.Println("*** Misc demo ***")
	fmt.Println("Call to custom variadic function:", max(4, 8, 1))

	fmt.Printf("\nnext up:\n")
	fmt.Println("*** Slice as queue or stack demo ***")
	queue := []int{}
	queue = append(queue, 1)                    // add element to a queue
	queue = append(queue, 2, 3, 4)              // add multiple elements
	queue = append(queue, []int{5, 6, 7, 8}...) // add multiple elements (from another list)
	queue = append([]int{42}, queue...)         // push value (i.e. 42) to front
	for len(queue) > 0 {
		fmt.Printf("%d...", queue[0])
		queue = queue[1:]
		// You pop multiple elements, and re-slice once. Can be used in BFS traversal.
		//queue[3:] means indexes 3,4,5,6...up to len(queue), i.e. queue[3:len(queue)]
	}
	fmt.Println("queue is now empty")

	stack := []int{}
	stack = append(stack, 1)                    // add element to a stack
	stack = append(stack, 2, 3, 4)              // add multiple elements
	stack = append(stack, []int{5, 6, 7, 8}...) // add multiple elements (from another list)
	for len(stack) > 0 {
		fmt.Printf("%d...", stack[len(stack)-1])
		stack = stack[:len(stack)-1] // last element index in slicing is non-inclusive
		//slice[1:5] means indexes 1,2,3,4
		//slice[:5] means indexes 0,1,2,3,4
	}
	fmt.Println("stack is now empty")

	fmt.Printf("\nnext up:\n")
	fmt.Println("*** Basic channels and concurrency demo ***")

	// second param is optional buffer (non-blocking up to buffer),
	// otherwise channels blocking from first message
	ch := make(chan int, 5)
	ch <- 3
	close(ch)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch { // reads up until channel is closed
			fmt.Printf("Got %d from chan of ints\n", num)
		}
	}()
	wg.Wait()

	// mutexes, can be used to guard concurrent access to variables and shared memory (use sparingly)
	// "share memory by communicating" is usually better pattern (using channels and pass over values)
	mu := &sync.Mutex{}
	mu.Lock()
	mu.Unlock()

	ch2 := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i + 1
		}
		close(ch2)
	}()

	ch = nil
loop:
	for {
		select {
		case num := <-ch: // nil chan always blocks, will not enter here. nil chan can be used to switch-off a case
			fmt.Printf("Will not get %d from nil chan\n", num)
		case num, ok := <-ch2: // second value indicates if chan in open (i.e. num is "ok" to use)
			// closed channels can be read from and they return default value (which is usually not "ok" to use)
			if ok {
				fmt.Printf("Got %d from chan of ints\n", num)
				continue
			}
			break loop // !ok, we are breaking out of for-select
		}
	}

	fmt.Printf("\nnext up:\n")
	fmt.Println("*** Strings demo ***")
	s := "Hello, 世界"
	for _, c := range s { // iterate over runes
		// if c='B', than c-'A' == 1
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c <= 255) { // a bit redundant, just to show comparison ops
			fmt.Printf("Non-ascii char: %c\n", c)
		}
	}
	for i := 0; i < len(s); i++ { //iterate over bytes
		_ = s[i]
	}
	for i, word := range strings.Split(s, ",") {
		fmt.Printf("Word #%d: %s\n", i+1, strings.TrimSpace(word))
	}
	fmt.Println("Checking strings.Contains() =>", strings.Contains(s, "llo"))     // true
	fmt.Println("Checking strings.HasPrefix() =>", strings.HasPrefix(s, "Hello")) // true
	fmt.Println("Checking strings.HasSuffix() =>", strings.HasSuffix(s, "界"))     // true
	fmt.Println("Checking strings.Join() =>", strings.Join([]string{"Hello", "World"}, " "))
	sb := strings.Builder{}
	_, err := sb.WriteString("Building formatted string")
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}
	fmt.Fprintf(&sb, " with some numbers: %6.2f and date: ", 12.5)
	//conversion to byte slice is not necessary (can use WriteString),
	//just demo of Write using bytes
	sb.Write([]byte(time.Now().Format("Mon Jan 02")))
	fmt.Println(sb.String())

	fmt.Printf("\nnext up:\n")
	fmt.Println("*** Heap demo ***")

	// minHeap := MinHeap{2, 1, 4, 3}
	// maxHeap := MaxHeap{MinHeap{2, 1, 4, 3}}
	// heap.Init(&maxHeap) // O(n) init
	// heap.Init(&minHeap)
	minHeap, maxHeap := MinHeap{}, MaxHeap{}
	for i := 1; i < 5; i++ {
		heap.Push(&minHeap, i)
		heap.Push(&maxHeap, i)
	}
	fmt.Printf("Max in max heap: %d\n", maxHeap.Peek())
	fmt.Printf("Min in min heap: %d\n", minHeap.Peek())
	for maxHeap.Len() > 0 { // or len(minHeap) >0, or minHeap.Len() >0
		fmt.Printf("Min: %d and max: %d\n", heap.Pop(&minHeap), heap.Pop(&maxHeap))
	}

}

// Variadic function: max(4,8,1) => 8
func max(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

// MinHeap is a min-heap of ints.
// will define container/heap.Interface functions (Len(),Less(),Swap(), plus Push(), Pop())
type MinHeap []int

// MaxHeap is a min-heap of ints.
type MaxHeap struct {
	MinHeap // we use struct embedding here, to get access to data and functions of existing type
	// with this style we will just need a small update to Less() function
}

// Also possible to define as `type MaxHeap []int`,
// which means defining 5 functions
// including 3 for sort interface func(Len(),Less(),Swap()), and 2 for heap-specific (Push(), Pop())

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//optional:
func (h *MinHeap) Peek() interface{} {
	return (*h)[0]
}

// For max heap we just need to redefine Less(), other functions are the same
func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }
