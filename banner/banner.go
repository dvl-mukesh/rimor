package banner

import "fmt"

var banner = `
8888888b.  8888888 888b     d888  .d88888b.  8888888b.  
888   Y88b   888   8888b   d8888 d88P" "Y88b 888   Y88b 
888    888   888   88888b.d88888 888     888 888    888 
888   d88P   888   888Y88888P888 888     888 888   d88P 
8888888P"    888   888 Y888P 888 888     888 8888888P"  
888 T88b     888   888  Y8P  888 888     888 888 T88b   
888  T88b    888   888   "   888 Y88b. .d88P 888  T88b  
888   T88b 8888888 888       888  "Y88888P"  888   T88b                                                                                                                                         
`

var (
	purple = color("\033[1;34m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func PrintBanner() {
	fmt.Println(purple(banner))
}
