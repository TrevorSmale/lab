Project Outline

MP3 Playback



1. Reading the MP3 file from some storage (like an SD card or flash memory).

2. Decoding the MP3 file.

3. Sending the decoded audio data to the I2S interface.


Here's a high-level outline of how you might approach this:


1. **Setup**:

   - Initialize the I2S interface.

   - Initialize the storage interface (e.g., SD card).


2. **Reading the MP3 File**:

   - Open the MP3 file from the storage.

   - Read chunks of the MP3 file into a buffer.


3. **Decoding the MP3 File**:

   - Use an MP3 decoder library to decode the buffered MP3 data. This is the tricky part because as of my last update in September 2021, there isn't a widely-used MP3 decoder written in Go. You might need to interface with a C library or find/write a Go-based decoder that's efficient enough for your microcontroller.


4. **Sending Audio Data to I2S**:

   - Send the decoded audio data to the I2S interface.


