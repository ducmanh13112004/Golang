package main

import "fmt"

func main() {
	// Khởi tạo map với giá trị ban đầu
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   40,
	}

	// Thêm một phần tử mới
	ages["Charlie"] = 35

	// Cập nhật giá trị
	ages["Alice"] = 31

	// Truy cập và in một phần tử
	fmt.Println("Tuổi của Alice là:", ages["Alice"])

	// Kiểm tra phần tử tồn tại
	if age, exists := ages["Bob"]; exists {
		fmt.Println("Tuổi của Bob là:", age)
	} else {
		fmt.Println("Bob không tồn tại trong map")
	}

	// Xóa một phần tử
	delete(ages, "Eve")

	// Duyệt qua map và in các cặp key-value
	for name, age := range ages {
		fmt.Printf("%s có tuổi là %d\n", name, age)
	}
}
