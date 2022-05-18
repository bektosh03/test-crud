package urlpaginator

import "fmt"

type URLPaginator struct {
	url      string
	maxPage  int
	currPage int
}

func New(url string, maxPage int) *URLPaginator {
	return &URLPaginator{
		maxPage:  maxPage,
		url:      url,
		currPage: 0,
	}
}

func (u *URLPaginator) NextPage() (ok bool) {
	if u.currPage >= u.maxPage {
		return false
	}

	u.currPage++
	return true
}

func (u *URLPaginator) URL() string {
	return fmt.Sprintf("%s?page=%d", u.url, u.currPage)
}

func (u *URLPaginator) CurrentPage() int {
	return u.currPage
}
