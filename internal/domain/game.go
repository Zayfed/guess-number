package domain

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"module1/pkg"
	"os"
	"strconv"
	"time"
)

type Game struct {
	Date     time.Time `json:"date"`
	Attempts int       `json:"attempts"`
	Status   int       `json:"status"`
}

func NewGame() *Game {
	return &Game{Date: time.Now()}
}

func (g *Game) GameCycle() {
	attempts, maxNum := chooseDifficulty()

	answer := rand.IntN(maxNum + 1)
	fmt.Println(answer)

	guesses := make([]int, 0)

	for i := 1; i <= attempts; i++ {
		n := pkg.ReadNumber(Yellow + "Попытка №" + strconv.Itoa(i) + " введите новое число" + Reset)
		guesses = append(guesses, n)

		fmt.Printf("Предыдущие попытки: %v\n", guesses)

		if checkAttempt(answer, n) {
			g.Status = win
			g.Attempts = i
			g.endGame()
		}
	}

	g.Status = loose
	g.Attempts = attempts
	fmt.Printf(Red+"Вы не сравились, загаданное число – %v\n"+Reset, answer)
	g.endGame()
}

func (g *Game) endGame() {
	err := SaveGameInfo(*g, "results.json")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Game data saved successfully.")
	}

	n := pkg.ReadNumber("Хотите сыграть еще раз?\n" +
		"1 – начать новую игру\n" +
		"2 – завершить программу\n")
	switch n {
	case 1:
		g.GameCycle()
	case 2:
		os.Exit(0)
	default:
		fmt.Println("error: this option is not available")
		g.endGame()
	}
}

func chooseDifficulty() (int, int) {
	n := pkg.ReadNumber("Выберите сложность:\n" +
		"1 – easy\n" +
		"2 – medium\n" +
		"3 – hard\n")
	switch n {
	case 1:
		return attemptsEasy, maxEasy
	case 2:
		return attemptsMedium, maxMedium
	case 3:
		return attemptsHard, maxHard
	default:
		fmt.Println("error: this option is not available")
		return chooseDifficulty()
	}
}

func checkAttempt(answer, n int) bool {
	temp := findTemp(answer, n)

	if n < answer {
		fmt.Println(temp + ", cекретное число больше👆")
	} else if n > answer {
		fmt.Println(temp + ", cекретное число меньше👇")
	} else {
		fmt.Println(Green + "Поздравляю, вы нашли загаданное число!\n" + Reset)
		return true
	}
	return false
}

func findTemp(answer, n int) string {
	if math.Abs(float64(answer-n)) <= hot {
		return "🔥 Горячо"
	} else if math.Abs(float64(answer-n)) <= warm {
		return "🙂 Тепло"
	} else {
		return "❄️ Холодно"
	}
}

func SaveGameInfo(g Game, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	gameData, err := json.Marshal(g)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	gameData = append(gameData, '\n')

	_, err = file.Write(gameData)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
