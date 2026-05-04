; count-10 -- count from 1 to 10
020$
   *N.1      ; #0 = 1
   +Q        ; dest = #0
   +I1       ; send dest + \n
   G         ; clear dest
   *O..1     ; #0 += 1
   +A.1.10.1 ; end if #0 > 10
   I-5       ; repeat
