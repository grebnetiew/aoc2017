target = ReadList[NotebookDirectory[]<>"puzzle03.in", Number][[1]];
Clear[memory];
pos = 0;
direction = -I; (* South, turning East immediately *)
For[i = 1, i != target, i++,
  memory[pos] = i;
  (* If the cell to our left is empty, turn left *)
  If[!NumberQ[memory[pos + direction I]], direction *= I];
  pos += direction;
];
Abs@Re[pos] + Abs@Im[pos]


Clear[memory];
memory[_] := 0; (* Default value *)pos = 0;
direction = -I; (* South, turning East immediately *)
For[i = 1, i != target, i++,
  memory[pos] = memory[pos+1-I] + memory[pos+1] + memory[pos+1+I] +
                memory[pos-I]                   + memory[pos+I] + 
                memory[pos-1-I] + memory[pos-1] + memory[pos-1+I];
  (* If the cell to our left is empty, turn left *)
  If[NumberQ[memory[pos + direction I]] == 0, direction *= I];
  pos += direction;
];
i
