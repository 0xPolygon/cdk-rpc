package rpc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponse_MarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name         string
		request      Request
		responseData []byte
		err          *ErrorObject
	}{
		{
			name: "success response without error",
			request: Request{
				JSONRPC: "2.0",
				ID:      float64(1),
				Method:  "test_method",
			},
			responseData: json.RawMessage(`{"key":"value"}`),
			err:          nil,
		},
		{
			name: "error response with error",
			request: Request{
				JSONRPC: "2.0",
				ID:      float64(1),
				Method:  "test_method",
			},
			responseData: nil,
			err: &ErrorObject{
				Code:    123,
				Message: "test error",
				Data:    nil,
			},
		},
		{
			name: "error response with error and data",
			request: Request{
				JSONRPC: "2.0",
				ID:      float64(1),
				Method:  "test_method",
			},
			responseData: nil,
			err: &ErrorObject{
				Code:    123,
				Message: "test error",
				Data:    "error data",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rpcErr *RPCError
			if tt.err != nil {
				rpcErr = NewRPCError(tt.err.Code, tt.err.Message, tt.err.Data)
			}

			resp := NewResponse(tt.request, tt.responseData, rpcErr)
			data, err := resp.Bytes()
			require.NoError(t, err)

			var unmarshaledResp Response
			err = json.Unmarshal(data, &unmarshaledResp)
			require.NoError(t, err)
			require.Equal(t, resp, unmarshaledResp)
		})
	}
}
