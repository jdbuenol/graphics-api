package routes

import (
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

func SharpenEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.Sharpen)
}

func EdgeDetectionEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.EdgeDetection)
}

func GaussianBlurEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.GaussianBlur)

}

func BoxBlurEndpoint(w http.ResponseWriter, r *http.Request) {
	processImage(w, r, &kernels.BoxBlur)
}

func CustomKernelEndpoint(w http.ResponseWriter, r *http.Request) {
	customKernel := kernels.Kernel{}
	processImage(w, r, &customKernel)
}

func RegisterRoutes(r *mux.Router) {
	r.Handle("/sharpen", enforceMultipartFormDataHandler(http.HandlerFunc(SharpenEndpoint))).Methods("GET")
	r.Handle("/edgedetection", enforceMultipartFormDataHandler(http.HandlerFunc(EdgeDetectionEndpoint))).Methods("GET")
	r.Handle("/gaussianblur", enforceMultipartFormDataHandler(http.HandlerFunc(GaussianBlurEndpoint))).Methods("GET")
	r.Handle("/boxblur", enforceMultipartFormDataHandler(http.HandlerFunc(BoxBlurEndpoint))).Methods("GET")
	r.Handle("/custom", enforceMultipartFormDataHandler(http.HandlerFunc(CustomKernelEndpoint))).Methods("GET")
}
