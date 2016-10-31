// hash project main.go
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
)

func main() {
	fmt.Println("=============================================================")
	test()
	fmt.Println("=============================================================")
	testSha1()
	fmt.Println("=============================================================")
	testSha256()
	fmt.Println("=============================================================")
	testSha512()
	fmt.Println("=============================================================")

}

func test() {
	dk := pbkdf2([]byte("password"), []byte("salt"), 1, 64, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "867f70cf1ade02cff3752599a3a53dc4af34c7a669815ae5d513554e1c8cf252c02d470a285a0501bad999bfe943c08f050235d7d68b1da55e63f73b60a57fce")
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}

func testSha1() {
	dk := pbkdf2([]byte("password"), []byte("salt"), 1, 20, sha1.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "0c60c80f961f0e71f3a9b524af6012062fe037a6")
	dk = pbkdf2([]byte("password"), []byte("salt"), 2, 20, sha1.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "ea6c014dc72d6f8ccd1ed92ace1d41f0d8de8957")
	dk = pbkdf2([]byte("password"), []byte("salt"), 4096, 20, sha1.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "4b007901b765489abead49d926f721d065a429c1")
	//	dk = pbkdf2([]byte("password"), []byte("salt"), 16777216, 20, sha1.New)
	//	fmt.Println(hex.EncodeToString(dk),
	//		hex.EncodeToString(dk) == "eefe3d61cd4da4e4e9945b3d6ba2158c2634e984")
	dk = pbkdf2([]byte("pass\x00word"), []byte("sa\x00lt"), 4096, 16, sha1.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "56fa6aa75548099dcc37d7f03425e0c3")
}

func testSha512() {
	dk := pbkdf2([]byte("password"), []byte("salt"), 1, 64, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "867f70cf1ade02cff3752599a3a53dc4af34c7a669815ae5d513554e1c8cf252c02d470a285a0501bad999bfe943c08f050235d7d68b1da55e63f73b60a57fce")
	dk = pbkdf2([]byte("password"), []byte("salt"), 2, 64, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "e1d9c16aa681708a45f5c7c4e215ceb66e011a2e9f0040713f18aefdb866d53cf76cab2868a39b9f7840edce4fef5a82be67335c77a6068e04112754f27ccf4e")
	dk = pbkdf2([]byte("password"), []byte("salt"), 4096, 64, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "d197b1b33db0143e018b12f3d1d1479e6cdebdcc97c5c0f87f6902e072f457b5143f30602641b3d55cd335988cb36b84376060ecd532e039b742a239434af2d5")
	dk = pbkdf2([]byte("passwordPASSWORDpassword"), []byte("saltSALTsaltSALTsaltSALTsaltSALTsalt"), 4096, 64, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "8c0511f4c6e597c6ac6315d8f0362e225f3c501495ba23b868c005174dc4ee71115b59f9e60cd9532fa33e0f75aefe30225c583a186cd82bd4daea9724a3d3b8")

	dk = pbkdf2([]byte("passDATAb00AB7YxDTT"), []byte("saltKEYbcTcXHCBxtjD"), 1, 63, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "cbe6088ad4359af42e603c2a33760ef9d4017a7b2aad10af46f992c660a0b461ecb0dc2a79c2570941bea6a08d15d6887e79f32b132e1c134e9525eeddd744")

	dk = pbkdf2([]byte("passDATAb00AB7YxDTT"), []byte("saltKEYbcTcXHCBxtjD"), 100000, 65, sha512.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "accdcd8798ae5cd85804739015ef2a11e32591b7b7d16f76819b30b0d49d80e1abea6c9822b80a1fdfe421e26f5603eca8a47a64c9a004fb5af8229f762ff41f7c")

	//	dk = pbkdf2([]byte("passDATA"), []byte("saltKEYbc"), 16777216, 7, sha512.New)
	//	fmt.Println(hex.EncodeToString(dk),
	//		hex.EncodeToString(dk) == "ab96c76400d08b")

}

func testSha256() {
	dk := pbkdf2([]byte("password"), []byte("salt"), 1, 32, sha256.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "120fb6cffcf8b32c43e7225256c4f837a86548c92ccc35480805987cb70be17b")
	dk = pbkdf2([]byte("password"), []byte("salt"), 2, 32, sha256.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "ae4d0c95af6b46d32d0adff928f06dd02a303f8ef3c251dfd6e2d85a95474c43")
	dk = pbkdf2([]byte("password"), []byte("salt"), 4096, 32, sha256.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "c5e478d59288c841aa530db6845c4c8d962893a001ce4e11a4963873aa98134a")
	//	dk = pbkdf2([]byte("password"), []byte("salt"), 16777216, 32, sha256.New)
	//	fmt.Println(hex.EncodeToString(dk),
	//		hex.EncodeToString(dk) == "cf81c66fe8cfc04d1f31ecb65dab4089f7f179e89b3b0bcb17ad10e3ac6eba46")
	dk = pbkdf2([]byte("passwordPASSWORDpassword"), []byte("saltSALTsaltSALTsaltSALTsaltSALTsalt"), 4096, 40, sha256.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "348c89dbcbd32b2f32d814b8116e84cf2b17347ebc1800181c4e2a1fb8dd53e1c635518c7dac47e9")
	dk = pbkdf2([]byte("pass\x00word"), []byte("sa\x00lt"), 4096, 16, sha256.New)
	fmt.Println(hex.EncodeToString(dk),
		hex.EncodeToString(dk) == "89b69d0516f829893c696226650a8687")
}

func pbkdf2(pass []byte, salt []byte, count int, dkLen int, h func() hash.Hash) []byte {
	dk := make([]byte, dkLen)
	prf := hmac.New(h, pass)
	hLen := prf.Size()
	l := int(math.Ceil(float64(dkLen) / float64(hLen)))
	r := dkLen - (l-1)*hLen
	ibuf := make([]byte, 4)
	var u []byte
	for i := 0; i < l; i++ {
		size := hLen
		if i == l-1 {
			size = r
		}
		prf.Reset()
		prf.Write(salt)
		binary.BigEndian.PutUint32(ibuf, uint32(i+1))
		prf.Write(ibuf)
		u = prf.Sum(nil)
		for j := 0; j < size; j++ {
			dk[j+i*hLen] = u[j]
		}
		for j := 1; j < count; j++ {
			prf.Reset()
			prf.Write(u)
			u = prf.Sum(nil)
			for k := 0; k < size; k++ {
				dk[k+i*hLen] ^= u[k]
			}
		}
	}
	return dk

	//1. If dkLen > (2^32 - 1) * hLen, output "derived key too long" and
	//   stop.

	//2. Let l be the number of hLen-octet blocks in the derived key,
	//   rounding up, and let r be the number of octets in the last
	//   block:

	//             l = CEIL (dkLen / hLen) ,
	//             r = dkLen - (l - 1) * hLen .

	//   Here, CEIL (x) is the "ceiling" function, i.e. the smallest
	//   integer greater than, or equal to, x.

	//3. For each block of the derived key apply the function F defined
	//   below to the password P, the salt S, the iteration count c, and
	//   the block index to compute the block:

	//             T_1 = F (P, S, c, 1) ,
	//             T_2 = F (P, S, c, 2) ,
	//             ...
	//             T_l = F (P, S, c, l) ,

	//   where the function F is defined as the exclusive-or sum of the
	//   first c iterates of the underlying pseudorandom function PRF
	//   applied to the password P and the concatenation of the salt S
	//   and the block index i:

	//             F (P, S, c, i) = U_1 \xor U_2 \xor ... \xor U_c

	//   where

	//             U_1 = PRF (P, S || INT (i)) ,
	//             U_2 = PRF (P, U_1) ,
	//             ...
	//             U_c = PRF (P, U_{c-1}) .

	//   Here, INT (i) is a four-octet encoding of the integer i, most
	//   significant octet first.

	//4. Concatenate the blocks and extract the first dkLen octets to
	//   produce a derived key DK:

	//             DK = T_1 || T_2 ||  ...  || T_l<0..r-1>

	//5. Output the derived key DK.

	//Note. The construction of the function F follows a "belt-and-
	//suspenders" approach. The iterates U_i are computed recursively to
	//remove a degree of parallelism from an opponent; they are exclusive-
	//ored together to reduce concerns about the recursion degenerating
	//into a small set of values.

}
