ReadList["..\\Projects\\aoc2017\\02\\puzzle02.in", Number, RecordLists -> True] //
  Max[#]-Min[#]& /@# & //
  Total  


ReadList["..\\Projects\\aoc2017\\02\\puzzle02.in", Number, RecordLists -> True] //
  Function[{row},
    Permutations[row, {2}] //
    If[Mod[#1, #2]==0, #1/#2, 0]& @@# & /@# & //
    Total
  ] /@# & // 
  Total
