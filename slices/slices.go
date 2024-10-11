package main

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func main() {
	// 1
	hobbies := [3]string{"programming", "reading", "playing games"}
	println(hobbies[0], hobbies[1], hobbies[2])

	// 2
	println(hobbies[0])
	newHobbies := hobbies[1:3]
	println(newHobbies[0], newHobbies[1])

	// 3
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	println(slice1[0], slice1[1], slice2[0], slice2[1])

	// 4
	slice1 = hobbies[1:3]

	// 5
	courseGoals := []string{"learn go", "have fun with go"}

	// 6
	courseGoals[1] = "learn go more"
	courseGoals = append(courseGoals, "learn go even even more")

	// 7
	type Product struct {
		title string
		id    int
		price float64
	}

	products := []Product{
		{title: "product1", id: 1, price: 1.0},
		{title: "product2", id: 2, price: 2.0},
	}

	products = append(products, Product{title: "product3", id: 3, price: 3.0})
}
