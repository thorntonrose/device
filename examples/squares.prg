; squares -- print square of each number received, delimited by spaces

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
   +A..32.4  ; skip 4 if #0 = 32 (space)
   *N#1.1    ; #1 = 1
   B2.1      ; dest = transmit
   X         ; src -> dest
   *O#2..1   ; #2 += 1

; send
022$
   +A#2...4  ; skip 4 if #2 == 0
   *L23      ; call square
   +I1       ; send transmit buf
   G         ; clear dest
   *N#2      ; #2 = 0

; square
023$
   B.1       ; dest = transmit
   *Q#3      ; #3 = dest buf
   *O#3.2.#3 ; #3 *= #3
   G         ; clear dest
   +Q#3      ; #3 -> dest
