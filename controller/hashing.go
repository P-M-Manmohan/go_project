package controller

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

func GenerateRandomSalt(size int) ([]byte, error){
    salt:= make([]byte, size)
    _, err:= rand.Read(salt)
    if err!=nil {
        return nil, err
    }
    return salt, nil
}

func HashPassword (password string) (string, string, error){
    salt, err := GenerateRandomSalt(16)
    if err!= nil{
        return "","",err
    }

    time := uint32(1)
    memory := uint32(64 * 1024)
    threads := uint8(2)
    keyLen := uint32(32)

    hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

    saltBase64 := base64.StdEncoding.EncodeToString(salt)
    hashBase64 := base64.StdEncoding.EncodeToString(hash)

    return saltBase64, hashBase64, nil
}

func VerifyPassword(password, storedHash, storedSalt string) bool {
    salt, err := base64.StdEncoding.DecodeString(storedSalt)
    if err != nil {
        log.Println("failed to decode salt")
        return false
    }
    
    time := uint32(1)
    memory := uint32(64 * 1024)
    threads := uint8(2)
    keyLen := uint32(32)

    hash := argon2.IDKey([]byte(password), salt, time, memory, threads , keyLen)

    hashEncoded := base64.StdEncoding.EncodeToString(hash)

    return hashEncoded == storedHash

}
