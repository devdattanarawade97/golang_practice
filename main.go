package main

import (
	"fmt"
	"reflect"
	"sort"
)

//closure function
func closure_function() func() int {
     a:=10;
	return func () int {
		a++;
		return a;
	}
}


func defer_function() {
	fmt.Println("defer function")
}


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

	//call by reference 
	y:=10;
	z:=20;
	fmt.Println("substraction is : ",minus(&y,&z));
   
	//call anoymous function
//anynoumous function
anoynomous_function:=func() {
	fmt.Println("anonymous function")
	
 }

anoynomous_function()


//call closure function
fmt.Println("closure function",closure_function()())
fmt.Println("before defer function")
//defer function	
defer defer_function()
fmt.Println("defer function end")


//call list function
fmt.Println("list function",list_function(1,2,3,4,5,6))


//array 
var arr [5]int=[5]int{1,2,3,4,5};
fmt.Println("array value",arr)

//elipsis 
var elipsis_arr =[...]int{1,2,3,4,5};
fmt.Println("elipsis array value",elipsis_arr)

//struct array 
record:=[]struct {
	name string
	age int
}{
	{
		name:"hello",
		age:10,
	},

}

fmt.Println("struct value",record)


//arr using make arr
int_arr:=make([]int,5,10);
fmt.Println("int array value",int_arr)


//2d slices 
double_array_slices:=[][]int{
	{1,2,3},
	{4,5,6},
	{7,8,9},
};

fmt.Println("2d slices value",double_array_slices)

//copy
slice_2:=[]int {5,20,4};
slice_1 := make([]int, len(slice_2))

copy(slice_1,slice_2);

fmt.Println("slices2 copy value",slice_2);
fmt.Println("slices1 copy value",slice_1);

//sort arr
sort.Ints(slice_1);
fmt.Println("slices1 sort value",slice_1);

//append slice 
slice_1=append(slice_1,20);
fmt.Println("slices1 append value",slice_1);

//compare two arr 
fmt.Println("compare two arr",reflect.DeepEqual(slice_1,slice_2));


//string 
var string_var string="hello";
fmt.Println("string value",string_var)

//iterate over string
for i, char := range string_var {
	fmt.Printf("Character at index %d: %c\n", i, char)
}

//string 
var string_var_1 string=`hello_
1`;
fmt.Println("string value",string_var_1)
}

//argument list function
func list_function(arg...int) int {
	sum:=0;
	for _,v:=range arg {
		sum+=v
	}
	return sum
}


func sum(a int , b int) int {
     
	return a+b
}

func minus(a *int , b *int) int {			

	return *a-*b}


