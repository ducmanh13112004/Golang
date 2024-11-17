package main

import (
	"fmt"
)

// Độ dài (Length): Số lượng phần tử có trong slice, truy cập bằng len(slice).
// Dung lượng (Capacity): Dung lượng tối đa mà slice có thể chứa, truy cập bằng cap(slice).
func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice1 := []int{1, 2, 3, 4, 5}
	//make
	slice2 := make([]int, 3, 5) // Slice với độ dài là 3 và dung lượng là 5
	slice3 := slice[1:3]
	//copy
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)
	copy(dest, src)
	//thêm phần tử vào slice
	slice = append(slice, 6, 7)
	// Thay đổi giá trị trong slice (ví dụ tại index 1)
	slice1[1] = 50  
	//       
	fmt.Println("---")
	fmt.Println("kết quả của copy: ", dest)
	fmt.Println("Kết quả thêm : ",slice)
    fmt.Println("Kết quả sau khi thay đổi : ",slice1)
	fmt.Println("---")
	fmt.Println(slice2)
	fmt.Println("---")
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice3)) //'sẽ in ra số lượng phần tử từ số 2 trể về sau

}