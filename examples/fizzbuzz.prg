; fizzbuzz -- multiple of 3 => fizz, 5 => buzz, 3 & 5 => fizzbuzz
004=fizz
005=buzz
006=fizzbuzz

; main
020$
   B.2         ; dest buf = receive
   +I5..0.2    ; read up to 2 bytes
   *Q          ; #0 = dest buf
   *N#1.1      ; #1 = 1
   B.1         ; dest buf = transmit
   G           ; clear dest buf
   *L21        ; call calc
   +I1         ; send transmit buf
   +A#1..#0.2  ; skip 2 if #1 == #0
   *O#1..1     ; #1 += 1
   I-5         ; repeat

; calc
021$
   *N#2.#1     ; #2 = #1
   *O#2.4.15.6 ; #1 %= 15; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.3.6  ; #1 %= 3; skip 6 if 0
   *N#2.#1     ; #2 = #1
   *O#2.4.5.6  ; #1 %= 5; skip 6 if 0
   +Q#1        ; #1 -> dest buf
   *M          ; return
   A6          ; 006 -> dest buf
   *M          ; return
   A4          ; 004 -> dest buf
   *M          ; return
   A5          ; 005 -> dest buf
