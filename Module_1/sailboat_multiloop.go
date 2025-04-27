// Program to draw a sailboat with a sail, mast and a hull
// This was the first attempt and is more complex than simple approach

package main

import "fmt"

func main() {
	// Lengths for components
	const sail_len = 5
	const mast_len = 1
	const hull_len = 1
	const offset = 10 // For keeping some distance from the left

	// For hull, we define lengths of top and bottom
	const hull_base_1 = 2 * sail_len + 15
	const hull_base_2 = hull_base_1 - 4

	// Draw the sail
	for i := 0; i < sail_len; i++ {
		for j := 0; j < (sail_len - i) + offset; j++ {
			fmt.Printf(" ");
		}
		fmt.Printf("/");

		for j := 0; j < 2*i; j++ {
			fmt.Printf(" ");
		}
		fmt.Printf(" \\\n"); // Need to escape the "\"

	}

	// Draw the base
	for i := 0; i < offset; i++ {
		fmt.Printf(" ");
	}

	for i := 0; i < 2*sail_len + 3; i++ {
		fmt.Printf("-");
	}

	fmt.Println();

	// Draw the mast
	for i := 0; i < mast_len; i++ {
		for i := 0; i < sail_len + offset; i++ {
			fmt.Printf(" ");
		}
		fmt.Println("| |")
	}

	// Draw the hull

	// Upper base for hull
	for i := 0; i < offset; i++ {
		fmt.Printf(" ");
	}

	for i := 0; i < hull_base_1; i++ {
		fmt.Printf("-");
	}

	fmt.Println()

	for i := 0; i < hull_len; i++ {
		for j := 0; j < hull_len + offset + i; j++ {
			fmt.Printf(" ");
		}
		fmt.Printf("\\"); // Need to escape the "\"

		for j := 0; j < hull_base_2 - i; j++ {
			fmt.Printf(" ");
		}
		fmt.Println("/")

	}

	// Lower base for hull 
	for i := 0; i < offset + 2; i++ {
		fmt.Printf(" ");
	}

	for i := 0; i < hull_base_2; i++ {
		fmt.Printf("-");
	}

	fmt.Println()
}
