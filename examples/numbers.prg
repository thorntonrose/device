; numbers -- send only numbers from received data, one per line

; main
020$
   B.2       ; set dest = receive
   G         ; clear dest
   +I5..2    ; read 1 byte; skip 2 if received
   *L22      ; call send
   *M        ; return
   *Q..1     ; #0 = dest buf
   *L21      ; call save-digit
   +A#1..1.1 ; skip 1 if #1 == 1 (digit)
   *L22      ; call send
   I-9       ; repeat

; save-digit
021$
   *N#1      ; #1 = 0
   +A.2.48.5 ; skip 5 if #0 < 48 (0)
   +A.1.57.4 ; skip 4 if #0 > 57 (9)
   *N#1.1    ; #1 = 1
   B2.1      ; dest = transmit
   X         ; src -> dest
   *O#2..1   ; #2 += 1

; send
022$
   +A#2...4 ; skip 4 if #2 == 0
   +I1      ; send transmit buf
   B.1      ; dest = transmit
   G        ; clear dest
   *N#2     ; #2 = 0
