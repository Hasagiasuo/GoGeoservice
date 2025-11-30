package http

import "net/http"

func NewRouter(dh *DriverHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/driver", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			dh.AddDriver(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusBadRequest)
		}
	})
	mux.HandleFunc("/drivers/nerdy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			dh.GetNerdyDrivers(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusBadRequest)
		}
	})
	return mux
}
