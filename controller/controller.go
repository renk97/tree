package controller

import (
	"MTree/model"
	"encoding/json"
	"log"
	"net/http"
)

func GetTreeController(id int, root string) (out_arr []model.IOTree, code int) {
	code = http.StatusOK

	resp, err := model.GetTreeModel(id, root)

	for _, tree := range resp {
		var output model.IOTree

		err := json.Unmarshal(tree.Leaves, &output)
		output.Id = tree.Id
		if err != nil {
			log.Println(err)
			return
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
