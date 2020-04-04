package tree

import (
	"encoding/base64"
	"testing"
)

func TestCensus(t *testing.T) {
	root := "0x3cebb9e1be69d489897db0f30cb5b106d12d27d25b21a9ff509a9ddf8ad18a04"
	proof := "0x000b00000000000000000000000000000000000000000000000000000000041f9d522f7e1301ca58127dcb83bc01311dfcbec9ddc30928b575586e5592de09096467620f8d363f9486b449a13d5884ff70eb652db626acf3a3e0bbbf28de63257462bfbb9bdca670d438daccd1ed881b45892b4100bfef967f0e9350ba98ef2e4ab2326fa3cb71a5d442c21ead9bb7e43a69cf27b26b7442765f262a1804c00dbcbc96a800a673048568a9fdad6bf7d28634099f2f0945a387daed6ede7b94262cb0b58aa6d04bf2e61a3468bd931a45039d3ce0976898d454c681e4a37e3113"
	claim := "HMqKwz4Vlw+hyKx697D8lr+efpFWbiPd6qWkqlCzczY="
	data, err := base64.StdEncoding.DecodeString(claim)
	if err != nil {
		t.Fatal(err)
	}
	valid, err := CheckProof(root, proof, data)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Errorf("proof is invalid, but should be valid")
	} else {
		t.Log("proof valid")
	}
}
