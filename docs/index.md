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


## Variables

Scripts have 10 variables named #0 to #9. They can hold integer or character values and can be referenced in any command that accepts parameters. Variables are initialized to 0.

**Examples:**

```
020$
  *N#1.1   ; #1 = 1
  *N#2.'A' ; #2 = 'A'
  B#1      ; src buf = #1 (1)
```


## Commands

...


## Memory

Device has a small amount of memory for data and programs.

**Memory Layout:**

Slot Number | Size (Bytes) | Description         |
------------|--------------|---------------------|
000 - 001   | -            | reserved            |
002         | 250          | buffer 1 (transmit) |
003         | 250          | buffer 2 (receive)  |
004 - 019   | -            | reserved            |
020 - 039   | 120          | general purpose     |


## Buffers

Buffers are special memory slots for data manipulation and I/O. Device has two buffers: buffer 1 (transmit) and buffer 2 (receive). Additionally, Device has two virtual buffers for viewing or manipulating the contents of a physical buffer: source (read) and destination (write).

Each virtual buffer has a pointer that indicates the current read/write position in the buffer. The pointer of the source buffer can be moved with various commands. The pointer of the destination buffer is always at the end of the data, and moves automatically as data is written to the buffer.

**Example:**

```
; preload receive buffer
003=HELLO WORLD

020$
   B2 ; set source buffer to buffer 2 (receive); set pointer to 0
   O6 ; move source pointer forward 6 positions
   X  ; copy "WORLD" from source buffer to destination buffer (transmit)
```
