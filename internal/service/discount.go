package service

import (
	"github.com/Sinet2000/Martix-Orders-Go/internal/entity"
	"time"
)

type DiscountService struct {
	ruleEngine RuleEngine
}

type DiscountRule struct {
	Type           string // percentage, fixed
	Amount         float64
	MinOrderAmount float64
	StartDate      time.Time
	EndDate        time.Time
	UserType       string // new, regular, premium
	ItemCategories []string
}

func (s *DiscountService) CalculateDiscount(order *entity.Order) float64 {
	var totalDiscount float64

	// Get applicable rules
	rules := s.ruleEngine.GetApplicableRules(order)

	for _, rule := range rules {
		switch rule.Type {
		case "percentage":
			discount := order.TotalAmount * (rule.Amount / 100)
			totalDiscount += discount

		case "fixed":
			if order.TotalAmount >= rule.MinOrderAmount {
				totalDiscount += rule.Amount
			}
		}
	}

	// Apply seasonal discounts
	if isHolidaySeason() {
		totalDiscount += s.calculateSeasonalDiscount(order)
	}

	// Apply loyalty discount
	if order.User.Type == "premium" {
		totalDiscount += s.calculateLoyaltyDiscount(order)
	}

	// Ensure discount doesn't exceed maximum allowed
	if totalDiscount > order.TotalAmount*0.7 {
		totalDiscount = order.TotalAmount * 0.7
	}

	return totalDiscount
}

func (s *DiscountService) calculateSeasonalDiscount(order *entity.Order) float64 {
	// Implement seasonal discount logic
	return 0
}

func (s *DiscountService) calculateLoyaltyDiscount(order *entity.Order) float64 {
	// Implement loyalty discount logic
	return 0
}

func isHolidaySeason() bool {
	// Implement holiday season check
	return false
}
