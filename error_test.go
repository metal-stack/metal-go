package metalgo

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/models"
	"github.com/stretchr/testify/assert"
)

func TestAttemptGenericError(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		priorErrorRegex string
		newErrorString  string
	}{
		{
			name: "swagger error gets printed nicely",
			err: &machine.FindMachinesDefault{
				Payload: &models.HttperrorsHTTPErrorResponse{
					Message:    strPtr("some error message"),
					Statuscode: int32Ptr(500),
				},
			},
			priorErrorRegex: regexp.QuoteMeta("[POST /v1/machine/find][0] findMachines default  &{Message:") + ".* Statuscode:.*" + regexp.QuoteMeta("}"),
			newErrorString:  "some error message (500)",
		},
		{
			name:            "non swagger error gets passed through",
			err:             fmt.Errorf("non-swagger-error"),
			priorErrorRegex: "non-swagger-error",
			newErrorString:  "non-swagger-error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := regexp.MustCompile(tt.priorErrorRegex)
			assert.Regexp(t, r, tt.err.Error())

			err := AttemptGenericError(tt.err)
			assert.Equal(t, err.Error(), tt.newErrorString)
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
