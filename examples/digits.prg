; digits -- send only digits from received data

; main
020$
   B.2    ; set dest = receive
   G      ; clear dest
   +I5..1 ; read 1 byte; skip 1 if received
   *M     ; return
   *Q..1  ; #0 = dest buf (as ASCII)
   *L21   ; call send
   I-6    ; repeat

; send
021$
   +A.2.48.5 ; skip 5 if #0 < 48 (0)
   +A.1.57.4 ; skip 4 if #0 > 57 (9)
   B2.1      ; src = receive, dest = transmit
   G         ; clear dest
   X         ; src -> dest
   +I        ; send transmit buf
