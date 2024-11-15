package gollama

import "errors"

func (c *Gollama) Vision(prompt string, images []string) (string, error) {
	if len(images) == 0 {
		return "", errors.New("no images provided")
	}

	base64images := make([]string, len(images))
	for i, image := range images {
		base64image, err := base64EncodeFile(image)
		if err != nil {
			return "", err
		}
		base64images[i] = base64image
	}

}
