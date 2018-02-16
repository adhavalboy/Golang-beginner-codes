package main
import ("fmt")

func main() {
x := []int{1,2,3}
//var c [10]float64
c := make([]int, 2, 5) //this copy two element from x and capacity is 5
copy(c,x)
fmt.Println(x,c)

}
