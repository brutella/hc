package http

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"
	"github.com/xiam/to"

	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type CharacteristicsResponse struct {
	Characteristics []CharacteristicResponse `json:"characteristics"`
}

type CharacteristicResponse struct {
	AccessoryID      uint64      `json:"aid"`
	CharacteristicID uint64      `json:"iid"`
	Value            interface{} `json:"value,omitempty"`

	// Status contains the status code. Should be interpreted as integer.
	// The property is omitted if not specified, which makes the payload smaller.
	Status *int `json:"status,omitempty"`
}

type CharacteristicRequest struct {
	AccessoryID      uint64      `json:"aid"`
	CharacteristicID uint64      `json:"iid"`
	Value            interface{} `json:"value"`
	Events           interface{} `json:"ev,omitempty"`
}

// Authenticate verfies that a sesson for the request is available.
func (srv *Server) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
		sess := srv.context.GetSessionForRequest(r)
		if sess == nil {
			w.WriteHeader(470) // this custom status code indicates an error
			if err := WriteJSON(w, r, &ErrResponse{Status: hap.StatusInsufficientPrivileges}); err != nil {
				log.Debug.Println(err)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Characteristics Handles GET and PUT requests for the /characteristics endpoint.
func (srv *Server) Characteristics(w http.ResponseWriter, r *http.Request) {
	sess := srv.context.GetSessionForRequest(r)
	conn := sess.Connection()

	switch r.Method {
	case hap.MethodGET:
		r.ParseForm()
		log.Debug.Printf("%v GET /characteristics %v", r.RemoteAddr, r.Form)

		var (
			err = false // indicates if a characteristic was not found
			arr []CharacteristicResponse
		)

		// id=1.4,1.5
		strs := strings.Split(r.Form.Get("id"), ",")
		for _, str := range strs {
			if ids := strings.Split(str, "."); len(ids) == 2 {
				aid := to.Uint64(ids[0]) // accessory id
				iid := to.Uint64(ids[1]) // instance id (= characteristic id)
				resp := CharacteristicResponse{AccessoryID: aid, CharacteristicID: iid}
				if ch := srv.getCharacteristic(aid, iid); ch != nil {
					resp.Value = ch.GetValueFromConnection(conn)
				} else {
					err = true
					status := hap.StatusServiceCommunicationFailure
					resp.Status = &status
				}
				arr = append(arr, resp)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if err == true {
			// Set 207 status when any of the response includes an error
			w.WriteHeader(http.StatusMultiStatus)
			for _, resp := range arr {
				if resp.Status == nil {
					ok := 0
					resp.Status = &ok // make sure that every response contains a status code (0 means OK)
				}
			}
		} else {
			w.WriteHeader(http.StatusOK)
		}

		WriteJSON(w, r, &CharacteristicsResponse{arr})
	case hap.MethodPUT:
		log.Debug.Printf("%v PUT /characteristics", r.RemoteAddr)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Debug.Println(err)
		} else {
			log.Debug.Println(string(b))
		}

		req := struct {
			Characteristics []CharacteristicRequest `json:"characteristics"`
		}{}
		err = JSONDecode(bytes.NewBuffer(b), &req)
		if err != nil {
			log.Debug.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			if err := WriteJSON(w, r, &ErrResponse{Status: hap.StatusInvalidValueInRequest /*?*/}); err != nil {
				log.Debug.Println(err)
			}
			return
		}

		resp := &CharacteristicsResponse{}
		for _, ch := range req.Characteristics {
			c := srv.getCharacteristic(ch.AccessoryID, ch.CharacteristicID)
			if c == nil {
				log.Info.Printf("Could not find characteristic with aid %d and iid %d\n", ch.AccessoryID, ch.CharacteristicID)
				continue
			}

			if ch.Value != nil {
				c.UpdateValueFromConnection(ch.Value, conn)
			}

			if ch.Events != nil {
				if !c.IsObservable() {
					status := hap.StatusNotificationNotSupported
					err := CharacteristicResponse{AccessoryID: ch.AccessoryID, CharacteristicID: ch.CharacteristicID, Status: &status}
					resp.Characteristics = append(resp.Characteristics, err)
					continue
				}

				if events, ok := ch.Events.(bool); ok == true {
					if events {
						sess.Subscribe(c)
					} else {
						sess.Unsubscribe(c)
					}
				}
			}
		}

		if len(resp.Characteristics) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		WriteJSON(w, r, resp)

	default:
		log.Debug.Println("Cannot handle HTTP method", r.Method)
	}
}
