# Cypher Tool

## Description

Cypher Tool is a command-line program that encrypts and decrypts single-line messages using simple cyphers.

## Usage
1. Run the program.
```
go run cyphertool.go
```
2. Select the operation (encrypt or decrypt).
```
Select operation (1/2):
1. Encrypt.
2. Decrypt.

Enter your choice: 
>>1
```
3. Select the encryption method.
```
Select cypher (1/3):
1. ROT13.
2. Reverse.
3. Custom.

Enter your choice:
>>1
```
4. Input message.
```
Enter the message: 
>>Hello, world!
```
5. Message is output
```
Encrypted message using shift:
Encrypted message using ROT13:
Uryyb, jbeyq!
```
6. Program exits

## Encryption Methods

### ROT13
A substitution cipher that replaces a letter with the 13th letter after it in the latin alphabet.

### Reverse
A substitution cipher that replaces a letter with it's reverse letter in the latin alphabet.

 ### Substitution
### Custom
A substitution cipher that replaces a letter with another letter according to predefined mapping.

 ## Authors

Joshua

Deji

Petteri