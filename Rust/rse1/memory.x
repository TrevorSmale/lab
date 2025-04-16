/* memory.x - RP2040 Memory Layout */

MEMORY
{
  /* 264KB SRAM (split in two banks in hardware, but seen as continuous here) */
  FLASH : ORIGIN = 0x10000000, LENGTH = 2M
  RAM   : ORIGIN = 0x20000000, LENGTH = 264K
}

/* Provide default stack top (start of RAM) */
_stack_start = ORIGIN(RAM) + LENGTH(RAM);
