; echo -- send what is received
020$
   B.2    ; dest buf = receive
   G      ; clear dest
   +I5..1 ; read 1 byte; skip 1 if received
   I5     ; skip to end
   B2.1   ; src buf = receive, dest buf = transmit
   X      ; src buf -> dest buf
   +I     ; send transmit buf
   G      ; clear dest
   I-8    ; repeat
