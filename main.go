package main

import (
	"fmt"
	"strconv"

	"github.com/tadvi/winc"
)

func CheckResult(m map[string]string) (Winner string) {
	Winner = "NoWinner"
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
		if key == "1" {
			if m["2"] == value && m["3"] == value {
				Winner = value
				break
			}
		}

		if key == "4" {
			if m["5"] == value && m["6"] == value {
				Winner = value
				break
			}
		}

		if key == "7" {
			if m["8"] == value && m["9"] == value {
				Winner = value
				break
			}
		}

		if key == "1" {
			if m["4"] == value && m["7"] == value {
				Winner = value
				break
			}
		}

		if key == "2" {
			if m["5"] == value && m["8"] == value {
				Winner = value
				break
			}
		}

		if key == "3" {
			if m["6"] == value && m["9"] == value {
				Winner = value
				break
			}
		}

		if key == "7" {
			if m["5"] == value && m["3"] == value {
				Winner = value
				break
			}
		}

		if key == "1" {
			if m["5"] == value && m["9"] == value {
				Winner = value
				break
			}
		}
	}

	return
}

func CreateAll() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300) // (width, height)
	mainWindow.SetText("Morpion")

	var currentTurn = "X"

	var Text = "Tour: " + currentTurn

	var BasePosX = 50
	var BasePosY = 50

	txt := winc.NewLabel(mainWindow)

	font := winc.NewFont("Aharoni", 14, winc.FontBold)
	txt.SetFont(font)
	txt.SetText(Text)
	txt.SetPos(2, 2)
	txt.SetSize(150, 20)

	resetButton := winc.NewPushButton(mainWindow)
	resetButton.SetFont(font)
	resetButton.SetText("Reset")
	resetButton.SetPos(200, 2)
	resetButton.SetSize(150, 20)

	var numCase = 0
	var numCaseCoche = 0

	var m map[string]string

	m = make(map[string]string)

	resetButton.OnClick().Bind(func(e *winc.Event) {
		mainWindow.Close()

		CreateAll()
	})

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			numCase = numCase + 1
			newButton := winc.NewPushButton(mainWindow)
			newButton.SetSize(50, 50)
			newButton.SetPos(BasePosX+x*50, BasePosY+y*50)
			name := strconv.Itoa(numCase)
			newButton.SetText(name)

			newButton.OnClick().Bind(func(e *winc.Event) {
				if m[name] == "" {
					numCaseCoche++
					newButton.SetText(currentTurn)
					m[name] = currentTurn

					if currentTurn == "X" {
						currentTurn = "O"
					} else {
						currentTurn = "X"
					}

					Text = "Tour: " + currentTurn
					txt.SetText(Text)
					Winner := CheckResult(m)
					fmt.Println(Winner)

					if Winner == "NoWinner" && numCaseCoche >= 9 {
						winText := "There was no winner!"
						mainWindow.Close()
						newWindow := winc.NewForm(nil)
						newWindow.SetSize(400, 300)

						txt := winc.NewLabel(newWindow)
						newWindow.SetText("Resultat")
						txt.SetText(winText)
						txt.SetPos(0, 0)
						txt.SetSize(400, 300)

						font := winc.NewFont("Aharoni", 100, winc.FontBold)
						txt.SetFont(font)

						newWindow.Center()
						newWindow.Show()

						restartButton := winc.NewPushButton(newWindow)
						restartButton.SetFont(font)
						restartButton.SetText("Restart")
						restartButton.SetPos(2, 200)
						restartButton.SetSize(150, 30)

						restartButton.OnClick().Bind(func(e *winc.Event) {
							newWindow.Close()
							CreateAll()
						})

						newWindow.OnClose().Bind(nwndOnClose)
					} else if Winner != "NoWinner" {

						winText := "The winner is " + Winner
						mainWindow.Close()
						newWindow := winc.NewForm(nil)
						newWindow.SetText("Resultat")
						newWindow.SetSize(400, 300)

						txt := winc.NewLabel(newWindow)
						txt.SetText(winText)
						txt.SetPos(0, 0)
						txt.SetSize(400, 300)

						font := winc.NewFont("Aharoni", 20, winc.FontBold)
						txt.SetFont(font)

						newWindow.Center()
						newWindow.Show()

						restartButton := winc.NewPushButton(newWindow)
						restartButton.SetFont(font)
						restartButton.SetText("Restart")
						restartButton.SetPos(2, 200)
						restartButton.SetSize(150, 30)

						restartButton.OnClick().Bind(func(e *winc.Event) {
							newWindow.Close()
							CreateAll()
						})

						newWindow.OnClose().Bind(nwndOnClose)
					}
				}
			})
		}
	}

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
}

func main() {
	CreateAll()

	winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func nwndOnClose(arg *winc.Event) {
	winc.Exit()
}
