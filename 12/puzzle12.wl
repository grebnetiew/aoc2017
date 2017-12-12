components = (
  ToExpression/@StringSplit[#, {", ", " <-> "}]& /@ ReadList[NotebookDirectory[]<>"puzzle12.in", Record] // 
  Function[{l}, UndirectedEdge[First[l], #]& /@ Rest[l]] /@# & // 
  Flatten // Graph // ConnectedComponents)

Length[Select[components,ContainsAll[#, {0}]&][[1]]]
Length[components]
