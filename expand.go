package expand

import (
	"fmt"
)

var expand_start int = 1
var _ = expand_start
var expand_first_final int = 5
var _ = expand_first_final
var expand_error int = 0
var _ = expand_error
var expand_en_main int = 1
var _ = expand_en_main

func (e *Expander) Expand(data string) ([]string, error) {
	e.clear()
	cs, p, pe := 0, 0, len(data)
	{
		cs = int(expand_start)

	}
	{
		switch cs {
		case 1:
			goto st_case_1
		case 5:
			goto st_case_5
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4

		}
		p += 1
	st_case_1:
		if p == pe {
			goto _out1

		}
		switch data[p] {
		case 44:
			{
				goto _st0

			}
		case 91:
			{
				goto _ctr3

			}
		case 93:
			{
				goto _st0

			}

		}
		goto _ctr2
	_ctr2:
		{
			e.chars = e.chars + string((data[p]))
		}
		goto _st5
	_ctr8:
		{
			e.doMap()
			e.doReduce()
		}
		goto _st5
	_st5:
		p += 1
	st_case_5:
		if p == pe {
			goto _out5

		}
		switch data[p] {
		case 44:
			{
				goto _st0

			}
		case 91:
			{
				goto _ctr3

			}
		case 93:
			{
				goto _st0

			}

		}
		goto _ctr2
	_st0:
	st_case_0:
		goto _out0
	_ctr3:
		{
			e.putChars()
		}
		goto _st2
	_st2:
		p += 1
	st_case_2:
		if p == pe {
			goto _out2

		}
		switch data[p] {
		case 44:
			{
				goto _st0

			}
		case 91:
			{
				goto _st0

			}
		case 93:
			{
				goto _st0

			}

		}
		goto _ctr5
	_ctr5:
		{
			e.opt = e.opt + string((data[p]))
		}
		goto _st3
	_st3:
		p += 1
	st_case_3:
		if p == pe {
			goto _out3

		}
		switch data[p] {
		case 44:
			{
				goto _ctr7

			}
		case 91:
			{
				goto _st0

			}
		case 93:
			{
				goto _ctr8

			}

		}
		goto _ctr5
	_ctr7:
		{
			e.doMap()
		}
		goto _st4
	_st4:
		p += 1
	st_case_4:
		if p == pe {
			goto _out4

		}
		switch data[p] {
		case 44:
			{
				goto _st0

			}
		case 91:
			{
				goto _st0

			}
		case 93:
			{
				goto _ctr8

			}

		}
		goto _ctr5
	_out1:
		cs = 1
		goto _out
	_out5:
		cs = 5
		goto _out
	_out0:
		cs = 0
		goto _out
	_out2:
		cs = 2
		goto _out
	_out3:
		cs = 3
		goto _out
	_out4:
		cs = 4
		goto _out
	_out:
		{

		}

	}
	if cs < expand_first_final {
		return nil, fmt.Errorf("expand parse error: %s", data)
	}

	e.putChars()

	return e.result, nil
}
