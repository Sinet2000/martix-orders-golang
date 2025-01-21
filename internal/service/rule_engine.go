package service

import (
	"github.com/Sinet2000/Martix-Orders-Go/internal/entity"
	"time"
)

type RuleEngine struct {
	ruleRepo RuleRepository
	cache    RuleCache
}

type RuleRepository interface {
	GetActiveRules(ctx context.Context) ([]DiscountRule, error)
}

type RuleCache interface {
	Get(key string) ([]DiscountRule, error)
	Set(key string, rules []DiscountRule, expiration time.Duration)
}

func (re *RuleEngine) GetApplicableRules(order *entity.Order) []DiscountRule {
	// Try getting rules from cache
	rules, err := re.cache.Get("active_rules")
	if err != nil {
		// Fetch from repository if not in cache
		rules, _ = re.ruleRepo.GetActiveRules(context.Background())
		re.cache.Set("active_rules", rules, 1*time.Hour)
	}

	var applicableRules []DiscountRule
	currentTime := time.Now()

	for _, rule := range rules {
		if isRuleApplicable(rule, order, currentTime) {
			applicableRules = append(applicableRules, rule)
		}
	}

	return applicableRules
}

func isRuleApplicable(rule DiscountRule, order *entity.Order, currentTime time.Time) bool {
	// Check time validity
	if currentTime.Before(rule.StartDate) || currentTime.After(rule.EndDate) {
		return false
	}

	// Check minimum order amount
	if order.TotalAmount < rule.MinOrderAmount {
		return false
	}

	// Check user type eligibility
	if rule.UserType != "" && order.User.Type != rule.UserType {
		return false
	}

	// Check item categories if rule is category-specific
	if len(rule.ItemCategories) > 0 {
		return hasMatchingCategories(order.Items, rule.ItemCategories)
	}

	// First purchase discount
	if rule.Type == "first_purchase" && !isFirstPurchase(order.User.ID) {
		return false
	}

	// Bulk order discount
	if rule.Type == "bulk_order" && !isBulkOrder(order.Items) {
		return false
	}

	return true
}

func hasMatchingCategories(items []entity.OrderItem, categories []string) bool {
	categoryMap := make(map[string]bool)
	for _, category := range categories {
		categoryMap[category] = true
	}

	for _, item := range items {
		if categoryMap[item.Category] {
			return true
		}
	}
	return false
}

func isFirstPurchase(userID string) bool {
	// Implementation to check if this is user's first purchase
	return false
}

func isBulkOrder(items []entity.OrderItem) bool {
	totalQuantity := 0
	for _, item := range items {
		totalQuantity += item.Quantity
	}
	return totalQuantity >= 10
}
