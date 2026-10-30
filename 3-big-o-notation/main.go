package main

import (
	"fmt"
	"strings"
)

// 1. Constant O(1): وقت ثابت بغض النظر عن حجم المدخلات
func addNumbers(a,b float64) float64{
	return  a + b
}

// 2. Linear O(n): الوقت يتناسب طردياً مع حجم الـ Slice
func addNumbersBySlice(nums []float64)float64{
	total := 0.0 // O(1)
	for _,n := range nums { // O(n)
		total += n // O(1)
	}
	return total
}

type Student struct {
	Id int
	Name string
}

// 3. Quadratic O(n^2): حلقات تكرارية متداخلة (Nested Loops)
func findStudentById(school [][]Student, studentId int) *Student {
	for _, classRoom := range school { // O(n2)
		for _, student := range classRoom { // O(n)
			if studentId == student.Id { // O(1)
				return &student
			}
		}
	}
	return nil
}

type Contact struct {
	Name string
	Phone string
}

// 4. Logarithmic O(log n): البحث الثنائي في دليل مرتب أبجدياً
func searchPhone(phonebook []Contact, name string) *Contact {
	// {1,2,3,4,5}
	low := 0
	high := len(phonebook) - 1

	for low <= high {
		mid := low + (high - low) / 2

		res := strings.Compare(phonebook[mid].Name, name)

		if res == 0 {
			return &phonebook[mid] // تم العثور على الاسم
		} else if res < 0 {
			low = mid + 1 // الاسم المستهدف بعد العنصر الحالي
		} else {
			high = mid - 1 // الاسم المستهدف قبل العنصر الحالي
		}
	}

	return  nil
}

// 5. Logarithmic-Linear O(n log n):
// المرور على المحافظات O(n) * البحث الثنائي داخل دليل كل محافظة O(log m)
func searchAcrossGovernorates(countryPhonebook map[string][]Contact, name string) *Contact {
	for _, phonebook := range countryPhonebook {
		foundContact := searchPhone(phonebook, name)
		if foundContact != nil {
			return foundContact
		}
	}
	return nil
}

var Counter int = 0

// 6. Exponential O(2^n): شجرة استدعاءات تتضاعف مع كل زيادة في n
func fibonacci(n int) int {
	Counter++
	fmt.Println(Counter)
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main(){
	//  result := addNumbers(1, 62)

	//  println(result)

	// ----------------------------------------

	// result := addNumbersBySlice([]float64{1,5.5,70,6.3})

	// println(result)

	// ----------------------------------------

	// school := [][]Student{}
	// school = append(school, []Student{})
	// school[0] = append(school[0], Student{Id: 111, Name: "Ahmed Ouda"})
	// school[0] = append(school[0], Student{Id: 112, Name: "Ahmed Mahmoud"})
	// school[0] = append(school[0], Student{Id: 113, Name: "Ibrahim Kamal"})
	// school[0] = append(school[0], Student{Id: 114, Name: "Yasser Mohamed"})
	// school[0] = append(school[0], Student{Id: 115, Name: "Mahmoud Ahmed"})

	// school = append(school, []Student{})
	// school[1] = append(school[1], Student{Id: 210, Name: "Hamed Mohamed"})
	// school[1] = append(school[1], Student{Id: 211, Name: "Belal Ahmed"})

	// result := findStudentById(school, 211)

	// if result == nil {
	// 	fmt.Println("Student not found")
	// } else {
	// 	fmt.Println(result.Name)
	// }

	// ----------------------------------------

	// phonebook := []Contact{
	// 	{Name: "Ahmed", Phone: "01000000001"},
	// 	{Name: "Ebrahim", Phone: "01000000002"},
	// 	{Name: "Hassan", Phone: "01000000003"},
	// 	{Name: "Mohamed", Phone: "01000000004"},
	// 	{Name: "Omar", Phone: "01000000005"},
	// 	{Name: "Tarek", Phone: "01000000006"},
	// }

	// foundContact := searchPhone(phonebook, "Mohame")

	// if foundContact != nil {
	// 	fmt.Println(foundContact.Phone)
	// } else {
	// 	fmt.Println("Contact not found")
	// }

	// ----------------------------------------

	// countryPhonebooks := map[string][]Contact{
	// 	"Cairo": {
	// 		{Name: "Ahmed", Phone: "01011111111"},
	// 		{Name: "Hassan", Phone: "01022222222"},
	// 		{Name: "Ziad", Phone: "01033333333"},
	// 	},
	// 	"Alexandria": {
	// 		{Name: "Ali", Phone: "01111111111"},
	// 		{Name: "Mona", Phone: "01122222222"},
	// 		{Name: "Tarek", Phone: "01133333333"},
	// 	},
	// 	"Giza": {
	// 		{Name: "Ebrahim", Phone: "01211111111"},
	// 		{Name: "Mohamed", Phone: "01222222222"},
	// 		{Name: "Sara", Phone: "01233333333"},
	// 	},
	// }

	// foundContactInCountry := searchAcrossGovernorates(countryPhonebooks, "Moh")

	// if foundContactInCountry != nil {
	// 	fmt.Println(foundContactInCountry.Phone)
	// } else {
	// 	fmt.Println("Contact not found")
	// }

	// ----------------------------------------

	// fibonacci(20)
}

/* 
	1,000,000 contact

	linear O(n) = 1,000,000

	logarithmic O(log(n)) = 1,000,000 / 2
	500,000 / 2
	250,000 / 2
	125,000 / 2
	62,500 / 2
	31,250 / 2
	15,625 / 2
	7,812 / 2
	3,906 / 2
	1,953 / 2
	976 / 2
	488 / 2
	244 / 2
	122 / 2
	61 / 2
	30 / 2
	15 / 2
	7 / 2
	3 / 2
	1 / 2
	0
*/