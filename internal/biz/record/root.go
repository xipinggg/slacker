package record

type Root struct {
	record *Record
}

func (r *Root) Record() *Record {
	return r.record
}
