package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataStorage := strings.Split(data, ",")
	if len(dataStorage) != 2 {
		log.Println("неправильное количество данных")
		return 0, 0, fmt.Errorf("неправильное количество данных")
	}

	steps, err := strconv.Atoi(dataStorage[0])
	if err != nil {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(dataStorage[1])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 || duration <= 0 {
		log.Println("Отрицательные значения шагов или длительности")
		return 0, 0, fmt.Errorf("Отрицательные значения шагов или длительности")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 || duration <= 0 {
		log.Println("Отрицательные значения шагов или длительности")
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return fmt.Sprint("Ошибка при расчете калорий: ", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)

}
