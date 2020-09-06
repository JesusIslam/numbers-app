package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestNumbersHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/numbersapi/:number")
	c.SetParamNames("number")
	c.SetParamValues("1")

	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"numberhandler", args{c}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NumbersHandler(tt.args.ctx); (err != nil) != (tt.wantErr == nil) {
				t.Errorf("NumbersHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
