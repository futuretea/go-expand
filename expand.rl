package expand

import (
  "fmt"
)

%%{
  machine expand;
  write data;
}%%

func (e *Expander) Expand(data string) ([]string, error) {
  e.clear()
  cs, p, pe := 0, 0, len(data)
  %%{

    action read_chars {
      e.chars = e.chars + string(fc)
    }

    action enter_option {
      e.putChars()
    }

    action read_opt {
       e.opt = e.opt + string(fc)
    }

    action read_comma {
      e.doMap()
    }

    action leave_option {
      e.doMap()
      e.doReduce()
    }

    chars = ^[\[,\],\,];
    option = (chars @read_opt)+ (',' @read_comma)?;
    options = ('[' @enter_option) option+ (']' @leave_option);
    patt = (chars @read_chars)+ options? |
           options (chars @read_chars)?;
    main := patt+;

    write init;
    write exec;
  }%%

  if cs < expand_first_final {
    return nil, fmt.Errorf("expand parse error: %s", data)
  }

  e.putChars()

  return e.result, nil
}
