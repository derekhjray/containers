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

func TestContainerInspect(t *testing.T) {
	runtime, err := New()
	if err != nil {
		t.Error(err)
		return
	}

	container, err := runtime.ContainerInspect(context.TODO(), "openssh-server")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(container)
}
