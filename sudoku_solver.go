package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	emptyCell string = "."
	boardSize        = 9
	squarSize        = 3
)

func inputBoardFromSite(driver selenium.WebDriver) [][]string {
	var testBoarNineXNine = [][]string{
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	for i := 0; i < 9*9; i++ {
		element, err := driver.FindElement(selenium.ByID, strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		text, err := element.Text()
		if err != nil {
			log.Fatal(err)
		}
		if text != "" {
			testBoarNineXNine[i/9][i%9] = text
		}
	}
	return testBoarNineXNine
}

func main() {

	service, err := selenium.NewChromeDriverService("chromedriver", 4444)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal(err)
	}
	err = driver.Get("https://absite.ru/sudoku/")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Print("0 - return\n1 - solve sudocu\n")
		comand := 0
		fmt.Scanf("%d\n", &comand)
		switch comand {
		case 0:
			driver.Quit()
			return
		case 1:
			testBoarNineXNine := inputBoardFromSite(driver)
			printBoard(testBoarNineXNine)
			if solveSudoku(testBoarNineXNine) {
				fmt.Println("Решил!")
				printBoard(testBoarNineXNine)
				inputBoardInSite(testBoarNineXNine)
			} else {
				fmt.Println("Не решается чета")
			}
		default:
			continue
		}
	}
}

func solveSudoku(board [][]string) bool {
	currentPositionX, currentPositionY, flag := searchEmpty(board)
	if !flag {
		return true
	}

	for num := 1; num < boardSize+1; num++ {
		currentNum := strconv.Itoa(num)

		if checkValidate(board, currentNum, currentPositionX, currentPositionY) {
			board[currentPositionX][currentPositionY] = currentNum
			if solveSudoku(board) {
				return true
			}
			board[currentPositionX][currentPositionY] = emptyCell
		}
	}

	return false
}

func checkValidate(board [][]string, num string, positionX int, positionY int) bool {

	//check cols
	for x := 0; x < boardSize; x++ {
		if board[x][positionY] == num && x != positionX {
			return false
		}
	}

	//check rows
	for y := 0; y < boardSize; y++ {
		if board[positionX][y] == num && y != positionY {
			return false
		}
	}

	//check box
	boxPositionX, boxPositionY := positionitionSquar(positionX, positionY)
	for x := boxPositionX; x < boxPositionX+squarSize; x++ {
		for y := boxPositionY; y < boxPositionY+squarSize; y++ {
			if board[x][y] == num && x != positionX && y != positionY {
				return false
			}
		}
	}
	return true
}

func positionitionSquar(positionX int, positionY int) (int, int) {
	return (positionX / squarSize) * squarSize, (positionY / squarSize) * squarSize
}

func searchEmpty(board [][]string) (int, int, bool) {
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			if board[x][y] == "." {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

func printBoard(board [][]string) {
	delimiter := strings.Repeat("-", boardSize*2+boardSize/6)
	fmt.Print("\n" + delimiter + "\n")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("|")
		for y := 0; y < boardSize; y++ {
			fmt.Printf(board[x][y])
			fmt.Printf("|")
		}
		fmt.Print("\n" + delimiter + "\n")
	}
}

var (
	mapping = map[string]int{
		"1": keybd_event.VK_1,
		"2": keybd_event.VK_2,
		"3": keybd_event.VK_3,
		"4": keybd_event.VK_4,
		"5": keybd_event.VK_5,
		"6": keybd_event.VK_6,
		"7": keybd_event.VK_7,
		"8": keybd_event.VK_8,
		"9": keybd_event.VK_9,
	}
)

func inputBoardInSite(board [][]string) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		time.Sleep(5 * time.Second)
	}

	for x := 0; x < boardSize; x++ {
		kb.SetKeys(keybd_event.VK_UP)
		kb.Press()
		time.Sleep(10 * time.Microsecond)
		kb.Release()
		kb.SetKeys(keybd_event.VK_LEFT)
		kb.Press()
		time.Sleep(10 * time.Microsecond)
		kb.Release()
	}

	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			kb.SetKeys(mapping[board[x][y]])
			kb.Press()
			kb.Release()
			kb.SetKeys(keybd_event.VK_RIGHT)
			kb.Press()
			time.Sleep(10 * time.Microsecond)
			kb.Release()
		}
	}
}
