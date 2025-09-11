package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataStorage := strings.Split(data, ",")
	if len(dataStorage) != 3 {
		return 0, "", 0, fmt.Errorf("неправильное количество данных")
	}

	steps, err := strconv.Atoi(dataStorage[0])
	if err != nil {
		return 0, "", 0, err
	}

	acrivity := dataStorage[1]

	duration, err := time.ParseDuration(dataStorage[2])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 || duration <= 0 {
		return 0, "", 0, fmt.Errorf("Не может быть 0 или отрицательными")
	}
	return steps, acrivity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLeght := height * stepLengthCoefficient
	totalMeters := float64(steps) * stepLeght
	distance := totalMeters / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	durationHours := duration.Hours()
	if durationHours <= 0 {
		return 0
	}

	speed := dist / durationHours
	return speed

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64

	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		err := fmt.Errorf("неизвестный тип тренировки: %s", activity)
		log.Println(err)
		return "", err
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Данные не могут быть 0 или отрицательными")
	}
	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * speed * durationMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Данные не могут быть 0 или отрицательными")
	}
	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	baseCalories := (weight * speed * durationMinutes) / minInH
	calories := baseCalories * walkingCaloriesCoefficient
	return calories, nil
}
