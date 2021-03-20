package routes

import (
	"encoding/json"
	"image"
	"io"
	"net/http"

	processing "github.com/drew138/go-graphics/processing"
	kernels "github.com/drew138/go-graphics/processing/kernels"
	"github.com/gorilla/mux"
)

type handler func(http.ResponseWriter, *http.Request)

func checkBody(f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := r.GetBody()
		if err != nil {

		}
		img := io.Reader(body)
		i, format, err := image.Decode(img)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"Error": "An error occurred while decoding the image"})
			w.WriteHeader(400)
			return
		}
		// b := r.Body()
		if true {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"Error": "An image was not provided, or its format is not currently supported"})
			w.WriteHeader(400)
			return
		}
		f(w, r)
	}
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/leaderboard", checkBody(func(w http.ResponseWriter, r *http.Request) {
		processing.TransformImage(i, kernels.BoxBlur)
	})).Methods("GET")
}

// ChangePassword - change password for a given user

func SharpenEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func EdgeDetectionEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func GaussianBlurEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func CustomKernelEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
