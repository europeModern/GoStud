package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name string
	TrendLevel int
	Category string
	Views float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme
	
	for _, meme := range memes {
		if meme.Views > minViews {
			result = append(result, meme)
		}
	}
	
	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})
	
	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)
	
	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}
	
	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string
	
	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			result = append(result, meme.Name)
		}
	}
	
	return result
}

func main() {
	fmt.Println("ЗАДАЧА #2: Анализ Brainrot-контента")
	
	memes := []BrainrotMeme{
		{Name: "Skibidi Toilet", TrendLevel: 10, Category: "Skibidi", Views: 120.5},
		{Name: "Sigma Male", TrendLevel: 8, Category: "Sigma", Views: 75.3},
		{Name: "Mewing Challenge", TrendLevel: 7, Category: "Mewing", Views: 45.2},
		{Name: "Subo Bratik Dance", TrendLevel: 9, Category: "Subo Bratik", Views: 88.7},
		{Name: "TUNTUNTUNSAHUR Beat", TrendLevel: 6, Category: "TUNTUNTUNSAHUR", Views: 30.1},
		{Name: "Sigma Grindset", TrendLevel: 5, Category: "Sigma", Views: 55.8},
		{Name: "Random Brainrot", TrendLevel: 4, Category: "Other", Views: 20.3},
	}
	
	fmt.Println("\nИсходный набор мемов:")
	for _, meme := range memes {
		fmt.Printf("  %s | TrendLevel: %d | Category: %s | Views: %.1fM\n",
			meme.Name, meme.TrendLevel, meme.Category, meme.Views)
	}
	
	fmt.Println("\n--- FindTopTrending (minViews = 50.0) ---")
	topTrending := FindTopTrending(memes, 50.0)
	for _, meme := range topTrending {
		fmt.Printf("  %s | TrendLevel: %d | Views: %.1fM\n",
			meme.Name, meme.TrendLevel, meme.Views)
	}
	
	fmt.Println("\n--- CalculateCategoryImpact ---")
	categoryImpact := CalculateCategoryImpact(memes)
	for category, totalViews := range categoryImpact {
		fmt.Printf("  %s: %.1fM просмотров\n", category, totalViews)
	}
	
	fmt.Println("\n--- FilterByComplexCondition ---")
	filteredNames := FilterByComplexCondition(memes)
	fmt.Println("  Мемы, удовлетворяющие условию (TrendLevel >= 7 ИЛИ (Views > 50 И Category == 'Sigma')):")
	for _, name := range filteredNames {
		fmt.Printf("  - %s\n", name)
	}
}
