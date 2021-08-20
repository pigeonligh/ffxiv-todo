package algorithm

import (
	"container/list"
	"fmt"
	"math"

	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type workItem struct {
	id     int
	amount int
	stage  int
	links  int
}

type Workshop struct {
	manager     *cache.Manager
	items       map[int]*workItem
	products    map[int]int
	inventories map[int]int

	workqueue *list.List
}

type Report struct {
	ItemName      map[int]string        `json:"items"`
	ItemGathering map[int][]string      `json:"gatherings"`
	ItemRecipe    map[int]*types.Recipe `json:"recipes"`
	Stages        map[int][]int         `json:"stages"`
}

func NewWorkshop(m *cache.Manager) *Workshop {
	return &Workshop{
		manager:     m,
		items:       make(map[int]*workItem),
		products:    make(map[int]int),
		inventories: make(map[int]int),
	}
}

func (w *Workshop) AddProduct(id, number int) {
	w.products[id] = w.products[id] + number
}

func (w *Workshop) AddInventory(id, number int) {
	w.inventories[id] = w.inventories[id] + number
}

func (w *Workshop) Calculate() error {
	for id, number := range w.products {
		w.fill(id, 0)
		w.items[id].amount += number
	}
	return w.calcAmount()
}

func (w *Workshop) fill(id, addlink int) {
	if item, found := w.items[id]; found {
		item.links += addlink
	} else {
		item := &workItem{
			id:     id,
			amount: 0,
			stage:  1,
			links:  addlink,
		}
		recipe := w.manager.Item.GetByID(id).Recipe
		if recipe.Amount > 0 {
			for ing := range recipe.Ingredients {
				w.fill(ing, 1)
				item.stage = maxInt(item.stage, w.items[ing].stage+1)
			}
		}
		w.items[id] = item
	}
}

func (w *Workshop) calcAmount() error {
	w.workqueue = list.New()
	for id := range w.products {
		if w.items[id].links == 0 {
			w.workqueue.PushBack(id)
		}
	}
	for w.workqueue.Len() > 0 {
		elem := w.workqueue.Front()
		id := elem.Value.(int)
		nowitem := w.items[id]
		nowitem.amount = maxInt(nowitem.amount - w.inventories[id])
		recipe := w.manager.Item.GetByID(id).Recipe
		if recipe.Amount > 0 {
			worktimes := int(math.Ceil(float64(nowitem.amount) / float64(recipe.Amount)))
			for ing, amount := range recipe.Ingredients {
				ingitem := w.items[ing]
				ingitem.amount += worktimes * amount
				ingitem.links--
				if ingitem.links == 0 {
					w.workqueue.PushBack(ingitem.id)
				}
			}
		}
		w.workqueue.Remove(elem)
	}

	for _, item := range w.items {
		if item.links > 0 {
			return fmt.Errorf("error calculate")
		}
	}
	return nil
}

func (w *Workshop) PrintResults() {
	stages := make(map[int][]*workItem)
	maxStage := 0
	for _, item := range w.items {
		if _, found := stages[item.stage]; !found {
			stages[item.stage] = make([]*workItem, 0)
		}
		stages[item.stage] = append(stages[item.stage], item)
		maxStage = maxInt(maxStage, item.stage)
	}
	for i := 1; i <= maxStage; i++ {
		fmt.Printf("stage %d:\n", i)
		for _, item := range stages[i] {
			realItem := w.manager.Item.GetByID(item.id)
			recipe := realItem.Recipe
			gathering := ""
			for place := range realItem.Gathering {
				placeName := w.manager.Place.GetByID(place.Place)
				gathering += fmt.Sprintf(" %s(%s-%d)",
					placeName.Name,
					types.GatheringType[place.Type],
					place.Level,
				)
			}
			fmt.Printf("  %s * %d   %s\n", realItem.Name, item.amount, gathering)
			if recipe.Amount > 0 {
				// worktimes := int(math.Ceil(float64(item.amount) / float64(recipe.Amount)))
				fmt.Printf("    %d <- ", recipe.Amount)
				firstEmpty := ""
				for ing, amount := range recipe.Ingredients {
					ingitem := w.manager.Item.GetByID(ing)
					fmt.Printf("%s%s * %d", firstEmpty, ingitem.Name, amount)
					firstEmpty = " + "
				}
				fmt.Println()
			}
		}
	}
}

func (w *Workshop) Report() *Report {
	report := &Report{
		ItemName:      make(map[int]string),
		ItemGathering: make(map[int][]string),
		ItemRecipe:    make(map[int]*types.Recipe),
		Stages:        make(map[int][]int),
	}

	stages := make(map[int][]*workItem)
	maxStage := 0
	for _, item := range w.items {
		if _, found := stages[item.stage]; !found {
			stages[item.stage] = make([]*workItem, 0)
		}
		stages[item.stage] = append(stages[item.stage], item)
		maxStage = maxInt(maxStage, item.stage)
	}

	for i := 1; i <= maxStage; i++ {
		stageItems := make([]int, 0)
		for _, item := range stages[i] {
			realItem := w.manager.Item.GetByID(item.id)
			recipe := realItem.Recipe
			gathering := make([]string, 0)
			for place := range realItem.Gathering {
				placeName := w.manager.Place.GetByID(place.Place)
				gathering = append(gathering, placeName.Name)
			}

			report.ItemName[item.id] = realItem.Name
			report.ItemRecipe[item.id] = recipe
			report.ItemGathering[item.id] = gathering
			stageItems = append(stageItems, item.id)
		}
		report.Stages[i] = stageItems
	}

	return report
}
