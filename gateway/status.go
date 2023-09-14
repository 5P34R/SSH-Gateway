package gateway

import "fmt"

func PrintUserCount(userCount map[string]int) {
	fmt.Println("Current Number of Users in Containers:")
	for containerAddress, count := range userCount {
		fmt.Printf("Container %s: %d users\n", containerAddress, count)
	}
	fmt.Println("------------------------")
}
