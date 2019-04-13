package units
func PToK(p Pound) Kg{
	return Kg(p*0.45)
}
func KToP(k Kg) Pound{return Pound(k/0.45)}