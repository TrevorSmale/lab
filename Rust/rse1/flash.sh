#!/bin/sh
set -e

elf2uf2-rs target/thumbv6m-none-eabi/release/rse1 target/thumbv6m-none-eabi/release/rse1.uf2
cp target/thumbv6m-none-eabi/release/rse1.uf2 /run/media/ts/RPI-RP2/
