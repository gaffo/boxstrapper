// line 1 "ops_parser.rl"
/* this file is outputted directly to the go file unless we are inside a %% block */

package ops_parser

// no imports
import (
	"errors"
	"fmt"
	"strings"
)

// line 14 "ops_parser.rl"

// line 115 "ops_parser.rl"

// line 24 "ops_parser.go"
const operations_start int = 1
const operations_first_final int = 17
const operations_error int = 0

const operations_en_main int = 1

// line 118 "ops_parser.rl"

type Dep struct {
	Operation string
	Params    []string
}

type Op struct {
	Name    string
	Params  []string
	Groups  []string
	Depends []Dep
}

func ParsePackages(data string) ([]Op, error) {
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

	// line 63 "ops_parser.go"
	{
		cs = operations_start
	}

	// line 148 "ops_parser.rl"

	// output the code for operating the parser

	// line 72 "ops_parser.go"
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
		case 7:
			goto st_case_7
		case 17:
			goto st_case_17
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
		case 14:
			goto st_case_14
		case 18:
			goto st_case_18
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		}
		goto st_out
	st_case_1:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:
		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st2
	tr9:
		// line 55 "ops_parser.rl"

		// fmt.Println("line")
		line++
		lastLine = p

		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		// line 164 "ops_parser.go"
		if data[p] == 40 {
			goto tr2
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
	tr2:
		// line 24 "ops_parser.rl"

		// fmt.Println("operation", curToken)
		op.Name = curToken
		curToken = ""

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		// line 195 "ops_parser.go"
		if data[p] == 126 {
			goto tr3
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	tr3:
		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		// line 225 "ops_parser.go"
		switch data[p] {
		case 41:
			goto tr4
		case 44:
			goto tr5
		case 126:
			goto tr3
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	tr4:
		// line 31 "ops_parser.rl"

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
		// line 262 "ops_parser.go"
		if data[p] == 58 {
			goto st6
		}
		goto st0
	tr21:
		// line 39 "ops_parser.rl"

		// fmt.Println("group", curToken)
		op.Groups = append(op.Groups, curToken)
		// fmt.Println(">", op.Groups)
		curToken = ""

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		// line 282 "ops_parser.go"
		if data[p] == 32 {
			goto st7
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
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
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
		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		// line 330 "ops_parser.go"
		switch data[p] {
		case 10:
			goto tr19
		case 32:
			goto tr20
		case 44:
			goto tr21
		case 62:
			goto tr22
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
	tr19:
		// line 39 "ops_parser.rl"

		// fmt.Println("group", curToken)
		op.Groups = append(op.Groups, curToken)
		// fmt.Println(">", op.Groups)
		curToken = ""

		// line 47 "ops_parser.rl"

		// fmt.Println("opSpec", op)
		ops = append(ops, op)
		op = Op{}

		goto st8
	tr23:
		// line 76 "ops_parser.rl"

		// fmt.Println("dep", dep)
		op.Depends = append(op.Depends, dep)
		dep = Dep{}

		// line 47 "ops_parser.rl"

		// fmt.Println("opSpec", op)
		ops = append(ops, op)
		op = Op{}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		// line 392 "ops_parser.go"
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr9
			}
		default:
			goto tr9
		}
		goto st0
	tr20:
		// line 39 "ops_parser.rl"

		// fmt.Println("group", curToken)
		op.Groups = append(op.Groups, curToken)
		// fmt.Println(">", op.Groups)
		curToken = ""

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		// line 421 "ops_parser.go"
		if data[p] == 62 {
			goto st10
		}
		goto st0
	tr22:
		// line 39 "ops_parser.rl"

		// fmt.Println("group", curToken)
		op.Groups = append(op.Groups, curToken)
		// fmt.Println(">", op.Groups)
		curToken = ""

		goto st10
	tr24:
		// line 76 "ops_parser.rl"

		// fmt.Println("dep", dep)
		op.Depends = append(op.Depends, dep)
		dep = Dep{}

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		// line 450 "ops_parser.go"
		if data[p] == 32 {
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
	tr12:
		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		// line 498 "ops_parser.go"
		if data[p] == 40 {
			goto tr13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
	tr13:
		// line 62 "ops_parser.rl"

		// fmt.Println("dep_operation", curToken)
		dep.Operation = curToken
		curToken = ""

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		// line 529 "ops_parser.go"
		if data[p] == 126 {
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
	tr14:
		// line 18 "ops_parser.rl"

		// fmt.Println("Char", string(fc))
		curToken += string(data[p])

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		// line 559 "ops_parser.go"
		switch data[p] {
		case 41:
			goto tr15
		case 44:
			goto tr16
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
	tr15:
		// line 69 "ops_parser.rl"

		// fmt.Println("dep_param", curToken)
		dep.Params = append(dep.Params, curToken)
		curToken = ""

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		// line 595 "ops_parser.go"
		switch data[p] {
		case 10:
			goto tr23
		case 44:
			goto tr24
		}
		goto st0
	tr16:
		// line 69 "ops_parser.rl"

		// fmt.Println("dep_param", curToken)
		dep.Params = append(dep.Params, curToken)
		curToken = ""

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		// line 617 "ops_parser.go"
		switch data[p] {
		case 32:
			goto st13
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
	tr5:
		// line 31 "ops_parser.rl"

		// fmt.Println("param", curToken)
		op.Params = append(op.Params, curToken)
		// fmt.Println(">", op.Params)
		curToken = ""

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		// line 652 "ops_parser.go"
		switch data[p] {
		case 32:
			goto st3
		case 126:
			goto tr3
		}
		switch {
		case data[p] < 65:
			if 46 <= data[p] && data[p] <= 57 {
				goto tr3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 17:
				// line 39 "ops_parser.rl"

				// fmt.Println("group", curToken)
				op.Groups = append(op.Groups, curToken)
				// fmt.Println(">", op.Groups)
				curToken = ""

				// line 47 "ops_parser.rl"

				// fmt.Println("opSpec", op)
				ops = append(ops, op)
				op = Op{}

			case 18:
				// line 76 "ops_parser.rl"

				// fmt.Println("dep", dep)
				op.Depends = append(op.Depends, dep)
				dep = Dep{}

				// line 47 "ops_parser.rl"

				// fmt.Println("opSpec", op)
				ops = append(ops, op)
				op = Op{}

				// line 725 "ops_parser.go"
			}
		}

	_out:
		{
		}
	}

	// line 151 "ops_parser.rl"

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
			char+1,
			ctxLine,
			ctx)

		return nil, errors.New(error)
	}

	return ops, nil
}
