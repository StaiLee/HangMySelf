package art

import (
	"fmt"
)

func Victory() {
	colorGreen := "\033[32m"
	colorWhite := "\033[37m"
	fmt.Println(string(colorGreen), "")
	fmt.Println("||             ||  |||||||||||||||||||   ||               ||    ||                                        ||     ||||||||||||||||||||    ||||             ||")
	fmt.Println(" ||           ||   ||               ||   ||               ||     ||                                      ||               ||             || ||            ||")
	fmt.Println("  ||         ||    ||               ||   ||               ||      ||                                    ||                ||             ||  ||           ||")
	fmt.Println("   ||       ||     ||               ||   ||               ||       ||                                  ||                 ||             ||   ||          ||")
	fmt.Println("    ||     ||      ||               ||   ||               ||        ||                                ||                  ||             ||    ||         ||")
	fmt.Println("     ||   ||       ||               ||   ||               ||         ||                              ||                   ||             ||     ||        ||")
	fmt.Println("      || ||        ||               ||   ||               ||          ||            ||||            ||                    ||             ||      ||       ||")
	fmt.Println("      ||||         ||               ||   ||               ||           ||          ||  ||          ||                     ||             ||       ||      ||")
	fmt.Println("      ||||         ||               ||   ||               ||            ||        ||    ||        ||                      ||             ||        ||     ||")
	fmt.Println("      ||||         ||               ||   ||               ||             ||      ||      ||      ||                       ||             ||         ||    ||")
	fmt.Println("      ||||         ||               ||   ||               ||              ||    ||        ||    ||                        ||             ||          ||   ||")
	fmt.Println("      ||||         ||               ||   ||               ||               ||  ||          ||  ||                         ||             ||           ||  ||")
	fmt.Println("      ||||         |||||||||||||||||||   |||||||||||||||||||                ||||            ||||                 |||||||||||||||||||     ||            |||||")
	fmt.Println(string(colorWhite), "")
}
