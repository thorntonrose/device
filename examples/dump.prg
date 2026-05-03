; dump -- dump memory
003=HELLO

; dump all
020$
  G    ; clear destination buffer
  Y2   ; append next non-empty location to destination buffer; skip 2 if none
  V    ; display destination buffer
  I-3  ; go back to G

; dump buffers
021$
  P2   ; display transmit buffer
  P3   ; display receive buffer
