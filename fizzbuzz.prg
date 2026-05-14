; fizzbuzz -- multiple of 3 => fizz, 5 => buzz, 3 & 5 => fizzbuzz
004=fizz
005=buzz

020$
   B.2         ; dest buf = receive
   G           ; clear dest
   +I5..0.2    ; read up to 2 bytes
   B2          ; dest buf = transmit
   *Q.2        ; #0 = src buf
   *N#1.1      ; #1 = 1
   G           ; clear dest buf
   ; switch
   *N#2.#1     ; #2 = #1
   *O#2.4.15.6 ; #1 %= 15; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.3.6  ; #1 %= 3; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.5.6  ; #1 %= 5; skip 6 if 0
   *L21        ; send number
   I5          ; skip 5
   *L22        ; send fizzbuzz
   I3          ; skip 3
   *L23        ; send fizz
   I1          ; skip 1
   *L24        ; send buzz
   ; count
   *O#1.1.1    ; #1 += 1
   +A#1.1.#0.1 ; skip 1 if #1 > #0
   I-16        ; repeat

; send number
021$


; send
022$
   A4 ; 004 -> transmit buf
   A5 ; 005 -> transmit buf
   +I ; send transmit buf
