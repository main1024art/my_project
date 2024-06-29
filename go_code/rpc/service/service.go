package service

import "go_code/rpc/service/data"

type Caluclator struct {
}

func (c Caluclator) Add(request *data.Caluclatorrequest, response *data.Caluclatorresponse) error {
	response.Result = request.A + request.B
	return nil
}
