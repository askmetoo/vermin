package images

import (
	"fmt"
	"strings"
)

func CheckCanMount(image string) error {
	remote, err := listRemoteImages(false)
	if err != nil {
		return err
	}

	dbImage, err := remote.findByName(image)
	if err != nil {
		return err
	}

	if !dbImage.Mount {
		mounted := remote.findByMount(true).names()
		return fmt.Errorf("Image '%s' cannot be used with mount, "+
			"images can be used are:\n%s", image, strings.Join(mounted, "\n"))
	}

	return nil
}
