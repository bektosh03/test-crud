package ports

import (
	"github.com/bektosh03/test-crud/data-service/app"
	"github.com/bektosh03/test-crud/genprotos/datapb"
)

type GrpcServer struct {
	app *app.App
	datapb.UnimplementedDataServiceServer
}


