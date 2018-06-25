package page

import (
	"errors"
	"math"
)

var (
	ErrEndPage = errors.New("end page")
)

type PageObj struct {
	Pageno     int64
	TotalRows  int64
	TotalPages int64

	RowCnt int64
	RowBgn int64
}

func NewPageObj(pageno, totalRows, rowCnt int64) (*PageObj, error) {
	p := new(PageObj)

	err := p.InitPage(pageno, totalRows, rowCnt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *PageObj) InitPage(pageno, totalRows, rowCnt int64) error {
	if pageno <= 0 {
		return errors.New("invalid pageno")
	}
	if totalRows < 0 {
		return errors.New("invalid  totalRows")
	}
	if rowCnt <= 0 {
		return errors.New("invalid rowCnt")
	}

	p.TotalRows = totalRows
	p.RowCnt = rowCnt
	if totalRows == 0 {
		p.TotalPages = 1
		p.Pageno = 1
	} else {
		tp := math.Ceil(float64(p.TotalRows) / float64(p.RowCnt))
		p.TotalPages = int64(tp)
		p.Pageno = pageno
	}

	p.RowBgn = (p.Pageno - 1) * p.RowCnt
	if p.RowBgn+p.RowCnt > p.TotalRows {
		p.RowCnt = p.TotalRows - p.RowBgn
	}

	return nil
}

func (p *PageObj) Rewind() error {
	return p.InitPage(1, p.TotalRows, p.RowCnt)
}

func (p *PageObj) Next() error {
	if p.Pageno == p.TotalPages {
		return ErrEndPage
	}

	return p.InitPage(p.Pageno+1, p.TotalRows, p.RowCnt)
}
