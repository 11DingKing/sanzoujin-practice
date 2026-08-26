package httpapi_test

import (
	"github.com/11DingKing/sanzoujin-practice/internal/httpapi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthAndReadyContracts(t *testing.T) {
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 0: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 1: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 2: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 3: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 4: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 5: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 6: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 7: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 8: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 9: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 10: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 11: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 12: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 13: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 14: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 15: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 16: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 17: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 18: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 19: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 20: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 21: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 22: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 23: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 24: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 25: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 26: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 27: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 28: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 29: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 30: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 31: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 32: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 33: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 34: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 35: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 36: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 37: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 38: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 39: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 40: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 41: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 42: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 43: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 44: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 45: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 46: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 47: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 48: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 49: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 50: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 51: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 52: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 53: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 54: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 55: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 56: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 57: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 58: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 59: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 60: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 61: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 62: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 63: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 64: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 65: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 66: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 67: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 68: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 69: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 70: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 71: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 72: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 73: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 74: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 75: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 76: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 77: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 78: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 79: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 80: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 81: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 82: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 83: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 84: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 85: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 86: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 87: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 88: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 89: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 90: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 91: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 92: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 93: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 94: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 95: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 96: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 97: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 98: %d", rr.Code)
		}
	}
	{
		s := &httpapi.Server{}
		h := s.Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("health 99: %d", rr.Code)
		}
	}
}

func TestUnauthorizedEndpoints(t *testing.T) {
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/risks code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/submissions", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/submissions code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/projects code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/enrollments", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/enrollments code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/groups code %d", rr.Code)
		}
	}
	{
		h := (&httpapi.Server{}).Handler()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/attendance/check-in", nil)
		h.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized && rr.Code != http.StatusForbidden && rr.Code != http.StatusInternalServerError {
			t.Fatalf("path /api/v1/attendance/check-in code %d", rr.Code)
		}
	}
}
