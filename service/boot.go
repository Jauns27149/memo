package service

var (
	MemoService *Memo
	PlanService *Plan
	EatService  *Eat
)

func Boot() {
	MemoService = newData()
	PlanService = NewPlan()
	EatService = NewEat()
	listenerPref()
}
