package query

type Filterer interface {
	ListFilterer
	DetailFilterer
}

type Factory struct {
	sf SorterFactory
	F  Filterer
}

func NewFactory(sf SorterFactory, f Filterer) *Factory {
	return &Factory{
		sf: sf,
		F:  f,
	}
}

func (f *Factory) NewDetail(query IDetail) (*Detail, error) {
	return NewDetailFrom(query, f.F)
}

func (f *Factory) NewList(query IList) (*List, error) {
	return NewListFrom(query, f.sf, f.F)
}
