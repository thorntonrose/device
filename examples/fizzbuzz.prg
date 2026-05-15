; NOT COMPLETED
; fizzbuzz -- multiple of 3 => fizz, 5 => buzz, 3 & 5 => fizzbuzz
004=fizz
005=buzz

; main
020$
   B.2         ; dest buf = receive
   G           ; clear dest
   +I5..0.2    ; read up to 2 bytes
   *Q          ; #0 = dest buf
   *N#1.1      ; #1 = 1
   B1          ; dest buf = transmit
   G           ; clear dest buf
   *L21        ; call send
   +A#1..#0.2  ; skip 2 if #1 == #0
   *O#1.1.1    ; #1 += 1
   I-4         ; repeat

; send
021$
   *N#2.#1     ; #2 = #1
   *O#2.4.15.6 ; #1 %= 15; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.3.6  ; #1 %= 3; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.5.6  ; #1 %= 5; skip 6 if 0
   *L22        ; send number
   *M          ; return
   *L23        ; send fizzbuzz
   *M          ; return
   *L24        ; send fizz
   *M          ; return
   *L25        ; send buzz

; send number
022$
   +Q#1.1 ; #1 -> transmit buf
   +I     ; send

; send fizzbuzz
023$
   A4 ; 004 -> transmit buf
   A5 ; 005 -> transmit buf
   +I ; send transmit buf

; send fizz
024$
   A4 ; 004 -> transmit buf
   +I ; send transmit buf

; send buzz
025$
   A5 ; 005 -> transmit buf
   +I ; send transmit buf
