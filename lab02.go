package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name string
	email string
	phone string
	gender string
}

func main()  {
	list := map[int]Student {
		1: {"Khang Tran", "khangtran@gmail.com", "0987654321", "Nam"},
		2: {"Mukund", "mukund@go.dev", "0123456789", "Nữ"},
		3: {"Trung Tran", "trungtran@gmail.com", "0987654321", "Nam"},
		4: {"Ngan Hoai", "nganhoai098@gmail.com", "0868246105", "Nữ"},
	}
	//getStudentList(list)
	addStudent(list, Student{"Tuan Nguyen", "tuannguyen@gmail.com", "0987654321", "Nam"})
	// deleteStudent(list, 1)
	updateStudent(list, 2, Student{"Huyen Pham", "huyen@go.dev", "0987654321", "Nữ"})
	// getStudentByKey(list, 3)
	// searchStudentByName(list, "huyen")
	// getStudentByGender(list, "Nam")
}

func getStudent(student Student, key int)  {
	fmt.Println("Mã số: ", key)
	fmt.Println("Tên sinh viên: ", student.name)
	fmt.Println("Email: ", student.email)
	fmt.Println("SĐT: ", student.phone)
	fmt.Println("Giới tính: ", student.gender)
	fmt.Println("-----------------------------")
}

func getStudentList(list map[int]Student)  {
	fmt.Println("DANH SÁCH SINH VIÊN: ")
	fmt.Println("-----------------------------")
	for key, student := range list {
		getStudent(student, key)
	}
}

func getStudentByKey(list map[int]Student, key int)  {
	if student, ok := list[key]; ok {
		fmt.Printf("Tìm thấy sinh viên có mã số %d: \n", key)
		getStudent(student, key)
	} else {
		fmt.Println("Không tìm thấy sinh viên!")
	}
}

func addStudent(list map[int]Student, student Student)  {
	list[len(list) + 1] = student
	fmt.Println("Thêm sinh viên thành công!")
	getStudentList(list)
}

func deleteStudent(list map[int]Student, key int)  {
	delete(list, key)
	fmt.Println("Xóa sinh viên thành công!")
	getStudentList(list)
}

func updateStudent(list map[int]Student, key int, student Student)  {
	list[key] = student
	fmt.Println("Cập nhật sinh viên thành công!")
	getStudentList(list)
}

func searchStudentByName(list map[int]Student, keyword string)  {
	fmt.Printf("DANH SÁCH SINH VIÊN CÓ TỪ KHÓA: %s\n", strings.ToUpper(keyword))
	fmt.Println("-----------------------------")
	for key, student := range list {
		if strings.Contains(strings.ToLower(student.name), strings.ToLower(keyword)) {
			getStudent(student, key)
		}
	}
}

func getStudentByGender(list map[int]Student, gender string)  {
	fmt.Printf("DANH SÁCH SINH VIÊN CÓ GIỚI TÍNH: %s\n", strings.ToUpper(gender))
	fmt.Println("-----------------------------")
	for key, student := range list {
		if strings.EqualFold(student.gender, gender) {
			getStudent(student, key)
		}
	}
}