package metalgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/models"
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
			Returns(http.StatusNotFound, "No free machine for allocation found", models.HttperrorsHTTPErrorResponse{}).
			Returns(http.StatusUnprocessableEntity, "Unprocessable Entity", models.HttperrorsHTTPErrorResponse{}))

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

	server, driver := startServerAndGetDriver()
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

func TestRetryAllocateMachine(t *testing.T) {
	// given
	errMsg := "Faked error"
	attemptTracker := &sync.Map{}
	machineID := "machineID"
	metalMachine := &models.V1MachineResponse{
		ID: &machineID,
	}
	ws.Routes()[0].Function = func(req *restful.Request, resp *restful.Response) {
		reqBody, _ := ioutil.ReadAll(req.Request.Body)
		allocMachineRequest := &models.V1MachineAllocateRequest{}
		_ = json.Unmarshal(reqBody, allocMachineRequest)

		value, _ := attemptTracker.LoadOrStore(allocMachineRequest.UserData, 0)
		attempt := value.(int)
		attempt++
		attemptTracker.Store(allocMachineRequest.UserData, attempt)
		if attempt < 4 {
			err := models.HttperrorsHTTPErrorResponse{
				Message: &errMsg,
			}
			errResp, _ := json.Marshal(err)
			_ = resp.WriteErrorString(http.StatusNotFound, string(errResp))

			return
		}

		_ = resp.WriteEntity(metalMachine)
	}

	server, driver := startServerAndGetDriver()
	defer func() {
		_ = server.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(7)

	// Verify failure with retry config that will be applied due to a matching status code but having not enough retries configured
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "1",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)
		allocMachine.WithContext(newRetryContext(2, 5*time.Millisecond, func(sc int) bool {
			return sc == http.StatusNotFound
		}))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.NotNil(t, err)
		require.Contains(t, err.Error(), errMsg)
		require.Nil(t, resp)
	}()

	// Verify success with retry config that will be applied due to a matching status code
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "2",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)
		allocMachine.SetContext(newRetryContext(3, 4*time.Millisecond, func(sc int) bool {
			return sc == http.StatusNotFound
		}))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, metalMachine, resp.Payload)
	}()

	// Verify success with retry config that will always be applied
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "3",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)
		allocMachine.WithContext(newRetryContext(3, 3*time.Millisecond))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, metalMachine, resp.Payload)
	}()

	// Verify success with retry config that will be applied due to a matching status code
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "4",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)
		allocMachine.SetContext(newRetryContext(3, 2*time.Millisecond, func(sc int) bool {
			return sc == http.StatusBadRequest || sc == http.StatusNotFound
		}))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, metalMachine, resp.Payload)
	}()

	// Verify failure with retry config that will not be applied due to a non-matching status code
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "5",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)
		allocMachine.WithContext(newRetryContext(3, time.Millisecond, func(sc int) bool {
			return sc == http.StatusBadRequest
		}))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.NotNil(t, err)
		require.Contains(t, err.Error(), errMsg)
		require.Nil(t, resp)
	}()

	// Verify failure due to a lacking retry config
	go func() {
		defer wg.Done()

		// given
		allocMachineRequest := &models.V1MachineAllocateRequest{
			UserData: "6",
		}
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.SetBody(allocMachineRequest)

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.NotNil(t, err)
		require.Contains(t, err.Error(), errMsg)
		require.Nil(t, resp)
	}()

	// Verify success with nil body
	go func() {
		defer wg.Done()

		// given
		allocMachine := machine.NewAllocateMachineParams()
		allocMachine.WithContext(newRetryContext(3, time.Millisecond, func(sc int) bool {
			return sc == http.StatusNotFound
		}))

		// when
		resp, err := driver.machine.AllocateMachine(allocMachine, driver.auth)

		// then
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, metalMachine, resp.Payload)
	}()

	wg.Wait()
}

func startServerAndGetDriver() (*http.Server, *Driver) {
	listener, _ := net.Listen("tcp", ":0")

	server := &http.Server{}

	go func() {
		_ = server.Serve(listener)
	}()

	time.Sleep(10 * time.Millisecond)

	port := listener.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("http://localhost:%d", port)

	driver, _ := NewDriver(addr, "", "")

	return server, driver
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
