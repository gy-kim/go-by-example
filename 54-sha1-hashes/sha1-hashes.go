package main

/*

https://gobyexample.com/sha1-hashes

SHA1 hashes are frequently used to compute short
identities for binary or text blobs. For example, the git
revision control system uses SHA1s extensively to identify
versioned files and directories. Here's how to compute
SHA1 hashes in Go.

Go implements serveral hash functions in various crypto/*
packages.

*/

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	/*
		The pattern for generating a hash is sha1.New(),
		sha1.Write(bytes), then sha1.Sum([]byte{}). Here we
		start with new hash.
	*/
	h := sha1.New()

	/*
		Write expects bytes. If you have a string s, use []byte(s)
		to coerce it ti bytes.
	*/
	h.Write([]byte(s))

	/*
		This gets the finalized hash result as a byte slice. The
		argument to Sum can be used to append to an existing byte
		slice: it usually isn't needed.
	*/
	bs := h.Sum(nil)

	/*
		SHA1 values are often printed in hex, for example in git
		commits. Use the %x format verb to convert a hash results
		to a hex string.
	*/
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	/*
		Running the program computes the hash and prints it in a
		human-readable hex format.

		You can compute other hashes using a similar pattern to
		the one shown above. For example, to compute MD5
		hashes import crypto/md5 and use md5.New().

		Note that if you need cryptographically secure  hashes, you
		should carefully research hash strenth.
	*/
}
