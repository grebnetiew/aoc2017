b = 108400;
c = 125400;
h = 0;
While[True,
  If[! PrimeQ[b], h++];
  If[b == c, Break[]];
   b += 17];
h