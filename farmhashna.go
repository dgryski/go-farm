package farm

func ShiftMix(val uint64) uint64 {
	return val ^ (val >> 47)
}

func HashLen16(u, v uint64) uint64 {
	return Hash128to64(uint128{u, v})
}

func HashLen16Mul(u, v, mul uint64) uint64 {
	// Murmur-inspired hashing.
	a := (u ^ v) * mul
	a ^= (a >> 47)
	b := (v ^ a) * mul
	b ^= (b >> 47)
	b *= mul
	return b
}

func HashLen0to16(s []byte) uint64 {
	slen := uint64(len(s))
	if slen >= 8 {
		mul := k2 + slen*2
		a := Fetch64(s, 0) + k2
		b := Fetch64(s, int(slen-8))
		c := Rotate64(b, 37)*mul + a
		d := (Rotate64(a, 25) + b) * mul
		return HashLen16Mul(c, d, mul)
	}

	if slen >= 4 {
		mul := k2 + slen*2
		a := Fetch32(s, 0)
		return HashLen16Mul(uint64(slen)+(uint64(a)<<3), uint64(Fetch32(s, int(slen-4))), mul)
	}
	if slen > 0 {
		a := s[0]
		b := s[slen>>1]
		c := s[slen-1]
		y := uint32(a) + (uint32(b) << 8)
		z := uint32(slen) + (uint32(c) << 2)
		return ShiftMix(uint64(y)*k2^uint64(z)*k0) * k2
	}
	return k2
}

// This probably works well for 16-byte strings as well, but it may be overkill
// in that case.
func HashLen17to32(s []byte) uint64 {
	slen := len(s)
	mul := k2 + uint64(slen*2)
	a := Fetch64(s, 0) * k1
	b := Fetch64(s, 8)
	c := Fetch64(s, slen-8) * mul
	d := Fetch64(s, slen-16) * k2
	return HashLen16Mul(Rotate64(a+b, 43)+Rotate64(c, 30)+d, a+Rotate64(b+k2, 18)+c, mul)
}

// Return a 16-byte hash for 48 bytes.  Quick and dirty.
// Callers do best to use "random-looking" values for a and b.
func WeakHashLen32WithSeedsWords(w, x, y, z, a, b uint64) (uint64, uint64) {
	a += w
	b = Rotate64(b+a+z, 21)
	c := a
	a += x
	a += y
	b += Rotate64(a, 44)
	return a + z, b + c
}

// Return a 16-byte hash for s[0] ... s[31], a, and b.  Quick and dirty.
func WeakHashLen32WithSeeds(s []byte, a, b uint64) (uint64, uint64) {
	return WeakHashLen32WithSeedsWords(Fetch64(s, 0),
		Fetch64(s, 8),
		Fetch64(s, 16),
		Fetch64(s, 24),
		a,
		b)
}

// Return an 8-byte hash for 33 to 64 bytes.
func HashLen33to64(s []byte) uint64 {
	slen := len(s)
	mul := k2 + uint64(slen)*2
	a := Fetch64(s, 0) * k2
	b := Fetch64(s, 8)
	c := Fetch64(s, slen-8) * mul
	d := Fetch64(s, slen-16) * k2
	y := Rotate64(a+b, 43) + Rotate64(c, 30) + d
	z := HashLen16Mul(y, a+Rotate64(b+k2, 18)+c, mul)
	e := Fetch64(s, 16) * mul
	f := Fetch64(s, 24)
	g := (y + Fetch64(s, slen-32)) * mul
	h := (z + Fetch64(s, slen-24)) * mul
	return HashLen16Mul(Rotate64(e+f, 43)+Rotate64(g, 30)+h, e+Rotate64(f+a, 18)+g, mul)
}

func Hash64(s []byte) uint64 {
	slen := len(s)
	const seed uint64 = 81
	if slen <= 32 {
		if slen <= 16 {
			return HashLen0to16(s)
		} else {
			return HashLen17to32(s)
		}
	} else if slen <= 64 {
		return HashLen33to64(s)
	}

	// For strings over 64 bytes we loop.  Internal state consists of
	// 56 bytes: v, w, x, y, and z.
	x := seed
	y := uint64(2480279821605975764) // == seed * k1 + 113; This overflows uint64 and is a compile error, so we expand the constant by hand
	z := ShiftMix(y*k2+113) * k2
	var v1, v2 uint64
	var w1, w2 uint64
	x = x*k2 + Fetch64(s, 0)

	// Set end so that after the loop we have 1 to 64 bytes left to process.
	endIdx := ((slen - 1) / 64) * 64
	last64Idx := endIdx + ((slen - 1) & 63) - 63
	last64 := s[last64Idx:]
	for len(s) > 64 {
		x = Rotate64(x+y+v1+Fetch64(s, 8), 37) * k1
		y = Rotate64(y+v2+Fetch64(s, 48), 42) * k1
		x ^= w2
		y += v1 + Fetch64(s, 40)
		z = Rotate64(z+w1, 33) * k1
		v1, v2 = WeakHashLen32WithSeeds(s, v2*k1, x+w1)
		w1, w2 = WeakHashLen32WithSeeds(s[32:], z+w2, y+Fetch64(s, 16))
		z, x = x, z
		s = s[64:]
	}
	mul := k1 + ((z & 0xff) << 1)

	// Make s point to the last 64 bytes of input.
	s = last64
	w1 += ((uint64(slen) - 1) & 63)
	v1 += w1
	w1 += v1
	x = Rotate64(x+y+v1+Fetch64(s, 8), 37) * mul
	y = Rotate64(y+v2+Fetch64(s, 48), 42) * mul
	x ^= w2 * 9
	y += v1*9 + Fetch64(s, 40)
	z = Rotate64(z+w1, 33) * mul
	v1, v2 = WeakHashLen32WithSeeds(s, v2*mul, x+w1)
	w1, w2 = WeakHashLen32WithSeeds(s[32:], z+w2, y+Fetch64(s, 16))
	z, x = x, z
	return HashLen16Mul(HashLen16Mul(v1, w1, mul)+ShiftMix(y)*k0+z,
		HashLen16Mul(v2, w2, mul)+x,
		mul)
}

func Hash64WithSeed(s []byte, seed uint64) uint64 {
	return Hash64WithSeeds(s, k2, seed)
}

func Hash64WithSeeds(s []byte, seed0, seed1 uint64) uint64 {
	return HashLen16(Hash64(s)-seed0, seed1)
}
