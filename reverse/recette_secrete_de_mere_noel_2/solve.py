# Binary compiled with the following commands:

# g++ ./miss_claus_secret_recipe_2.cpp -s -o ./recipe2.exe && ./recipe2.exe flag{le_pole_nord_n_est_pas_trop_fort_en_crypto} > secret_secret_recipe.log
    # on windows.

# g++ ./miss_claus_secret_recipe_2.cpp -s -o ./recipe2 && ./recipe2 flag{le_pole_nord_n_est_pas_trop_fort_en_crypto} > secret_secret_recipe.log
    # on linux.

# Get the second iteration of the recipe in the logs (the first one is a fake!)
out="""* 56 g de farine.
* 81 g de sucre.
* 53 oeufs.
* 83 mL de lait.
* 45 cuilleres de vanille.
* 70 mL d'extrait d'arc-en-ciel.
* 36 g de cacao.
* 124 mL d'huile vegetale.
* 9 mg de sel.
* 111 mg de bananes ecrasees.
* 4 mL de jus de pomme.
* 102 mg d'arachides' (sauf en cas d'allergie).
* 62 pepites au chocolat.
* 85 mL de creme fraiche.
* 59 mg de romarin seche.
* 78 mL de jus de citron.
* 43 flocons de neige.
* 113 mg de gentillesse.
* 40 mL de sirop d'erable.
* 128 g de farine.
* 226 g de sucre.
* 150 oeufs.
* 227 mL de lait.
* 185 cuilleres de vanille.
* 206 mL d'extrait d'arc-en-ciel.
* 184 g de cacao.
* 204 mL d'huile vegetale.
* 148 mg de sel.
* 229 mg de bananes ecrasees.
* 160 mL de jus de pomme.
* 216 mg d'arachides' (sauf en cas d'allergie).
* 173 pepites au chocolat.
* 243 mL de creme fraiche.
* 146 mg de romarin seche.
* 250 mL de jus de citron.
* 141 flocons de neige.
* 254 mg de gentillesse.
* 166 mL de sirop d'erable.
* 196 g de farine.
* 171 g de sucre.
* 241 oeufs.
* 147 mL de lait.
* 230 cuilleres de vanille.
* 168 mL d'extrait d'arc-en-ciel.
* 221 g de cacao.
* 174 mL d'huile vegetale.
* 198 mg de sel.
* 188 mg de bananes ecrasees.
Fini les 42 biscuits!"""

out_number_list = [56, 81, 53, 83, 45, 70, 36, 124, 9, 111, 4, 102, 62, 85, 59, 78, 43, 113, 40, 128, 226, 150, 227, 185, 206, 184, 204, 148, 229, 160, 216, 173, 243, 146, 250, 141, 254, 166, 196, 171, 241, 147, 230, 168, 221, 174, 198, 188]

# Formula must be reversed from binary
def goback(inp, iv):
    return (((inp - 1) ^ 0b110) - 2) ^ iv

iv = 73 # Printed out as the amount of fake "cookies"
decoded = []
for c in out_number_list:
    res = goback(c, iv)
    decoded.append(chr(res))
    iv = c

print("".join(decoded))
