
// line 1 "ops_parser/ops_parser.rl"
/* this file is outputted directly to the go file unless we are inside a %% block */

package ops_parser

// no imports
import (
  "fmt"
  "errors"
  "strings"
)


// line 14 "ops_parser/ops_parser.rl"



// line 114 "ops_parser/ops_parser.rl"



// line 24 "ops_parser/ops_parser.go"
const operations_start int = 1
const operations_first_final int = 18
const operations_error int = 0

const operations_en_main int = 1


// line 117 "ops_parser/ops_parser.rl"

type Dep struct {
  Operation string
  Params []string
}

type Op struct {
  Name string
  Params []string
  Groups []string
  Depends []Dep
}

func ParseOps(data string) ([]Op, error) {
  // Declare std operating variables for the parser, ?, pointer, point end
  cs, p, pe := 0, 0, len(data)
  eof := pe

  // declare our work variables
  curToken := ""
  ops := make([]Op, 0, 32)
  op := Op{}
  dep := Dep{}
  line := 1
  lastLine := 0

  // %% this just fixes our syntax highlighting...

  // output the code for setting up the parser
  
// line 63 "ops_parser/ops_parser.go"
	{
	cs = operations_start
	}

// line 147 "ops_parser/ops_parser.rl"

  // output the code for operating the parser
  
// line 72 "ops_parser/ops_parser.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 18:
		goto st_case_18
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 19:
		goto st_case_19
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	}
	goto st_out
tr9:
// line 55 "ops_parser/ops_parser.rl"


    // fmt.Println("line")
    line++
    lastLine = p
  
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
// line 134 "ops_parser/ops_parser.go"
		if data[p] == 32 {
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st2
tr10:
// line 55 "ops_parser/ops_parser.rl"


    // fmt.Println("line")
    line++
    lastLine = p
  
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
// line 183 "ops_parser/ops_parser.go"
		if data[p] == 40 {
			goto tr3
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
tr3:
// line 24 "ops_parser/ops_parser.rl"


    // fmt.Println("operation", curToken)
    op.Name = curToken
    curToken = ""
  
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
// line 214 "ops_parser/ops_parser.go"
		if data[p] == 126 {
			goto tr4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr4
			}
		default:
			goto tr4
		}
		goto st0
tr4:
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
// line 244 "ops_parser/ops_parser.go"
		switch data[p] {
		case 41:
			goto tr5
		case 44:
			goto tr6
		case 126:
			goto tr4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr4
			}
		default:
			goto tr4
		}
		goto st0
tr5:
// line 31 "ops_parser/ops_parser.rl"


    // fmt.Println("param", curToken)
    op.Params = append(op.Params, curToken)
    // fmt.Println(">", op.Params)
    curToken = ""
  
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
// line 281 "ops_parser/ops_parser.go"
		if data[p] == 58 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 32 {
			goto st6
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr8:
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
// line 320 "ops_parser/ops_parser.go"
		switch data[p] {
		case 10:
			goto tr21
		case 32:
			goto tr22
		case 44:
			goto tr23
		case 62:
			goto tr24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr21:
// line 39 "ops_parser/ops_parser.rl"


    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  
// line 47 "ops_parser/ops_parser.rl"


    // fmt.Println("opSpec", op)
    ops = append(ops, op)
    op = Op{}
  
	goto st7
tr25:
// line 76 "ops_parser/ops_parser.rl"


    // fmt.Println("dep", dep)
    op.Depends = append(op.Depends, dep)
    dep = Dep{}
  
// line 47 "ops_parser/ops_parser.rl"


    // fmt.Println("opSpec", op)
    ops = append(ops, op)
    op = Op{}
  
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
// line 382 "ops_parser/ops_parser.go"
		if data[p] == 32 {
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr10
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto st0
tr22:
// line 39 "ops_parser/ops_parser.rl"


    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
// line 414 "ops_parser/ops_parser.go"
		if data[p] == 62 {
			goto st9
		}
		goto st0
tr24:
// line 39 "ops_parser/ops_parser.rl"


    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  
	goto st9
tr26:
// line 76 "ops_parser/ops_parser.rl"


    // fmt.Println("dep", dep)
    op.Depends = append(op.Depends, dep)
    dep = Dep{}
  
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
// line 443 "ops_parser/ops_parser.go"
		if data[p] == 32 {
			goto st10
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
tr13:
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
// line 491 "ops_parser/ops_parser.go"
		if data[p] == 40 {
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
tr14:
// line 62 "ops_parser/ops_parser.rl"


    // fmt.Println("dep_operation", curToken)
    dep.Operation = curToken
    curToken = ""
  
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
// line 522 "ops_parser/ops_parser.go"
		if data[p] == 126 {
			goto tr15
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr15:
// line 18 "ops_parser/ops_parser.rl"


    // fmt.Println("Char", string(fc))
    curToken += string(data[p])
  
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
// line 552 "ops_parser/ops_parser.go"
		switch data[p] {
		case 41:
			goto tr16
		case 44:
			goto tr17
		case 126:
			goto tr15
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr16:
// line 69 "ops_parser/ops_parser.rl"


    // fmt.Println("dep_param", curToken)
    dep.Params = append(dep.Params, curToken)
    curToken = ""
  
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
// line 588 "ops_parser/ops_parser.go"
		switch data[p] {
		case 10:
			goto tr25
		case 44:
			goto tr26
		}
		goto st0
tr17:
// line 69 "ops_parser/ops_parser.rl"


    // fmt.Println("dep_param", curToken)
    dep.Params = append(dep.Params, curToken)
    curToken = ""
  
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
// line 610 "ops_parser/ops_parser.go"
		switch data[p] {
		case 32:
			goto st12
		case 126:
			goto tr15
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr23:
// line 39 "ops_parser/ops_parser.rl"


    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
// line 645 "ops_parser/ops_parser.go"
		if data[p] == 32 {
			goto st16
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr6:
// line 31 "ops_parser/ops_parser.rl"


    // fmt.Println("param", curToken)
    op.Params = append(op.Params, curToken)
    // fmt.Println(">", op.Params)
    curToken = ""
  
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
// line 695 "ops_parser/ops_parser.go"
		switch data[p] {
		case 32:
			goto st3
		case 126:
			goto tr4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr4
			}
		default:
			goto tr4
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 18:
// line 39 "ops_parser/ops_parser.rl"


    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  
// line 47 "ops_parser/ops_parser.rl"


    // fmt.Println("opSpec", op)
    ops = append(ops, op)
    op = Op{}
  
		case 19:
// line 76 "ops_parser/ops_parser.rl"


    // fmt.Println("dep", dep)
    op.Depends = append(op.Depends, dep)
    dep = Dep{}
  
// line 47 "ops_parser/ops_parser.rl"


    // fmt.Println("opSpec", op)
    ops = append(ops, op)
    op = Op{}
  
// line 770 "ops_parser/ops_parser.go"
		}
	}

	_out: {}
	}

// line 150 "ops_parser/ops_parser.rl"

  // fmt.Println("Final State:", cs, p, pe)

  if cs < operations_first_final {
    char := p - lastLine

    ctx := ""
    for i := 0; i < char; i += 1 {
      ctx += "-"
    }
    ctx += "^"

    ctxLine := data[lastLine:]
    endOfLineIndex := strings.IndexAny(ctxLine, "\n")
    if endOfLineIndex != -1 {
      ctxLine = ctxLine[:endOfLineIndex]
    }

    error := fmt.Sprintf(
      "Parse error on line %d char %d:\n%s\n%s",
      line,
      char + 1,
      ctxLine,
      ctx)

    return nil, errors.New(error)
  }

  return ops, nil
}
