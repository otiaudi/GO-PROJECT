package main
import (
	"fmt"
)
/* ARRAYS
	Declare and initialize an array
	Access elements in an array
	Iterate over an array
	Modify elements in an array
	Multidimensional arrays
	Declare and initialize an array 
	func todo(){
		 
		 arr1 :=[]int{1,2,3,4,5}
		 for i:=0; i<len(arr1); i++{
			 fmt.Println(arr1[i])
		 } // Iterate over an array
	

		 arr2 :=[] string{"hello","world"}
		 arr2=append(arr2, "that is the good news")
		 fmt.Println(arr2)
	}


*/


/* fuction to demonstrate the use of pointers

	func swap(m1 *int, m2 *int) {
	// Swap the values of m1 and m2
	// Use a temporary variable to hold the value of m1
	// Swap the values

	 temp := *m2 // Store the value of m2 in temp
		*m2=*m1 // Assign the value of m1 to m2
		*m1 =temp // Assign the value of temp to m1
	}
*/

/*
STRUCTURE
	Structs are used to create a custom data type that can hold multiple values of different types.
	structs are used to group related data together.
	it is a composite data type that can hold multiple values of different types.
	this is called data encapsulation.
	
	// Declare a struct
   type Car struct {
		name string
		age int
		modelNo int
   }

   // Method of the struct
   func(c Car) Drive(){
	fmt.Println("The car is driving")
   }
	// Method of the struct
	func(c Car) CarName() string{
		return c.name
	}
*/


/*INTERFACES
	Interfaces are used to define a contract that a struct must implement.
	Interfaces are used to define a set of methods that a struct must implement.
	Interfaces are used to enforce a style of programming that is more flexible and extensible.
	It helps you when you dont know the type of the data you are going to use.
	*/
type Car interface {
		Drive()
		//Stop()
}
	type Lambo struct{
		lamboModel string
	}

	type Toyota struct{
		toyotaModel string
	}

	func(l *Lambo) Drive(){
		fmt.Println("Lambo is moving")
		fmt.Println(l.lamboModel)
	}
	
	func(t *Toyota) Drive(){
		fmt.Println("Toyota is on the move")
		fmt.Println("it is of type:", t.toyotaModel)
	}
	func Anything(x interface{}){
		fmt.Println(x)
		}

func main(){
/*	fmt.Println("Arrays")
	todo()
	Slices
	Declare and initialize a slice
	Access elements in a slice
	Iterate over a slice
*/

/*POINTERS
	Declare and initialize a pointer
	---------------------------------------------
	p1, p2 :=5,9  
	fmt.Println(p1, p2) 
	swap(&p1, &p2) // call the swap function with pointers
	Swap the values of p1 and p2
	fmt.Println(p1, p2)	
	---------------------------------------------

*/

/*STRUCTS
	Declare and initialize a struct

	c:= Car{
		name:"Toyota", 
		age: 6,
		modelNo: 1234}
	Access the fields of the struct
	fmt.Println(c.name, c.age, c.modelNo)
	Call the method of the struct
	c.Drive()
  
	Call the method of the struct
	c.CarName()
	fmt.Println("The car name is", c.name)
*/

 l:=Lambo{"K234LONG20"}
 t:=Toyota{"CMT1245SLIN"}

 l.Drive()
 t.Drive()
 
 fmt.Println("this is")
 Anything("Silas Odero")
 Anything(5)
 Anything(5.5)
 Anything(true)
 Anything([]int{1,2,3,4,5})
 Anything(map[string]int{"one":1, "two":2})
 Anything(struct{Person string}{"Silas Odero"})
}



