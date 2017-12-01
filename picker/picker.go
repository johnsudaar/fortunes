package picker

import (
	"bufio"
	"math/rand"
	"os"
)

type Picker struct {
	Fortunes []string
}

func LoadPicker(path string) (Picker, error) {
	result := Picker{}
	file, err := os.Open(path)
	if err != nil {
		return result, nil
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	cont := true
	for cont {
		line, err := reader.ReadString('\n')
		if err != nil {
			cont = false
		}
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}

	result.Fortunes = lines
	return result, nil
}

func (p Picker) Pick() string {
	choice := rand.Int() % len(p.Fortunes)
	return p.Fortunes[choice]
}
