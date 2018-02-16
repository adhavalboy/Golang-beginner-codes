package main
import "fmt"
func factorial(x int64) int64 {
	if x==1 {
		return 1
	} 
		return x*factorial(x-1)
}
func main(){
add := func(x,y int) (int,int) {
return x+y, 5
}
c,d :=add(3,5)
fmt.Println(c,d)
fmt.Println(factorial(5));
}

