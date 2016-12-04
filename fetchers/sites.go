package fetchers

type SitesFetcher struct {
	client Client
}

func NewSitesFetcher(client Client) *SitesFetcher {
	return &SitesFetcher{client: client}
}

func (f *SitesFetcher) Sites() string {
	f.client.Request("/sites")
	return "foobar"
}
