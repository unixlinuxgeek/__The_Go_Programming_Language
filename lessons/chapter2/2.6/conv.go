package tempconv

func CToF(c Cel) Far {
	return Far(c*9/5 + 32)
}

func FToC(f Far) Cel {
	return Cel((f - 32) * 5 / 9)
}

// C = K – 273.15
func KToC(k Kel) Cel {
	return Cel(k - 273.15)
}

// K = C + 273.15
func CToK(c Cel) Kel {
	return Kel(c + 273.15)
}
