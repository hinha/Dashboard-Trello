package app

type DashboardContent struct {
	User   interface{}
	Errors map[string]string
	Any    map[string]string
	Page   map[string]int
	Token  string
	Data   map[string]interface{}
}

type MenuType int

const (
	HomeMenu MenuType = iota
	HomeMenuInbox
	HpmeMenuAttendece
	HomeMenuActivity
	HomeMenuEmployee
	AnalyticsSummary
	AnalyticsResult
	TrelloMenu
	SettingDetails
	SettingUsers
)
