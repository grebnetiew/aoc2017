l = ReadList[NotebookDirectory[]<>"puzzle05.in", Number];
For[it = 0; cursor = 1, 1 <= cursor <= Length[l], cursor += l[[cursor]]++, it++];
it


l = ReadList[NotebookDirectory[]<>"puzzle05.in", Number];
it = 0; cursor = 1;
While[1 <= cursor <= Length[l],
  oldval = l[[cursor]];
  If[oldval < 3, l[[cursor]]++, l[[cursor]]--];
  cursor += oldval;
  it++;
];
it
