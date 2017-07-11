package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unicode"

	"github.com/nsf/termbox-go"
	"github.com/thomasboyt/go-selecta/selecta"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "godesk"
	app.Usage = "godesk"
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) {
		goDesk()
	}

	app.Run(os.Args)
}

func goDesk() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "desk"
	cmdArgs := []string{"list"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	bytes := string(cmdOut)

	// bytes, _ := ioutil.ReadAll(os.Stdin)
	choices := strings.Split(string(bytes), "\n")

	// create a search
	initialSearch := "" //c.String("search")

	// set up termbox
	err2 := termbox.Init()
	if err2 != nil {
		panic(err2)
	}

	_, height := termbox.Size()
	search := selecta.BlankSearch(choices, initialSearch, height-1)

	termbox.SetInputMode(termbox.InputEsc)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	DrawApp(search)
	EventLoop(search)

	termbox.Close()

	binary, lookErr := exec.LookPath("desk")
	if lookErr != nil {
		panic(lookErr)
	}
	selectedChoice := strings.TrimSpace(search.SelectedChoice())
	args := []string{"desk", "go", selectedChoice}
	env := os.Environ()
	fmt.Println("go desk " + selectedChoice)
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func EventLoop(s *selecta.Search) {
	for !s.Done {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				termbox.Close()
				os.Exit(0)
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				s.Backspace()
			case termbox.KeyEnter:
				s.Done = true
			case termbox.KeyArrowDown, termbox.KeyCtrlN:
				s.Down()
			case termbox.KeyArrowUp, termbox.KeyCtrlP:
				s.Up()
			case termbox.KeyCtrlW:
				s.DeleteWord()
			default:
				char := rune(ev.Ch)
				if !unicode.IsControl(char) {
					s.AppendQuery(string(char))
				} else if ev.Key == termbox.KeySpace {
					s.AppendQuery(" ")
				}
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			DrawApp(s)
		}
	}
}

// Rendering

func DrawApp(s *selecta.Search) {
	matchCounter := "(" + (fmt.Sprintf("%v", len(s.Matches))) + ") > "
	WriteLine(0, matchCounter+s.Query, false)
	termbox.SetCursor(len(matchCounter)+len(s.Query), 0)

	for i, match := range s.Matches {
		if i >= s.VisibleChoices {
			break
		}
		choice := match.Value
		highlight := false
		if s.Index == i {
			highlight = true
		}
		WriteLine(i+1, choice, highlight)
	}

	termbox.Flush()
}

func WriteLine(row int, str string, highlight bool) {
	bgColor := termbox.ColorDefault
	textColor := termbox.ColorDefault
	width, _ := termbox.Size()
	if highlight {
		bgColor = termbox.ColorWhite
		textColor = termbox.ColorBlack
	}

	for col := 0; col < width; col++ {
		if col < len(str) {
			termbox.SetCell(col, row, rune(str[col]), textColor, bgColor)
		} else {
			termbox.SetCell(col, row, ' ', textColor, bgColor)
		}
	}
}
