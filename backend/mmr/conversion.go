package mmr

func MapTrueSkillToMMR(mu float64, sigma float64) float64 {
	return mu * 60
	//// Define constants for TrueSkill to MMR mapping
	//trueSkillMin := 0.0  // Minimum possible TrueSkill value (adjust as needed)
	//trueSkillMax := 50.0 // Maximum possible TrueSkill value (adjust as needed)
	//mmrMin := 1000.0     // Minimum MMR value
	//mmrMax := 2000.0     // Maximum MMR value
	//
	//// Calculate the conservative estimate (mu - 3 * sigma)
	//trueSkillValue := mu - 3.0*sigma
	//
	//// Perform linear interpolation to map TrueSkill value to MMR
	//mmr := mmrMin + ((trueSkillValue-trueSkillMin)/(trueSkillMax-trueSkillMin))*(mmrMax-mmrMin)
	//
	//// Round MMR to the nearest integer
	//return mmr
}
