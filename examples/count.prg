; count-10 -- count to 10
020$
   *N.1       ; #0 = 1
   +Q         ; #0 -> dest buf
   +I1        ; send transmit buf
   G          ; clear dest buf
   *O..1      ; #0 += 1
   +A.1.10.1  ; end if #0 > 10
   I-5        ; repeat
