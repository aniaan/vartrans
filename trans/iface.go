package trans

type Fetcher interface {
	fetch(q string) ([]string, error)
}