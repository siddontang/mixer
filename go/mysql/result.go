package mysql

type Result struct {
	Status uint16

	InsertId     uint64
	AffectedRows uint64
}

func (r *Result) GetStatus() uint16 {
	return r.Status
}

func (r *Result) LastInsertId() (int64, error) {
	return int64(r.InsertId), nil
}

func (r *Result) RowsAffected() (int64, error) {
	return int64(r.AffectedRows), nil
}
