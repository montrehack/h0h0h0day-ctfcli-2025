package xmazegift

import "strings"

const (
	FLAG_0     = "FLAG_0_w4th_a_S0NnY_D4Y"
	FLAG_1     = "FLAG_1_c0ll1sion_f4ilure"
	FLAG_2     = "FLAG_2_n4me_is_4_B4D_s33d"
	FLAG_3     = "FLAG_3_i_h4te_SQL_iNj3CtIon"
	FLAG_4     = "FLAG_4_k0nam1_CODe_wh4t_4_suRpRisE"
	EASTER_EGG = "wwssadadba"
)

func checkEasterEgg(moves string) bool {
	return strings.Contains(moves, EASTER_EGG)
}
