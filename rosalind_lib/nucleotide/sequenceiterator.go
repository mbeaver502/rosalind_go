package nucleotide

// SequenceIterator represents a basic iterator
// on a Nucleotide Sequence.
type SequenceIterator struct {
	seq *Sequence
	pos int
}

// Iter returns a new iterator on the Sequence.
func (ns Sequence) Iter() *SequenceIterator {
	return &SequenceIterator{
		seq: &ns,
		pos: 0,
	}
}

// Next advances the iterator by one position.
func (it *SequenceIterator) Next() *SequenceIterator {
	return &SequenceIterator{
		seq: it.seq,
		pos: it.pos + 1,
	}
}

// HasNext determines whether the iterator still has values.
func (it *SequenceIterator) HasNext() bool {
	return it.pos >= 0 && it.pos < len(*it.seq)
}

// Value returns the current value, if any, covered by the iterator.
func (it *SequenceIterator) Value() *Nucleotide {
	if it.HasNext() {
		return &(*it.seq)[it.pos]
	}
	return nil
}
