package gateway

import (
	"fmt"
	"ssh-gateway/counter"
)


func PrintUserCount() {
	usersCount := counter.GlobalUserCount
	fmt.Println(usersCount)
	fmt.Println("Current Number of Users in Containers:")
	for containerAddress, count := range usersCount.GetCounts() {
		fmt.Printf("Container %s: %d users\n", containerAddress, count)
	}
	fmt.Println("------------------------")
}