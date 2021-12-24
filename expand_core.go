package expand

import (
	"fmt"
	"strconv"
	"strings"
)

type Expander struct {
	chars  string
	opt    string
	result []string
	tmp    [][]string
}

func NewExpander() Expander {
	return Expander{
		result: []string{},
	}
}

func (e *Expander) clear() {
	e.chars = ""
	e.opt = ""
	e.result = []string{}
	e.tmp = [][]string{}
}

func (e *Expander) putChars() {
	if len(e.chars) == 0 {
		return
	}
	if len(e.result) == 0 {
		e.result = []string{""}
	}
	for i, str := range e.result {
		e.result[i] = str + e.chars
	}
	e.chars = ""
}

func (e *Expander) doReduce() {
	e.result = []string{}
	for _, tmpOne := range e.tmp {
		for _, s := range tmpOne {
			e.result = append(e.result, s)
		}
	}
	e.tmp = [][]string{}
}

func (e *Expander) putOpt(opt string) {
	if len(opt) == 0 {
		return
	}
	if len(e.result) == 0 {
		e.result = []string{""}
	}
	exp := make([]string, len(e.result))
	for i, str := range e.result {
		exp[i] = str + opt
	}
	e.tmp = append(e.tmp, exp)
}

func (e *Expander) doMap() {
	if len(e.opt) == 0 {
		return
	}
	optRange := strings.Split(e.opt, "-")
	if len(optRange) != 2 {
		e.putOpt(e.opt)
	} else {
		start, startErr := strconv.Atoi(optRange[0])
		end, endErr := strconv.Atoi(optRange[1])
		if startErr != nil || endErr != nil {
			e.putOpt(e.opt)
		} else {
			if end > start {
				for n := start; n <= end; n++ {
					e.putOpt(fmt.Sprintf("%d", n))
				}
			} else {
				for n := start; n >= end; n-- {
					e.putOpt(fmt.Sprintf("%d", n))
				}
			}
		}
	}
	e.opt = ""
}
