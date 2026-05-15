; echo2 -- send what is received (using subroutines)
020$
   *L21    ; call receive
   +A..4.2 ; skip 2 if #0 == EOT
   *L22    ; call send
   I-3     ; repeat

; receive
021$
   B.2    ; dest buf = receive
   G      ; clear dest
   +I5..1 ; read 1 byte; skip 1 if received
   *N.4   ; append 4 (ASCII EOT)

; send
022$
   B2.1 ; src buf = receive, dest buf = transmit
   X    ; src buf -> dest buf
   +I   ; send transmit buf
   G    ; clear dest
