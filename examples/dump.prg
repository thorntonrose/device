; dump -- dump memory
003=HELLO

; dump all
020$
   G   ; clear dest
   Y2  ; next non-empty slot -> dest; skip 2 if none
   V   ; display dest
   I-3 ; repeat

; dump buffers
021$
   P2 ; display transmit buf
   P3 ; display receive buf
