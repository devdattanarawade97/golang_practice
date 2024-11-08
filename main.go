package main

import (	
	"fmt"	
)		


func main() {
	name:= "Hello, World!"	
    
	//int value 
	var a int =2
	fmt.Println(name)
	fmt.Println(a)
    
	//pointers
	var c int=20;
	fmt.Println(c)	
	var d =&c;
	fmt.Println(d)
	fmt.Println(*d)
    //arr value
	var arrval [5]int=[5]int{1,2,3,4,5};
	fmt.Println(arrval)

	var mapvalue=make(map[int]string);
	mapvalue[1]="hello";
	mapvalue[2]="world";
	fmt.Println("map value",mapvalue)

	const i int =10
	fmt.Println("constant value",i);

	var e int=10;
	fmt.Println(e==c);

	 f :=30;
	fmt.Println(f)

	 x:=1;

	for x<10 {
		fmt.Println(x)
		x++
	}


	fmt.Println("sum : ",sum(1,2))

}


func sum(a int , b int) int {
     
	return a+b
}