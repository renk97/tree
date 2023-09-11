package controller

import (
	"MTree/model"
	"net/http"
	"strings"
)

func GetTreeController(id int, root string) ([]model.IOTree, int) {
	var out_arr []model.IOTree
	code := http.StatusOK

	resp, err := model.GetTreeModel(id, root)

	for _, tree := range resp {
		Leaves_arr := strings.Split(tree.Leaves, ",")
		output := model.IOTree{
			Id:     tree.Id,
			Root:   tree.Root,
			Leaves: Leaves_arr,
		}
		out_arr = append(out_arr, output)
	}

	if err != nil {
		code = http.StatusInternalServerError
	}

	return out_arr, code
}

func CreateTreeController(input model.IOTree) int {
	code := http.StatusOK

	err := model.CreateTreeModel(input)

	if err != nil {
		code = http.StatusInternalServerError
	}

	return code
}

func UpdateLeafController(input model.IOTree) int {
	code := http.StatusOK

	err := model.UpdateLeafModel(input)

	if err != nil {
		code = http.StatusInternalServerError
	}

	return code
}

func DeleteTreeController(id int, root string) int {
	code := http.StatusOK

	err := model.DeleteTreeModel(id, root)

	if err != nil {
		code = http.StatusInternalServerError
	}

	return code
}
