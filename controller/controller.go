package controller

import (
	"MTree/model"
	"net/http"
	"strings"
)

func GetTreeController() (model.OTree, int) {
	code := http.StatusOK

	resp, err := model.GetTreeModel()
	struct_arr := strings.Split(resp.Struct, ",")
	output := model.OTree{
		Id:     resp.Id,
		Root:   resp.Root,
		Struct: struct_arr,
	}

	if err != nil {
		code = http.StatusInternalServerError
	}

	return output, code
}
