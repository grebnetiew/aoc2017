StringSplit[ReadList[NotebookDirectory[]<>"\\puzzle01.in",String][[1]],""] //
  Transpose[{#, Join[Rest[#], {First[#]}]}] & //
  If[#1==#2,ToExpression[#1],0] & @@# & /@# & //
  Total


StringSplit[ReadList[NotebookDirectory[]<>"\\puzzle01.in",String][[1]],""] //
  Transpose[{#, Join[
    #[[Length[#]/2+1 ;; -1]],
    #[[1 ;; Length[#]/2]]]
  }] & //
  If[#1==#2, ToExpression[#1], 0]& @@# & /@# & //
  Total
