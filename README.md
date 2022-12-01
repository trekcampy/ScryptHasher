# ScryptHasher
This simple tool generates a 64 byte hashing using Scrypt key derivation function. It takes a pass phrase and an optional 8 byte salt. Without salt, the tool generates a random salt.

## Build Notes
go build -o ScryptHasher

## Usage
```
ScryptHasher -passphrase <string> -salt <hex string>
Where:  
  -passphrase string 
        pass phrase used in generating hash (default "passphrase")  
  -salt string 
        8 byte salt in hex format     
```
