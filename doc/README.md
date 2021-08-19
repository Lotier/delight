# DELIGHT Keyboard

[Build Guide \[To be updated\]]()

[VIA Configurator Guide (Easiest)\[To be updated\]]()

[QMK Toolbox Flashing Guide (not needed if using VIA)\[To be updated\]]()

[Command Line Flashing Guide (Advanced)\[To be updated\]]()

[Bootloader\[To be updated\]]()

[BOM (Parts List)\[To be updated\]]()

- Upload Bootloader ROM to chip 
  - Bootloader sources differs so we use a atmega32a chip, but otherwises it's just a standard atmega32 chip
- Hit chip with fuse settings so its set to the 16mhz timing crystal.
  - Find exact command line arguement should be (`-U lfuse:w:0x7f:m -U hfuse:w:0xd9:m` or `-U lfuse:w:0x1f:m -U hfuse:w:0xc0:m`) I think
- Reboot chip
- In QMK Toolbox upload firmware.