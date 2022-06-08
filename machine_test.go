package metalgo

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emicklei/go-restful/v3"
	"github.com/metal-stack/metal-go/client/machine"
	"github.com/metal-stack/metal-go/models"
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

	server, client, _ := startServerAndGetDriver()
	defer func() {
		_ = server.Close()
	}()

	// when
	resp, err := client.Machine().AllocateMachine(machine.NewAllocateMachineParams().WithBody(&models.V1MachineAllocateRequest{
		UUID: machineID,
	}), nil)

	// then
	require.Nil(t, err)
	require.NotNil(t, resp)
	require.Equal(t, metalMachine, resp.Payload)
}

func TestMachineCreateDriver(t *testing.T) {
	// given
	machineID := "machineID"
	metalMachine := &models.V1MachineResponse{
		ID: &machineID,
	}
	ws.Routes()[0].Function = func(req *restful.Request, resp *restful.Response) {
		_ = resp.WriteEntity(metalMachine)
	}

	server, _, driver := startServerAndGetDriver()
	defer func() {
		_ = server.Close()
	}()

	mcr := &MachineCreateRequest{}

	// when
	resp, err := driver.MachineCreate(mcr)

	// then
	require.Nil(t, err)
	require.NotNil(t, resp)
	require.Equal(t, metalMachine, resp.Machine)
}

func startServerAndGetDriver() (*http.Server, Client, *Driver) {
	//nolint:gosec
	listener, _ := net.Listen("tcp", ":0")

	server := &http.Server{}

	go func() {
		_ = server.Serve(listener)
	}()

	time.Sleep(10 * time.Millisecond)

	port := listener.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("http://localhost:%d", port)

	client, driver, _ := NewDriver(addr, "", "")

	return server, client, driver
}

func Test_translateNetworks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		request             MachineCreateRequest
		expectedAutoAcquire bool
		name                string
	}{
		{
			name:                "given false expected false",
			expectedAutoAcquire: false,
			request: MachineCreateRequest{
				Networks: []MachineAllocationNetwork{
					{NetworkID: "network", Autoacquire: false},
				},
			},
		},
		{
			name:                "given true expected true",
			expectedAutoAcquire: true,
			request: MachineCreateRequest{
				Networks: []MachineAllocationNetwork{
					{NetworkID: "network", Autoacquire: true},
				},
			},
		},
		{
			name:                "network ids do not differ",
			expectedAutoAcquire: true,
			request: MachineCreateRequest{
				Networks: []MachineAllocationNetwork{
					{NetworkID: "network1", Autoacquire: true},
					{NetworkID: "network2", Autoacquire: true},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		translatedNetworks := test.request.translateNetworks()
		assert.NotNil(translatedNetworks)
		assert.Len(translatedNetworks, len(test.request.Networks))

		// deep equals compare check to verify the network translation is valid
		comp := func() bool {
			for _, network := range test.request.Networks {
				for _, translatedNetwork := range translatedNetworks {
					if network.NetworkID == *translatedNetwork.Networkid {
						// verify conversion of boolean pointer to boolean works as expected
						return network.Autoacquire == *translatedNetwork.Autoacquire
					}
				}
			}
			return false
		}
		assert.Condition(comp, "translated networks: %+v do not equal the origin: %+v")
	}
}
