; echo-two -- output two of each byte of input

; main
020$
   B.2    ; dest buf = receive
   G      ; clear dest
   +I5..1 ; read 1 byte; skip 1 if received
   *M     ; return
   *L21   ; call double
   +I     ; send transmit buf
   G      ; clear dest
   I-7    ; repeat

; double
021$
   B2.1   ; src buf = receive, dest buf = transmit
   X      ; src buf -> dest buf
   O-1    ; src pointer -= 1
   X      ; src buf -> dest buf
