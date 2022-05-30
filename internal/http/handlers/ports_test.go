package handlers

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"test/internal/service"
	"testing"
)

//go:embed testdata/full_ports.json
var fullJSON string

//go:embed testdata/array.json
var arrayJSON string

//go:embed testdata/invalid.json
var invalidJSON string

//go:embed testdata/empty_object.json
var emptyObjectJSON string

//go:embed testdata/invalid_port_object.json
var invalidPortObject string

//go:embed testdata/invalid_port_array.json
var invalidPortArray string

//go:embed testdata/one_port.json
var onePortJSON string

func TestPortsEndpoint(t *testing.T) {

	t.Run("full json 200", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(fullJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusNoContent {
			tt.Errorf("204 http status code expected, got: %d", res.Code)
		}
	})

	t.Run("json array 400", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		//mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(arrayJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusBadRequest {
			tt.Errorf("400 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "expected {, got [" {
			tt.Errorf("body `expected {, got [ expected, got: %s", res.Body.String())
		}

	})

	t.Run("invalid json 400", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		//mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(invalidJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusBadRequest {
			tt.Errorf("400 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "invalid character 'i' looking for beginning of value" {
			tt.Errorf("body `invalid character 'i' looking for beginning of value` expected, got: %s", res.Body.String())
		}
	})

	t.Run("empty object json 200", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(emptyObjectJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusNoContent {
			tt.Errorf("204 http status code expected, got: %d", res.Code)
		}
	})

	t.Run("empty object json 200", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(emptyObjectJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusNoContent {
			tt.Errorf("204 http status code expected, got: %d", res.Code)
		}
	})

	t.Run("invalid port object 500", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(invalidPortObject))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusInternalServerError {
			tt.Errorf("500 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "invalid character ','" {
			tt.Errorf("body `invalid character ','` expected, got: %s", res.Body.String())
		}
	})

	t.Run("invalid port object 500", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		//mockService.EXPECT().Add(gomock.Any()).Times(1632).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(nil)
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(invalidPortArray))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusInternalServerError {
			tt.Errorf("500 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "json: cannot unmarshal array into Go value of type map[string]interface {}" {
			tt.Errorf("body json: cannot unmarshal array into Go value of type map[string]interface {}` expected, got: %s", res.Body.String())
		}
	})

	t.Run("error service add", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		mockService.EXPECT().Add(gomock.Any()).Return(errors.New("test"))
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(fullJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusInternalServerError {
			tt.Errorf("500 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "test" {
			tt.Errorf("body 'test' expected, got: %s", res.Body.String())
		}
	})

	t.Run("error service flush", func(tt *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := service.NewMockPortService(ctrl)
		mockService.EXPECT().Add(gomock.Any()).Times(1).Return(nil)
		mockService.EXPECT().PortsBufferFlush().Times(1).Return(errors.New("error flush"))
		portsHandler := Ports{service: mockService}
		res := httptest.NewRecorder()
		portsRequest, err := http.NewRequest(http.MethodPost, "http://localtest/ports", bytes.NewBufferString(onePortJSON))
		if err != nil {
			tt.Error(err)
		}
		portsHandler.ServeHTTP(res, portsRequest)
		if res.Code != http.StatusInternalServerError {
			tt.Errorf("500 http status code expected, got: %d", res.Code)
		}
		if res.Body.String() != "error flush" {
			tt.Errorf("body 'error flush' expected, got: %s", res.Body.String())
		}
	})
}
