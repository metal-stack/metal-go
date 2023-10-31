package metalgo

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/stretchr/testify/require"
)

var ws *restful.WebService

func init() {
	ws = new(restful.WebService)
	ws.
		Path("/v1/machine").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Route(ws.POST("/allocate").To(func(request *restful.Request, response *restful.Response) {}).
			Reads(models.V1MachineAllocateRequest{}).
			Returns(http.StatusOK, "OK", models.V1MachineResponse{}).
			Returns(http.StatusNotFound, "No free machine for allocation found", httperrors.HTTPErrorResponse{}).
			Returns(http.StatusUnprocessableEntity, "Unprocessable Entity", httperrors.HTTPErrorResponse{}))

	restful.DefaultContainer.Add(ws)
}

func TestMachineCreate(t *testing.T) {
	// given
	machineID := "machineID"
	metalMachine := &models.V1MachineResponse{
		ID: &machineID,
	}
	ws.Routes()[0].Function = func(req *restful.Request, resp *restful.Response) {
		_ = resp.WriteEntity(metalMachine)
	}

	server, client := startServerAndGetDriver(t)
	defer func() {
		_ = server.Close()
	}()

	// when
	resp, err := client.Machine().AllocateMachine(machine.NewAllocateMachineParams().WithBody(&models.V1MachineAllocateRequest{
		UUID: machineID,
	}), nil)

	// then
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, metalMachine, resp.Payload)
}

func startServerAndGetDriver(t *testing.T) (*http.Server, Client) {
	//nolint:gosec
	listener, _ := net.Listen("tcp", ":0")

	server := &http.Server{ReadHeaderTimeout: time.Minute}

	go func() {
		_ = server.Serve(listener)
	}()

	time.Sleep(10 * time.Millisecond)

	port := listener.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("http://localhost:%d", port)

	client, err := NewDriver(addr, "", "")
	require.NoError(t, err)

	return server, client
}
