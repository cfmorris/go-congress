package update

import "fmt"

func CheckUpdate() {
	fmt.Println("update something? (y/n)")
	var doUpdate string
	fmt.Scanln(&doUpdate)
	if doUpdate == "y" {
		fmt.Println("update members? (y/n)")

		var first string

		fmt.Scanln(&first)

		if first == "y" {
			GetSenators()
			GetRepresentatives()
		}

		fmt.Println("update comittees? (y/n)")

		var updateComittee string
		fmt.Scanln(&updateComittee)

		if first == "y" {
			fmt.Println("Committee Updates Are Not Ready")
		}
	}
}
