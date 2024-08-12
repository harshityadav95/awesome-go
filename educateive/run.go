package main
import (
	"fmt"
)

var total int

func main(){

	var sum int
	{
		var count int=5
		sum =calculateSum(&count,2)
		fmt.Println(sum)
		fmt.Println(count)
	}
	total=sum
	{
		count:= 10
		sum = calculateSum(&count,3)
		fmt.Println(sum)
		fmt.Println(count)

	}
	total := total+sum
	fmt.Printf("Total : %d is ",total)

}
func calculateSum(count *int, factor int) int {
	
	sum:= *count * factor
	return sum
}