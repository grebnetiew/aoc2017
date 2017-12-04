DuplicateFreeQ /@
  ReadList["..\\Projects\\aoc2017\\04\\puzzle04.in", Word, RecordLists -> True] //
  Count[#, True]&


DuplicateFreeQ[Sort /@ Characters /@ #] & /@
  ReadList["..\\Projects\\aoc2017\\04\\puzzle04.in", Word, RecordLists -> True] //
  Count[#, True]&
