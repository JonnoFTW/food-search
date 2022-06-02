package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type Food struct {
	// All values should be percentage per 100g
	name     string
	category string
	calories float64
	fat      float64
	protein  float64
	carbs    float64
}

type Target struct {
	calories float64
	fat      float64
	protein  float64
	carbs    float64
}
type Portion struct {
	food    *Food
	serving float64 // amount of food in grams
}
type Meal struct {
	menu []Portion
}

func (m *Meal) totals() (float64, float64, float64, float64, float64) {
	totalGrams := 0.0
	totalCalories := 0.0
	totalFat := 0.0
	totalProtein := 0.0
	totalCarbs := 0.0
	for i := 0; i < len(m.menu); i++ {
		portionGrams := m.menu[i].serving
		totalGrams += portionGrams
		portionAs100Grams := portionGrams / 100
		totalCalories += m.menu[i].food.calories * portionAs100Grams
		totalFat += m.menu[i].food.fat * portionAs100Grams
		totalCarbs += m.menu[i].food.carbs * portionAs100Grams
		totalProtein += m.menu[i].food.protein * portionAs100Grams
	}
	return totalGrams, totalCalories, totalFat, totalProtein, totalCarbs
}

func distance(m *Meal, target *Target) float64 {
	totalGrams, totalCalories, totalFat, totalProtein, totalCarbs := m.totals()
	pctCalories := totalCalories / target.calories
	pctFat := totalFat / totalGrams
	pctCarbs := totalCarbs / totalGrams
	pctProtein := totalProtein / totalGrams

	sum := math.Pow(1-pctCalories, 2)
	sum += math.Pow(target.fat-pctFat, 2)
	sum += math.Pow(target.carbs-pctCarbs, 2)
	sum += math.Pow(target.protein-pctProtein, 2)
	return math.Sqrt(sum)

}
func (m *Meal) summary() {
	for _, portion := range m.menu {
		fmt.Printf("\t%.2fg: %s\n", portion.serving, portion.food.name)
	}
	totalGrams, totalCalories, totalFat, totalProtein, totalCarbs := m.totals()
	fmt.Println()
	fmt.Printf("\tTotal Grams:    %.2f\n", totalGrams)
	fmt.Printf("\tTotal Calories: %.2f\n", totalCalories)
	fmt.Printf("\tTotal Fat:      %.2f  (%.2f%%)\n", totalFat, totalFat/totalGrams*100)
	fmt.Printf("\tTotal Protein:  %.2f  (%.2f%%)\n", totalProtein, totalProtein/totalGrams*100)
	fmt.Printf("\tTotal Carbs:    %.2f  (%.2f%%)\n", totalCarbs, totalCarbs/totalGrams*100)

}
func (m *Meal) score(target *Target) float64 {
	return distance(m, target)
}
func removePortion(s []Portion, i int) []Portion {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func removeMeal(s []Meal, i int) []Meal {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

type IntRange struct {
	min, max int
}

func (ir *IntRange) NextRandom() int {
	return rand.Intn(ir.max-ir.min+1) + ir.min
}

// mutate copies a meal and returns its mutated offspring
func mutate(m Meal, foods []Food) Meal {
	// Do one of
	//  0. Remove random portion
	//  1. Add random portion
	//  2. Adjust portion size of a random food
	choice := rand.Intn(3)
	mLen := len(m.menu)
	if mLen == 1 && choice == 0 {
		choice++
	}
	switch choice {
	case 0:
		delPos := rand.Intn(mLen)
		m.menu = removePortion(m.menu, delPos)
	case 1:
		m.menu = append(m.menu, NewPortion(foods))
	case 3:
		adjustPos := rand.Intn(mLen)
		intRange := IntRange{-8, 8}
		m.menu[adjustPos].serving = m.menu[adjustPos].serving + float64(intRange.NextRandom())
	}
	return m
}
func NewPortion(foods []Food) Portion {
	randomFood := &foods[rand.Intn(len(foods))]
	return Portion{
		randomFood,
		rand.Float64()*31 + 1,
	}
}

func NewMeal(foods []Food) Meal {
	meal := Meal{}
	// add a random food in random amounts
	meal.menu = append(meal.menu, NewPortion(foods))
	return meal
}

func (f *Food) UnmarshalJSON(p []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	f.name = tmp[1].(string)
	f.fat = tmp[7].(float64)
	f.carbs = tmp[9].(float64)
	f.protein = tmp[6].(float64)
	f.calories = tmp[3].(float64) / 4.184 // convert kj to calories
	f.category = tmp[56].(string)
	return nil
}

func loadAllFoods() []Food {
	rand.Seed(time.Now().UnixNano())
	var foods []Food
	jsonFile, err := os.Open("foods.json")
	if err != nil {
		log.Fatal(err)
	}
	// permitted food categories
	//allowed
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []Food
	json.Unmarshal([]byte(byteValue), &result)
	ignoreCatsSlice := []string{
		"Coffee and coffee substitutes",
		"Tea",
		"Cordials",
		"Wines",
		"Spirits",
		"Beers",
		"Cider and perry",
		"Other alcoholic beverages",
		"Soft drinks, and flavoured mineral waters",
		"Reptiles",
		"Insects",
		"Amphibia",
	}

	ignoreCats := map[string]bool{}
	for _, s := range ignoreCatsSlice {
		ignoreCats[s] = true
	}
	for i := 0; i < len(result); i++ {
		if _, ok := ignoreCats[result[i].category]; !ok {
			foods = append(foods, result[i])
		}
	}
	return foods
}

func main() {
	// target goal of 2000 calories
	target := Target{2000, 0.1, 0.4, 0.5}
	// load in the foods to a slice
	foods := loadAllFoods()

	// create a population of meals
	var meals []Meal
	for i := 0; i < 256; i++ {
		newMeal := NewMeal(foods)
		meals = append(meals, newMeal)
	}
	// begin the tournament
	fmt.Println("Starting")
	round := 1
	for {
		// pick 2 random meals from the population
		mealAPos := rand.Intn(len(meals))
		mealBPos := rand.Intn(len(meals))

		mealA := meals[mealAPos]
		mealB := meals[mealBPos]

		if mealA.score(&target) > mealB.score(&target) {
			// kill off mealA, keep B and mutate
			meals = removeMeal(meals, mealAPos)
			meals = append(meals, mutate(mealB, foods))
		} else {
			// kill off mealB, keep A and mutate
			meals = removeMeal(meals, mealBPos)
			meals = append(meals, mutate(mealA, foods))
		}

		if round%1000000 == 0 {
			// calculate the current best meal
			var bestMeal Meal
			bestScore := 9999999.9
			for i := 0; i < len(meals); i++ {
				newScore := meals[i].score(&target)
				if newScore < bestScore {
					bestScore = newScore
					bestMeal = meals[i]
				}
			}
			fmt.Printf("Round %d: top score=%.3f meal is\n", round, bestScore)
			bestMeal.summary()
		}
		round += 1
	}
}
