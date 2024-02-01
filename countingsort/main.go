package main

import "fmt"

func main() {
	Q3()
}

func Q1() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 4, 5, 6, 7, 8, 2, 3, 10, 4, 5, 6}

	/* 값이 몇개 있는지 확인 */
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	// fmt.Println(count)

	/* 값에 갯수에 따라 정렬 */
	/* 방법 1
	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	*/
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	fmt.Println(sorted)
}

func Q2() {
	str := "dsanfjksdnubjhsdfgegadsyfasdfsadfwaedasf"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + 1)
		}
	}

	fmt.Printf("%c %d\n", maxCh, maxCount)
}

func Q3() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyte", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Uejdf", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
		{Name: "Kevin", Height: 164.8},
	}

	// for i := 0; i < len(students); i++ {
	// 	if students[i].Height >= 170.0 && students[i].Height < 180.0 {
	// 		fmt.Println(students[i].Name, students[i].Height)
	// 	}
	// }

	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		heightMap[int(students[i].Height*10)] = append(heightMap[int(students[i].Height*10)], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name: ", name, " height: ", float64(i)/10)
		}
	}

	// 180 ~ 210
	for i := 1800; i < 2100; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name: ", name, " height: ", float64(i)/10)
		}
	}

}
