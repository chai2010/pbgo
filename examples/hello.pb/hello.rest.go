package hello_pb

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/chai2010/pbgo"
)

func HelloServiceHandler(svc HelloServiceInterface) http.Handler {
	router := httprouter.New()

	var re = regexp.MustCompile(`(\*|\:)(\w|\.)+`)

	router.Handle("GET", "/hello/:value",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			for _, fieldPath := range re.FindAllString("/hello/:value", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	router.Handle("POST", "/hello",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			if strings.Contains(r.Header.Get("Accept"), "application/json") {
				w.Header().Set("Content-Type", "application/json")
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}

			if err := json.NewDecoder(r.Body).Decode(&protoReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			for _, fieldPath := range re.FindAllString("/hello", -1) {
				fieldPath := strings.TrimLeft(fieldPath, ":*")
				err := pbgo.PopulateFieldFromPath(&protoReq, fieldPath, ps.ByName(fieldPath))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			if err := pbgo.PopulateQueryParameters(&protoReq, r.URL.Query()); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := svc.Hello(&protoReq, &protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	return router
}
