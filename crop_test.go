package imageprocess

import "testing"

func Test_Crop(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	img = CropImage(img, CropOption{Position: SouthEast, Width: 900, Height: 900})
	SaveImage(img, "examples/out.jpg", f, 100)
}
