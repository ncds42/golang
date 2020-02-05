package main;

import "fmt"


func main(){
	last, current := 1, 1;
	for (current < 100000){
		last, current = fib(last,current);
	}
	fmt.Println(last);
}


func fib(last, current int ) (int, int){
	return (current), (last + current);
}