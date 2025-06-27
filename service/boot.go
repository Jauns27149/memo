package service

var (
	MemoService *Data
	PlanService *Plan
	EatService  *Eat
)

func Boot() {
	MemoService = newData()
	PlanService = NewPlan()
	EatService = NewEat()
}
