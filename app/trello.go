package app

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TrelloRepository interface {
	Store(in *TrelloUserCard) (*TrelloUserCard, error)
	FindCardCategory(id string) ([]CardCategory, error)
	CategoryByDate(id string) ([]CardGroupBy, error)
	ListTrelloUser() ([]*Trello, error)
	StoreUser(in TrelloAddMember) (TrelloAddMember, error)
	FindMemberID(id string) (*Trello, error)
}

type Trello struct {
	ID           uint64           `json:"id" gorm:"primary_key"`
	BoardName    string           `json:"board_name" gorm:"type:varchar(50);not null"`
	BoardID      string           `json:"board_id" gorm:"type:varchar(50);not null"`
	CardMemberID string           `json:"card_member_id" gorm:"type:varchar(120);null;index:card_member_id_id_k"`
	AccountID    string           `json:"-" gorm:"type:varchar(50);null"`
	CreatedAt    time.Time        `json:"created_at"`
	CardItems    []TrelloUserCard `json:"-" gorm:"ForeignKey:CardMemberID;references:CardMemberID"`
}

func (Trello) TableName() string {
	return "trello"
}

type TrelloUserCard struct {
	gorm.Model
	CardID                   string    `json:"card_id" gorm:"type:varchar(120);not null"`
	CardName                 string    `json:"card_name" gorm:"type:text;not null"`
	CardCategory             string    `json:"card_category" gorm:"type:varchar(120);not null"`
	CardVotes                int64     `json:"card_votes" gorm:"type:int(12);not null"`
	CardCheckItems           int64     `json:"card_check_items" gorm:"type:int(12);not null"`
	CardCheckLists           int64     `json:"card_check_lists" gorm:"type:int(12);not null"`
	CardCommentCount         int64     `json:"comment_count" gorm:"type:int(12);not null"`
	CardAttachmentsCount     int64     `json:"attachments_count" gorm:"type:int(12);not null"`
	CardCheckItemsComplete   int64     `json:"card_check_items_complete" gorm:"type:int(12);not null"`
	CardCheckItemsInComplete int64     `json:"card_check_items_incomplete" gorm:"type:int(12);not null"`
	CardUrl                  string    `json:"card_url" gorm:"type:text;not null"`
	CardMemberID             string    `json:"card_member_id" gorm:"type:varchar(120);not null"`
	CardMemberName           string    `json:"card_member_name" gorm:"type:text;not null"`
	CardMemberUsername       string    `json:"card_member_username" gorm:"type:text;not null"`
	CardCreatedAt            time.Time `json:"card_created_at" gorm:"not null;"`
}

func (TrelloUserCard) TableName() string {
	return "trello_user_card"
}

type CardCategory struct {
	Label string `json:"label"`
	Count int    `json:"count"`
}

type TrelloItemList struct {
	User       []*Accounts `json:"user"`
	TrelloUser []*Trello   `json:"trello_user"`
}

type Performance struct {
	CardCategory  []CardCategory `json:"card_category"`
	Product       interface{}    `json:"product"`
	Daily         interface{}    `json:"daily"`
	Task          interface{}    `json:"task"`
	TimelineGantt interface{}    `json:"timeline_gantt"`
	OnlineUsers   interface{}    `json:"online_users"`
}

type CardGroupBy struct {
	Category string
	Count    int
	Date     string
}

func (m *Performance) splitData(groupby []CardGroupBy) ([]string, []string, []int, []int, []int) {
	var labels []string
	var dates []string
	var lineVal1 []int // ON PROGRESS
	var lineVal2 []int // TO DO
	var lineVal3 []int // DONE
	for _, data := range groupby {
		labels = append(labels, data.Category)
		dates = append(dates, data.Date)

		if data.Category == string(CardTypePROGRESS) {
			lineVal1 = append(lineVal1, data.Count)
		} else {
			lineVal1 = append(lineVal1, 0)
		}
		if data.Category == string(CardTypeTODO) {
			lineVal2 = append(lineVal2, data.Count)
		} else {
			lineVal2 = append(lineVal2, 0)
		}
		if data.Category == string(CardTypeDONE) {
			lineVal3 = append(lineVal3, data.Count)
		} else {
			lineVal3 = append(lineVal3, 0)
		}
	}

	// remove duplicate label
	labels = m.removeDuplicateStr(labels)

	return labels, dates, lineVal1, lineVal2, lineVal3
}

func (m *Performance) LineChart(groupby []CardGroupBy) *charts.Line {

	labels, dates, lineVal1, lineVal2, lineVal3 := m.splitData(groupby)

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Data: labels,
			Show: true,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
			Show:    true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside", Start: 0, End: 100}, opts.DataZoom{Start: 0, End: 10}),
		charts.WithYAxisOpts(opts.YAxis{Type: "value"}),
		charts.WithXAxisOpts(opts.XAxis{Type: "category", Data: dates}),
	)

	itemProgress := make([]opts.LineData, 0)
	for _, val := range lineVal1 {
		itemProgress = append(itemProgress, opts.LineData{Value: val})
	}

	itemTodo := make([]opts.LineData, 0)
	for _, val := range lineVal2 {
		itemTodo = append(itemTodo, opts.LineData{Value: val})
	}
	itemDone := make([]opts.LineData, 0)
	for _, val := range lineVal3 {
		itemDone = append(itemDone, opts.LineData{Value: val})
	}

	line.SetXAxis(dates).
		AddSeries(string(CardTypeTODO), itemTodo).
		AddSeries(string(CardTypeDONE), itemDone).
		AddSeries(string(CardTypePROGRESS), itemProgress)

	return line
}

func (m *Performance) PieChart(groupby []CardGroupBy) *charts.Pie {

	duplicateFrequency := make(map[string]int)
	for _, dat := range groupby {
		_, exist := duplicateFrequency[dat.Category]
		if exist {
			duplicateFrequency[dat.Category] += dat.Count
		} else {
			// count index value
			duplicateFrequency[dat.Category] += dat.Count
		}
	}

	items := make([]opts.PieData, 0)
	for k, item := range duplicateFrequency {
		items = append(items, opts.PieData{Name: k, Value: item})

	}

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Trigger: "item", Show: true}),
	)

	pie.AddSeries("Task", items).
		SetSeriesOptions(
			charts.WithSunburstOpts(opts.SunburstChart{Animation: true, SelectedMode: true}),
			charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{b}: {c}"}),
			charts.WithEmphasisOpts(opts.Emphasis{ItemStyle: &opts.ItemStyle{}}))

	return pie
}

func (m *Performance) removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type TrelloAddMember struct {
	BoardName string            `json:"board_name"`
	MemberID  string            `json:"member_id"`
	UserID    string            `json:"user_id"`
	BoardID   string            `json:"board_id"`
	Errors    map[string]string `json:"errors"`
}

func (m *TrelloAddMember) Validate() bool {
	m.Errors = make(map[string]string)

	if strings.TrimSpace(m.BoardName) == "" {
		m.Errors["board_name"] = "Please enter a board name"
	} else if strings.TrimSpace(m.MemberID) == "" {
		m.Errors["member_id"] = "Please enter a member id"
	} else if strings.TrimSpace(m.UserID) == "" {
		m.Errors["user_id"] = "Please enter a user id"
	} else if strings.TrimSpace(m.BoardID) == "" {
		m.Errors["board_id"] = "Please enter a board id"
	}
	return len(m.Errors) == 0
}

type CardCategoryType string

const (
	CardTypeTODO     CardCategoryType = "TO DO"
	CardTypePROGRESS CardCategoryType = "ON PROGRESS"
	CardTypeDONE     CardCategoryType = "DONE"
	CardTypeReview   CardCategoryType = "TESTING"
)
