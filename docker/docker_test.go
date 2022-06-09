package docker

import (
	"context"
	"testing"
)

func TestImageInspect(t *testing.T) {
	runtime, err := New()
	if err != nil {
		t.Error(err)
		return
	}

	image, err := runtime.ImageInspect(context.TODO(), "tomcat:latest")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(image)
}
