package main

import "fmt"

// func main(){
// 	// array , slice , map

// 	scores := []float64{90,85,80,82}

// 	fmt.Println(scores)

// 	scores = append(scores, 5, 6, 7, 8, 90)

// 	fmt.Println(scores)
// 	fmt.Println(scores[2])
// 	fmt.Println(scores[5])
// }

// func main(){
// 	// array , slice , map

// 	scoresMap := map[string]([]float64){}

// 	fmt.Println(scoresMap)

// 	scoresMap["ahmed"] = []float64{90,80,85,76}

// 	scoresMap["mahmoud"] = []float64{87,80,95,93}

// 	// fmt.Println(scoresMap)
// 	fmt.Println(scoresMap["ahmed"])
// 	fmt.Println(scoresMap["Ahmed"])
// }


func main(){
	// array , slice , map

	classes := [](map[string]int){}

	classes = append(classes, map[string]int{})

	classes[0]["ahmed mohamed"] = 123456
	classes[0]["ahmed Essa"] = 123457
	classes[0]["ahmed Hamdy"] = 123458
	classes[0]["Hamed Mohamed"] = 123460

	classes = append(classes, map[string]int{})

	classes[1]["Mahoud ahmed"] = 130044
	classes[1]["Mahoud Ibrahim"] = 130045
	classes[1]["Mohamed Rabeea"] = 130090

	fmt.Println(classes)
}

