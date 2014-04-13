package tgf

import (
    "os"
    "io"
    "bufio"
    "strings"
    "code.google.com/p/go-uuid/uuid"
)

func Parse(file io.Reader) (map[string]Node, map[string]Edge, error) {
  scannerState := 0
  nodes        := make(map[string]Node)
  edges        := make(map[string]Edge)
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
        node  := Node{id, label, make([]string, 0), make([]string, 0)}

        nodes[id] = node

      } else if scannerState == 1 {
        id         := uuid.New()
        fromNodeId := words[0]
        fromNode   := nodes[fromNodeId]
        toNodeId   := words[1]
        toNode     := nodes[toNodeId]
        label      := strings.Join(words[descriptionStartsAt:], " ")

        fromNode.ToEdgeIds = append(fromNode.ToEdgeIds, id)
        nodes[fromNodeId] = fromNode

        toNode.FromEdgeIds = append(toNode.FromEdgeIds, id)
        nodes[toNodeId] = toNode

        edge := Edge{id, label, fromNodeId, fromNode, toNodeId, toNode}
        edges[id] = edge
      }
    }
  }

  return nodes, edges, nil
}

func ParseFile(filename string) (map[string]Node, map[string]Edge, error) {
  file, err := os.Open(filename)
  if err != nil {
    return make(map[string]Node), make(map[string]Edge), err
  }
  defer file.Close()

  return Parse(file)
}