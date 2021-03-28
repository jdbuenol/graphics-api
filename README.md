# GRAPHICS API

API designed to perform image processing on jpg, jpeg and png images.

## AVAILABLE ENDPOINTS

Currently available endpoints listed below:

```text
/api/sharpen
/api/edgedetection
/api/gaussianblur
/api/boxblur
/api/custom
```

Supplying an image in a multipart/form is required for all of the endpoints.
In addition, the `/api/custom` requires provissioning a convolution matrix in the form `[[val1,val2,val3],[val4,val5,val6],[val7,val8,val9]]`. Both of these must be supplied as `image` and `kernel` attributes respectively in the form.
