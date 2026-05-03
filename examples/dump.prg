; dump -- dump memory
003=HELLO

; dump all
020$
   G    ; clear dest
   Y2   ; append next non-empty slot to dest; skip 2 if none
   V    ; display dest
   I-3  ; repeat

; dump buffers
021$
   P2   ; display transmit
   P3   ; display receive
