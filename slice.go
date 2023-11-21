package main

import "fmt"

func main() {
	 month := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	 }

	 slice1 := month[4:7]
	 slice2 := month[:4]
	 slice3 := month[2:]
	 slice4 := month[:]

	 fmt.Println(slice1, slice2, slice3, slice4)
	 fmt.Println(len(slice1))
	 fmt.Println(cap(slice1))

	 fmt.Println("-------------")

	 days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thurssday",
		"Friday",
		"Saturday",
		"Sunday",
	 }

	 daySlice1 := days[5:]
	 daySlice1[0] = "New Saturday"
	 daySlice1[1] = "New Sunday"

	 fmt.Println(daySlice1)
	 fmt.Println(days)

	 daySlice2 := append(daySlice1, "New Holiday")
	 fmt.Println(daySlice2)
	 fmt.Println(daySlice1)

	 fmt.Println("-------------")

	 var newSlice []string = make([]string, 2, 5)
	 newSlice[0] = "Muhtarom"
	 newSlice[1] = "Zain"
	//  newSlice[2] = "Muhammad" ,akan terjadi error tidak seperti array, jika ingin menambahkan data diluar panjangnya maka perlu menggunakan append

	 fmt.Println(newSlice)
	 fmt.Println(len(newSlice))
	 fmt.Println(cap(newSlice))

	 fmt.Println("-------------")

	 newSlice2 := append(newSlice, "Muhammad")
	 fmt.Println(newSlice2)
	 fmt.Println(len(newSlice2))
	 fmt.Println(cap(newSlice2))

	 fmt.Println("-------------")

	 fromSlice := days[:]
	 toSlice := make([]string, len(fromSlice), cap(fromSlice))

	 copy(toSlice, fromSlice)

	 fmt.Println(fromSlice)
	 fmt.Println(toSlice)

	 fmt.Println("-------------")

	 iniArray := [...]int{1,2,3,4,5}
	 iniSlice := []int{1,2,3,4,5} // slice tidak ada jumlah pada tipe datanya

	 fmt.Println(iniArray)
	 fmt.Println(iniSlice)
}