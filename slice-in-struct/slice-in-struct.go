//struct is passed by value but slice inside struct is reference type although if its an array instead of slice then it will be passed by value
//modifying the slice inside the struct will modify the original slice but if its an array then it will not modify the original array
//reassigning the slice inside the struct will not modify the original slice but if its an array then it will not modify the original array
//appending to the slice inside the struct will not modify the original slice

package main

import "fmt"

type SliceStruct struct {
	Data []int
}

func modSlice(s SliceStruct) {
	s.Data[0] = 100
}

func modData(s SliceStruct) {
	s.Data = []int{5, 4, 3, 2, 1}

}

func modAppendData(s SliceStruct) {
	s.Data = append(s.Data, 6, 7, 8)
}

func modAppendDataWithPointer(s *SliceStruct) {
	s.Data = append(s.Data, 7, 8, 9)
}

func main() {
	s := SliceStruct{
		Data: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("Slice in Struct:", s.Data)
	modSlice(s)
	fmt.Println("After modification:", s.Data)
	modData(s)
	fmt.Println("After modifying data:", s.Data)
	modAppendData(s)
	fmt.Println("After appending data:", s.Data)
	modAppendDataWithPointer(&s)
	fmt.Println("After appending data with pointer:", s.Data)

}
