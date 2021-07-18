package app

type DashboardContent struct {
	User   interface{}
	Errors map[string]string
	Any    map[string]string
	Page   map[string]int
}

type MenuType int

const (
	HomeMenu MenuType = iota
	TrelloMenu
)
