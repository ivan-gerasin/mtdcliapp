package mtdcli

import (
	"fmt"
	"github.com/ivan-gerasin/mtdcore"
	"sort"
)

func renderAll(item *mtdCore.ToDoItem) {
	status := " "

	if item.Done {
		status = "✓"
	}
	fmt.Println(fmt.Sprintf("#%d - [%s] %s", item.Id, status, item.Summary))
}

func renderInProgress(item *mtdCore.ToDoItem) {
	if item.Done {
		return
	}
	fmt.Println(fmt.Sprintf("#%d - %s", item.Id, item.Summary))
}

func Render(list *mtdCore.ToDoGlobal, showComplete bool, sortByPriority bool) {
	if sortByPriority {
		sort.Slice(*list, func(i, j int) bool {
			return (*list)[i].Priority > (*list)[j].Priority
		})
	}
	for key := range *list {
		if showComplete {
			renderAll(&(*list)[key])
		} else {
			renderInProgress(&(*list)[key])
		}
	}
}
