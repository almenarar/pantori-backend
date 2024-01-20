package mocks

type ImageMock struct {
	Invocation *string
}

func (im *ImageMock) GetImageURL(string) string {
	*im.Invocation = *im.Invocation + "-GetImage"
	return "url"
}
