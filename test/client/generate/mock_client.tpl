// Code generated by generate_mock_client.go. DO NOT EDIT.
package client

import (
	"testing"

	metalgo "github.com/metal-stack/metal-go"

{{ range $svc := . -}}
	"github.com/metal-stack/metal-go/api/client/{{ $svc.Name | lower | replace "switch" "switch_" }}"
	{{ $svc.Name | lower }}mocks "github.com/metal-stack/metal-go/test/mocks/{{ $svc.Name | lower | replace "switch" "switch_" }}"
{{ end }}

	"github.com/stretchr/testify/mock"
)

type MetalMockFns struct {
{{ range $svc := . -}}
	{{ $svc.Name }}    func(mock *mock.Mock)
{{ end }}
}

type MetalMockClient struct {
{{ range $svc := . -}}
	{{ $svc.Name | lower }}   *{{ $svc.Name | lower }}mocks.ClientService
{{ end }}
}

func NewMetalMockClient(t *testing.T, mockFns *MetalMockFns) (*MetalMockClient, metalgo.Client) {
	client := &MetalMockClient{
{{ range $svc := . -}}
		{{ $svc.Name | lower }}:   {{ $svc.Name | lower }}mocks.NewClientService(t),
{{ end }}
	}

	if mockFns == nil {
		return client, client
	}

{{ range $svc := . -}}
	if mockFns.{{ $svc.Name }} != nil {
		mockFns.{{ $svc.Name }}(&client.{{ $svc.Name | lower }}.Mock)
	}
{{ end }}

	return client, client
}

{{ range $svc := . -}}
func (c *MetalMockClient) {{ $svc.Name }}() {{ $svc.Name | lower | replace "switch" "switch_" }}.ClientService {
	return c.{{ $svc.Name | lower }}
}

{{ end }}
