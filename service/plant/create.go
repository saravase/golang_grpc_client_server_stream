package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_client_server_stream/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePlant used to implement client streaming
// To insert the new plant data into the datastore
func (server *PlantServer) CreatePlant(stream pb.PlantService_CreatePlantServer) error {

	var plants []string

	for {
		plant, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Created plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		server.store.Save(plant)
		plants = append(plants, plant.Id)
	}

}
