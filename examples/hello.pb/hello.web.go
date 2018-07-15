package hello_pb

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chai2010/pbgo"
)

func HelloServiceHandler(svc HelloServiceInterface) http.Handler {
	router := httprouter.New()

	router.Handle("GET", "/hello/:value",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			var (
				protoReq   String
				protoReply String
			)

			for _, fieldPath := range []string{"value"} {
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

			if err := json.NewDecoder(r.Body).Decode(&protoReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			for _, fieldPath := range []string{"value"} {
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
