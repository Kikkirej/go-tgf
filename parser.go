package tgf

import (
    "os"
    "io"
    "bufio"
    "strings"
    "github.com/didip/gotgf/ast"
)

func Parse(file io.Reader) (map[string]ast.Node, map[string]ast.Edge, error) {
  scannerState := 0
  scanner      := bufio.NewScanner(file)

  for scanner.Scan() {
    line      := scanner.Text()
    firstChar := line[0]
    words     := strings.Fields(line)

    if firstChar == '#' {
      scannerState = 1
    } else {
      descriptionStartsAt := scannerState + 1

      if scannerState == 0 {
        id    := words[0]
        label := strings.Join(words[descriptionStartsAt:], " ")

        ast.MakeNode(id, label)

      } else if scannerState == 1 {
        InboundNodeId  := words[0]
        OutboundNodeId := words[1]
        label          := strings.Join(words[descriptionStartsAt:], " ")

        ast.MakeEdge(label, InboundNodeId, OutboundNodeId)
      }
    }
  }

  return ast.Nodes, ast.Edges, nil
}

func ParseFile(filename string) (map[string]ast.Node, map[string]ast.Edge, error) {
  file, err := os.Open(filename)
  if err != nil {
    return make(map[string]ast.Node), make(map[string]ast.Edge), err
  }
  defer file.Close()

  return Parse(file)
}