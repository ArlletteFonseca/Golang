package main



type num []int


func newEven() num {
	numbers := num{0,1,2,3,4,5,6,7,8,9,10}
	// even := []int {2,4,6,8,10}
    // typeofNum := []string {}
	for _, num := range numbers {
		if(num % 2){
			fmt.Println(num + "is even")
		} else {
			fmt.Println(num + " is odd")
		}

}
}