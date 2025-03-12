package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/marco-lancini/goscan/core/cli"
	"github.com/marco-lancini/goscan/core/utils"
	"strings"
)

var (
	author  string
	version string
)

// ---------------------------------------------------------------------------------------
// INIT
// ---------------------------------------------------------------------------------------
func showBanner() {
	name := fmt.Sprintf("goscan (v.%s)", version)
	banner := `
 _____ _____ _____ _   _______  ___   ______  _______   __
/  ___|  ___/  __ \ | | | ___ \/ _ \  | ___ \|  _  \ \ / /
\ `--.| |__ | /  \/ | | | |_/ / /_\ \ | |_/ /| | | |\ V / 
 `--. \  __|| |   | | | |    /|  _  | | ___ \| | | |/   \ 
/\__/ / |___| \__/\ |_| | |\ \| | | | | |_/ /\ \_/ / /^\ \
\____/\____/ \____/\___/\_| \_\_| |_/ \____/  \___/\/   \/
                                                          
                                                          
	`

	// Shell width
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[1])

	// Print Centered
	fmt.Println(banner)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}

func initCore() {
	// Check sudo
	utils.CheckSudo()
	// Show banner
	showBanner()
	// Initialize global config (db, logger, etc.)
	// From now on it will be accessible as utils.Config
	utils.InitConfig()
}

// ---------------------------------------------------------------------------------------
// MAIN
// ---------------------------------------------------------------------------------------
func main() {
	// Setup core
	initCore()

	// Start CLI
	p := prompt.New(
		cli.Executor,
		cli.Completer,
		prompt.OptionTitle("securabox: Interactive Network Scanner"),
		prompt.OptionPrefix("[securabox] > "),
		prompt.OptionInputTextColor(prompt.White),
	)
	p.Run()
}
