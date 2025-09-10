package daysteps

import (
	"fmt"
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
		return 0, 0, fmt.Errorf("неправильное количество данных")
	}

	stepsStr := strings.TrimSpace(dataStorage[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("неправильное количество шагов")
	}

	durationStr := strings.TrimSpace(dataStorage[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
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
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories := spentcalories.WalkingSpentCalories(weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, calories)

}
