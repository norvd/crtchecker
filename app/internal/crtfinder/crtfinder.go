package crtfinder

import "github.com/norvd/crtchecker/internal/requester"

type CrtFinder struct {
	cr *requester.CrtRequester
}

func New(url string) *CrtFinder {
	return &CrtFinder{
		cr: requester.New(url + "/json"),
	}
}

func (cf *CrtFinder) Find(identity string) []byte {
	params := map[string]string{"q": identity}

	result := cf.cr.GetWithParams(params)

	return result

}
