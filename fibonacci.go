package main
import "fmt"


func fib(number int) int {
	if number==0{
		return 0
	}else if number==1{
		return 1
	}else {
		return fib(number-1)+ fib(number-2)
	}

}
func main(){
	var iteration int
	fmt.Println("Hi ! I am mirmahyar ! May I ask you to enter the number of fibonacci elements to be printed? ")
	fmt.Scanf("%d",&iteration)
	fmt.Printf("Ok, Thank you! Here is the first %d elements of fibonacci sequence: \n",iteration)
	for i:=0;i<iteration;i++{
		fmt.Println(fib(i))
	}

}