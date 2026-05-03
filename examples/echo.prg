; echo -- send what is received
020$
   B.2      ; dest = receive
   G        ; clear dest
   B2.1     ; src = receive, dest = transmit
   +I5.1.1  ; wait for 1 character; skip 1 if received
   I5       ; skip to end
   X        ; src -> dest
   +I       ; send
   G        ; clear dest
   I-8      ; repeat
