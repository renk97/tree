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
		struct_arr := strings.Split(tree.Struct, ",")
		output := model.IOTree{
			Id:     tree.Id,
			Root:   tree.Root,
			Struct: struct_arr,
		}
		out_arr = append(out_arr, output)
	}

	if err != nil {
		code = http.StatusInternalServerError
	}

	return out_arr, code
}

// func CreateTreeController(input model.IOTree) int {
// 	code := http.StatusOK

// 	struct_arr := strings.Split(resp.Struct, ",")
// 	output := model.IOTree{
// 		Id:     resp.Id,
// 		Root:   resp.Root,
// 		Struct: struct_arr,
// 	}
// 	resp, err := model.CreateTreeModel()

// 	if err != nil {
// 		code = http.StatusInternalServerError
// 	}

// 	return output, code
// }
