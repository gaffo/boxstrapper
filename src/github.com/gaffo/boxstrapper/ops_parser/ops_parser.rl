/* this file is outputted directly to the go file unless we are inside a %% block */

package ops_parser

// no imports
import (
  "fmt"
  "errors"
  "strings"
)

%%{
  machine operations;
}%%

%%{
  # every character we append to curToken, our work
  action char {
    // fmt.Println("Char", string(fc))
    curToken += string(fc)
  }

  # set the current OPs name to be curToken
  action op {
    // fmt.Println("operation", curToken)
    op.Name = curToken
    curToken = ""
  }

  # append the cur token to the params list
  action param {
    // fmt.Println("param", curToken)
    op.Params = append(op.Params, curToken)
    // fmt.Println(">", op.Params)
    curToken = ""
  }

  # append cur token to the groups list
  action group {
    // fmt.Println("group", curToken)
    op.Groups = append(op.Groups, curToken)
    // fmt.Println(">", op.Groups)
    curToken = ""
  }

  # append the currently built up op to the list of ops
  action opSpec {
    // fmt.Println("opSpec", op)
    ops = append(ops, op)
    op = Op{}
  }

  # for error book keeping we track the current line
  # and the pointer of where the last line started
  action line {
    // fmt.Println("line")
    line++
    lastLine = p
  }

  # set the current deps operation to be the current token
  action depOperation {
    // fmt.Println("dep_operation", curToken)
    dep.Operation = curToken
    curToken = ""
  }

  # append the current token to the dep params
  action depParam {
    // fmt.Println("dep_param", curToken)
    dep.Params = append(dep.Params, curToken)
    curToken = ""
  }

  # push the dep
  action dep {
    // fmt.Println("dep", dep)
    op.Depends = append(op.Depends, dep)
    dep = Dep{}
  }

  # parameters are alnumbers and some path possiblities,
  # may have to expand this to everything except )
  # on every ($) one of these, call char action
  p = (alnum | '/' | '.' | '~') $char;

  # param is one or me Ps. When we finish this state (%)
  # call param action
  param = p+ %param;

  # params are surrounded by (), 
  # and seperated by commas with an optional trailing space
  # lists are usually done with: zero or more items with sepeartors, with final item
  params = '(' (param ',' ' '?)* param ')';

  # group is alphanumeric repeating, call action group at the end
  group = (alnum $char)+ %group;

  #groups are list of comma seperated group-s
  groups = (group ',' ' '?)* group;


  operationName = (alnum $char)+ %op;
  operationDecl = operationName params ':' ' '?;

  depParam = p+ %depParam;
  depParams = '(' (depParam ',' ' '?)* depParam ')';
  depOperation = (alnum $char)+ %depOperation;
  dep = depOperation depParams %dep;
  depsDecl = ' '? '>' ' '? (dep ',' ' '?)* dep;

  operationSpec = operationDecl groups depsDecl? %opSpec;

  main := (operationSpec '\n' %line)* operationSpec;
}%%

%% write data;

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
  %% write init;

  // output the code for operating the parser
  %% write exec;

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
