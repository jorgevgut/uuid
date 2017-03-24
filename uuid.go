package main

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// UUID is the returned struct for UUID
type UUID struct {
	String string
	Bytes  []byte
}

// NewUUID generates a new UUID
func NewUUID() (*UUID, error) {
	// Generate UUID v4
	// Algorithm according to RFC 4122
	/*
		4.4.  Algorithms for Creating a UUID from Truly Random or
		      Pseudo-Random Numbers

			   The version 4 UUID is meant for generating UUIDs from truly-random or
			   pseudo-random numbers.

			   The algorithm is as follows:

			   o  Set the two most significant bits (bits 6 and 7) of the
			      clock_seq_hi_and_reserved to zero and one, respectively.

			   o  Set the four most significant bits (bits 12 through 15) of the
			      time_hi_and_version field to the 4-bit version number from
			      Section 4.1.3.

			   o  Set all the other bits to randomly (or pseudo-randomly) chosen
			      values.
	*/
	raw := make([]byte, 16, 16) //UUID has a fixed length of 128 bit (16 bytes)
	// all bits are 0
	// set bits 6 and 7 to 0 and 1 (bit 6 is already 0)
	//	var n uint
	//var pos uint
	var block byte = 0x01
	var bit6 byte = 0xfb

	// config for bits 12,13,14,15 (time_hi_and_version) = 0100 in binary (0x4)
	// bits 9 - 16 are in raw[1]
	// we require these 0 bits
	// xxx0x00x -> this means value &= 11101001   (E9 in Hex)
	var thav0 byte = 0xe9

	// we require this 1 bit
	// xxxx1xxx -> this means value |= 00001000 ... (8 in Hex)
	var thav1 byte = 0x8

	// Read random numbers and fill in byte
	_, err := rand.Read(raw)
	if err != nil {
		return nil, err
	}
	//pos = 8
	raw[0] |= block
	raw[0] &= bit6
	raw[1] &= thav0
	raw[1] |= thav1

	var parts []string
	//add strings depending on octect and "-"
	parts = append(parts, hex.EncodeToString(raw[:4]))
	parts = append(parts, "-")
	parts = append(parts, hex.EncodeToString(raw[4:6]))
	parts = append(parts, "-")
	parts = append(parts, hex.EncodeToString(raw[6:8]))
	parts = append(parts, "-")
	parts = append(parts, hex.EncodeToString(raw[8:10]))
	parts = append(parts, "-")
	parts = append(parts, hex.EncodeToString(raw[10:16]))
	s := strings.Join(parts, "")

	return &UUID{Bytes: raw, String: s}, nil
}
