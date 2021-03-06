package search

type SynonymIterator struct {
	synonyms []Synonym
	pos      int
}

func newSynonymIterator(synonyms []Synonym) *SynonymIterator {
	return &SynonymIterator{synonyms: synonyms}
}

func (it *SynonymIterator) Next(opts ...interface{}) (Synonym, error) {
	if it.pos >= len(it.synonyms) {
		return nil, NoMoreSynonymsErr
	}
	synonym := it.synonyms[it.pos]
	it.pos++
	return synonym, nil
}
