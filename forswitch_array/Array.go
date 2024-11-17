//Tính tổng và giá trị trung bình của các phần tử trong một mảng
package main
import "fmt"
func testArray(){
	arr :=[...]float64{1,2,3,4,5}
	sum :=0.0
	
	for _,value :=range arr{
		sum +=value
	}

	average :=sum/float64(len(arr))
	fmt.Println("Tổng các phần tử trong mảng là:", sum)
	fmt.Println("Giá trị trung bình của các phần tử trong mảng là:", average)
}