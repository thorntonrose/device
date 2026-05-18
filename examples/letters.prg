; letters -- send only letters from received data

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
   +A.2.65.8  ; skip 8 if #0 < 65 (A)
   +A.1.90.1  ; skip 1 if #0 > 90 (Z)
   I2         ; skip 2
   +A.2.97.5  ; skip 5 if #0 < 97 (a)
   +A.1.122.4 ; skip 4 if #0 > 122 (z)
   B2.1       ; src = receive, dest = transmit
   G          ; clear dest
   X          ; src -> dest
   +I         ; send transmit buf
