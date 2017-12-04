DuplicateFreeQ /@
  ReadList[NotebookDirectory[]<>"\\puzzle04.in", Word, RecordLists -> True] //
  Count[#, True]&


DuplicateFreeQ[Sort /@ Characters /@ #] & /@
  ReadList[NotebookDirectory[]<>"puzzle04.in", Word, RecordLists -> True] //
  Count[#, True]&
