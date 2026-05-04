; echo -- send what is received
020$
  B.2      ; dest buf = receive
  G        ; clear dest
  +I5.1.1  ; wait for 1 character; skip if received
  I5       ; skip to end
  B2.1     ; src buf = receive, dest buf = transmit
  X        ; src -> dest
  +I       ; send transmit buf
  G        ; clear dest
  I-8      ; skip to beginning
