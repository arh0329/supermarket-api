package main

import (
	"github.com/arh0329/supermarket-api/pkg/api"
	logger "github.com/arh0329/supermarket-api/pkg/logging"
)

func main() {
	r := api.CreateRouter()

	err := r.Run(":8000")
	if err != nil {
		logger.Log().WithError(err).Fatal("Error occurred running http server")
	}
}
