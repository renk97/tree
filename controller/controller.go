package controller

import (
	"MTree/model"
	"encoding/json"
	"log"
	"net/http"
)

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

func CreateHashTreeController(input model.HashIOTree) int {
	code := http.StatusOK

	err := model.CreateHashTreeModel(input)

	if err != nil {
		code = http.StatusInternalServerError
	}

	return code
}

func GetTreeController(leaf_type string, id int, root string) (out_arr []interface{}, code int) {
	code = http.StatusOK

	resp, err := model.GetTreeModel(leaf_type, id, root)
	if err != nil {
		code = http.StatusInternalServerError
	}

	switch leaf_type {
	case "hash":
		for _, tree := range resp {
			var output model.HashIOTree

			err := json.Unmarshal(tree.Leaves, &output)
			output.Id = tree.Id
			if err != nil {
				log.Println(err)
				return
			}
			out_arr = append(out_arr, output)
		}
	default:
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
	}

	return out_arr, code
}
