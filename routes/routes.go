package routes

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"

	"github.com/drew138/go-graphics/processing"
	"github.com/drew138/go-graphics/processing/formats"
	"github.com/drew138/go-graphics/processing/kernels"
	"github.com/gorilla/mux"
)

func enforceMultipartFormDataHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "multipart/form-data" {
			mediatype, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}
			if mediatype != "multipart/form-data" {
				http.Error(w, "Content-Type header must be multipart/form-data", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func processImage(w http.ResponseWriter, r *http.Request, k *kernels.Kernel) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("image")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	img := io.Reader(file)
	i, format, err := image.Decode(img)
	if !formats.IsSupportedFormat(format) {
		http.Error(w, "Unsupported image format", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	outputImg := processing.TransformImage(i, k)
	switch format {
	case "png":
		png.Encode(w, outputImg)
	case "jpg":
		jpeg.Encode(w, outputImg, nil)
	case "jpeg":
		jpeg.Encode(w, outputImg, nil)
	}
}

func sharpenEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.Sharpen)
}

func edgeDetectionEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.EdgeDetection)
}

func gaussianBlurEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.GaussianBlur)

}

func boxBlurEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.BoxBlur)
}

func customKernelEndpoint(w http.ResponseWriter, r *http.Request) {

	customKernel := kernels.Kernel{}
	vals := r.FormValue("kernel")
	err := json.Unmarshal([]byte(vals), &customKernel.Kernel)
	if err != nil {
		http.Error(w, "Error while parsing kernel", http.StatusInternalServerError)
		return
	}

	processImage(w, r, &customKernel)
}

func RegisterRoutes(r *mux.Router) {
	r.Handle("/api/sharpen", enforceMultipartFormDataHandler(http.HandlerFunc(sharpenEndpoint))).Methods("GET")
	r.Handle("/api/edgedetection", enforceMultipartFormDataHandler(http.HandlerFunc(edgeDetectionEndpoint))).Methods("GET")
	r.Handle("/api/gaussianblur", enforceMultipartFormDataHandler(http.HandlerFunc(gaussianBlurEndpoint))).Methods("GET")
	r.Handle("/api/boxblur", enforceMultipartFormDataHandler(http.HandlerFunc(boxBlurEndpoint))).Methods("GET")
	r.Handle("/api/custom", enforceMultipartFormDataHandler(http.HandlerFunc(customKernelEndpoint))).Methods("GET")
}
