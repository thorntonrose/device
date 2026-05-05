# Device

Device is an emulator for a fictional I/O control device.


## Usage

```
device [<flags>] <file>
```

Run with `-h` for help.


## Memory

Device has a small amount of memory for data and programs. The memory layout is as follows:

Slot Number | Size (Bytes) | Description |
------------|--------------|------------ |
000 - 001   | -            | reserved
002         | 250          | buffer 1 (transmit)
003         | 250          | buffer 2 (receive)
004 - 019   | -            | reserved
020 - 039   | 120          | general purpose


## Buffers

Buffers are special memory slots for data manipulation and I/O. Device has two buffers: buffer 1 for transmit and buffer 2 for receive. Additionally, Device provides two virtual buffers for viewing or manipulating the content of a buffer: the source (read) buffer and the destination (write) buffer.

Each virtual buffer has a pointer that indicates the current position in the buffer. The pointer of the source buffer can be moved with various commands. The pointer of the destination buffer is always at the end of content, and moves automatically as data is written to the buffer.


## Programs

Programs are text files (usually with a .prg extension) containing sequences of directives for assigning data and scripts to memory slots. Scripts are sequences of commands for manipulating data and performing I/O operations.

### Syntax
