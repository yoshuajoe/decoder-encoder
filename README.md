# Encoder - Decoder
# Encode
`Encode()` will transform `uint32(2500)` into bunches of array of bytes.
## Method
Simply by using `bit shift`, then using AND of `0xFF` to wrap up the corresponding bits. 

# Decode
`Decode` will transform `[]bytes(...)` into `uint32`.
## Method
Simply by using 
```golang
for i = 1; i < unsafe.Sizeof(uint32(0)); i++ {
	result += uint32(bytes[i]) * 256 * uint32(i)
}
``` 
