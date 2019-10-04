package sushi

import (
	"errors"
	"strconv"
)

var (
	// ErrDoesNotExist is returned when there are problems with the request.
	ErrDoesNotExist = errors.New("Sushi with this ID does not exist")

	// ErrPageMoved is returned when a 301/302 is returned.
	ErrPageMoved = errors.New("Page Moved")
)

var rolls []Roll

func init() {

	rolls = append(rolls, Roll{ID: "1", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})
	rolls = append(rolls, Roll{ID: "2", Name: "California Roll", Ingredients: "Crab, Cucumber, Avocado, Nori, Rice"})
	rolls = append(rolls, Roll{ID: "3", Name: "Caterpillar Roll", Ingredients: "Tempura Shrimp, Cucumber, Avocado, Nori, Rice, Eel Sauce"})
}

// GetMenu provide the full menu
func GetMenu() []Roll {
	return rolls
}

// AddRoll adds a new roll to the menu
func AddRoll(newRoll Roll) Roll {
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	return newRoll
}

// UpdateRoll updates a given roll
func UpdateRoll(id string, changedRoll Roll) (Roll, error) {
	for i, item := range rolls {
		if item.ID == id {
			rolls = append(rolls[:i], rolls[i+1:]...)
			rolls = append(rolls, changedRoll)
			changedRoll.ID = id
			return changedRoll, nil
		}
	}
	return Roll{}, ErrDoesNotExist
}

// DeleteRoll deletes a role from the sushi menu
func DeleteRoll(id string) error {
	for i, item := range rolls {
		if item.ID == id {
			rolls = append(rolls[:i], rolls[i+1:]...)
			return nil
		}
	}
	return ErrDoesNotExist
}
