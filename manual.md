# Device

Device is an emulator for a fictional I/O control device.


## Usage

```
device [<flags>] <file>
```

Run with `-h` for help.


## Programs

Programs are text files containing sequences of directives for assigning data and scripts to memory slots. Scripts are sequences of commands. Comments are allowed and are denoted with ";".

### Syntax

```
<program> ::= <statement>+
<statement> ::= <comment> | <data-directive> | <script-directive>

<comment> ::= ';'[<text>]<eol>

<data-directive> ::= <slot>'='<text>(<comment>)<eol>
<slot> ::= <digit>+

<script-directive> ::= <slot>'$'<script>
<script> ::= (<space>*<command><space>*[<comment>]<eol>)+
<command> ::= ['+' | '*']<'A'...'Z'>[<parameters>]
<parameters> ::= <parameter>('.'<parameter>)*
<parameter> ::= <variable> | ['-']<integer> | '<text>'
<variable> ::= '#'<digit>

<eol> ::= <newline> | <end>
```

### Example

```
; (data directive)
002=HELLO ; "HELLO" -> slot 2

; (script directive with one command)
020$+I1 ; send transmit buffer with newline

; (script directive with multiple commands)
021$
   G    ; clear dest
   Y2   ; next non-empty slot -> dest; skip 2 if none
   V    ; display dest
   I-3  ; repeat
```


## Memory

Device has a small amount of memory for data and programs.

**Memory Layout:**

Slot Number | Size (Bytes) | Description         |
------------|--------------|---------------------|
000 - 001   | -            | reserved            |
002         | 250          | buffer 1 (transmit) |
003         | 250          | buffer 2 (receive)  |
004 - 006   | -            | buffers 3-5         |
007 - 019   | -            | reserved            |
020 - 099   | 120          | general purpose     |
100 - 112   | -            | reserved            |
113 - 949   | 120          | general purpose     |
950 - 999   | -            | reserved            |


## Buffers

Buffers are special memory slots for data manipulation and I/O. Device has two buffers: buffer 1 (transmit) and buffer 2 (receive). Additionally, Device has two virtual buffers for viewing or manipulating the contents of a physical buffer: source (read) and destination (write).

Each virtual buffer has a pointer that indicates the current read/write position in the buffer. The pointer of the source buffer can be moved with various commands. The pointer of the destination buffer is always at the end of the data and moves automatically as data is written to the buffer.

**Example:**

```
; preload receive buffer
003=HELLO WORLD

020$
   B2 ; set source buffer to buffer 2 (receive); set pointer to 0
   O6 ; move source pointer forward 6 positions
   X  ; copy "WORLD" from source buffer to destination buffer (transmit)
```


## Variables

Programs have 10 variables named #0 to #9. They can hold integer or character values and can be referenced in any command that accepts parameters. Variables are initialized to 0 when a program starts.

**Examples:**

```
020$
  *N#1.1   ; #1 = 1
  *N#2.'A' ; #2 = 'A'
  B#1      ; src buf = #1 (1)
```


## Commands

### A[m.s]

Append contents of memory slot to destination buffer, moving pointer.

Parameters:

  - m: memory slot; default: 0
  - s: commands to skip if slot is empty; integer; default: 0

### +A[#v.a.c.s]

Compare variable.

Parameters:

  - #v: variable: #0 (default) - #9
  - a: comparison operation: 0 (default) = equal, 1 = greater than, 2 = less than
  - c: constant value (default: 0)
  - s: commands to skip if true: integer; default: 0

### B[b1.b2]

Set source and destination buffers, optionally resetting pointers.

Parameters:

  - b1: source buffer: 0 (default) = no change, 1 - \<max-buffers> = set buffer and reset pointer, 9 = reset pointer
  - b2: destination buffer: 0 (default) = no change, 1 - \<max-buffers> = set buffer and reset pointer

### D[n]

Delete characters from end of destination buffer, moving pointer.

Parameters:

  - n: number of characters to delete; integer; default: 1

### G

Clear destination buffer, resetting pointer.

### H[s.c]

Search for string in source buffer, moving pointer to the start of the string if found.

Parameters:

  - s: commands to skip if string not found; integer; default: 0
  - c: character or string to search for: 0 - 255 | '\<string>'; default: tab (ASCII 9)

### I[s.a.c]

Compare value to source buffer and skip.

Parameters:
  - n: commands to skip if comparison is true; integer; default: 0
  - a: comparison code: 0 (default) = true, 1 = equal, 2 = not equal, 3 = less than, 4 = greater than
  - c: constant value to compare: 0 - 255 | ‘\<string>’; default: tab (ASCII 9)

Note:
    - Skipping past the beginning or end of the script slot terminates the script.
    - If the source buffer starting from the pointer is empty, the skip always occurs.

### +I[a.t.s.n]

Send and receive data.

Parameters:
  - a: operation code: 0 (default) = send transmit buffer to stdout, 1 = send with newline, 5 = wait for n characters from stdin then append to receive buffer
  - t: reserved
  - s: number of commands to skip if data received: 0 (default)
  - n: number of characters to receive; positive integer; default: 1

### *L[m]

Call subroutine (script).

Parameters:

  - m: memory slot; default: 0
  - Note: To return from the subroutine, skip past the beginning or end of the script or run the *M command.

### *M

Return from subroutine (script).

### *N[#v.c]

Set variable to constant value.

Parameters:

  - #v: variable: #0 (default) - #9
  - c: constant value (default: 0)

### O[n]

Move source buffer pointer.

Parameters:

  - n: number of characters to move; integer; default: 1

### *O[#v.o.c.s]

Do math operation on variable.

Parameters:

  - #v: variable: #0 (default) - #9
  - o: operation: 0 (default) = add, 1 = subtract, 2 = multiply, 3 = divide, 4 = modulo
  - c: constant value (default: 0)
  - s: commands to skip if result is 0: integer; default: 0

### P[m]

Display contents of memory slot.

Parameters:

  - m: memory slot (default: 0)

### *Q[#v.b.a]

Set variable from buffer.

Parameters:
  - #v: variable: #0 (default) - #9
  - b: buffer to read: 0 (default) = destination buffer, 1 - \<max-buffers> = buffer
  - a: type of conversion: 0 (default) = string (“2748” -> 2748), 1 = ASCII (“V” -> 86, “XY” -> 22617 [0x5859])

Note: The buffer pointer is ignored.

### +Q[#v.a]

Append variable to destination buffer, moving pointer.

Parameters:

  - #v: variable: #0 (default) - #9
  - a: type of conversion: 0 (default) = string (2748 -> “2748”), 1 = ASCII (86 -> “V”, 22617 -> “XY”)

### R[c]

Append constant value to destination buffer, moving pointer.

Parameters:

  - c: constant value: 0 - 255 | ‘\<string>’; default: tab (ASCII 9)

### T[m.o.c.s]

Do math operation on contents of memory slot, treating empty slot as 0.

Parameters:

  - m: memory slot: 0 (default) - \<max-memory-slots>
  - o: operation: 0 (default) = add, 1 = subtract, 2 = multiply, 3 = divide, 4 = modulo
  - c: constant value (default: 0)
  - s: commands to skip if result is 0: integer; default: 0

### V[b]

Display contents of buffer.

Parameters:
  - b: buffer: 0 - \<max-buffers>, 0 (default) = destination buffer

### X[n.c]

Copy source buffer to destination buffer, moving pointer.

Parameters:

  - n: number of characters to copy: 0 (default) = copy to stop character
  - c: stop character: 0 (default) = end of buffer

### Y[s]

Append next non-empty memory slot to destination buffer, moving pointer.

Parameters:

  - s: commands to skip after all memory slots are read; integer; default: 0
