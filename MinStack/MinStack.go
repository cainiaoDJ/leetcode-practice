package main
import "fmt"
type MinStack struct {
	size int
	min int
	data []int
}

const initSize = 10

/** initialize your data structure here. */
func Constructor() MinStack {
	var m = MinStack{}
	// m.data = make( map[int]int,0)
	return m
}


func (this *MinStack) Push(x int)  {
	if(this.size==0 ||this.min>x){
		this.min = x
	}
	this.data = append(this.data,x)
	// this.data[this.top] = x
	this.size++
}


func (this *MinStack) Pop()  {
	if this.size == 0{
		return
	}

	if this.min == this.data[this.size-1] {
	this.min = this.data[0]
		for i:=1;i<this.size-1;i++{
			if(this.min>this.data[i]){
				this.min = this.data[i]
			}
		
		}
	}
	this.size--;
	this.data = append(this.data[:this.size])
	
}


func (this *MinStack) Top() int {
    return this.data[this.size-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}

func main(){
	obj := Constructor()
	obj.Push(7);
	obj.Push(9);
	obj.Push(12);
	obj.Pop()
	fmt.Println(obj)
	fmt.Println(obj.Top())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */