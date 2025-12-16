# Binary compiled with the following commands:

    # on windows.
# g++ ./miss_claus_secret_recipe.cpp -s -o ./recipe.exe && ./recipe.exe flag{l_ingredient_secret_est_} > secret_recipe.log

    # on linux.
# g++ ./miss_claus_secret_recipe.cpp -s -o ./recipe && ./recipe flag{l_ingredient_secret_est_} > secret_recipe.log

# Get the second iteration of the recipe in the logs (the first one is a fake!)
out="""* 26 g de farine.
* 40 g de sucre.
* 11 oeufs.
* 25 mL de lait.
* 93 cuilleres de vanille.
* 40 g de cacao.
* 1 mL d'huile vegetale.
* 35 mg de sel.
* 50 mg de bananes ecrasees.
* 25 mL de jus de pomme.
* 70 mg d'arachides (sauf en cas d'allergie).
* 31 pepites au chocolat.
* 16 mL de creme fraiche.
* 35 mg de romarin seche.
* 31 mL de jus de citron.
* 50 g de farine.
* 64 g de sucre.
* 1 oeufs.
* 69 mL de lait.
* 31 cuilleres de vanille.
* 21 g de cacao.
* 70 mL d'huile vegetale.
* 31 mg de sel.
* 64 mg de bananes ecrasees.
* 1 mL de jus de pomme.
* 31 mg d'arachides (sauf en cas d'allergie).
* 69 pepites au chocolat.
* 64 mL de creme fraiche.
* 1 mg de romarin seche.
* 103 mL de jus de citron.
Fini les 42 biscuits!"""

out_number_list = [26, 40, 11, 25, 93, 40, 1, 35, 50, 25, 70, 31, 16, 35, 31, 50, 64, 1, 69, 31, 21, 70, 31, 64, 1, 31, 69, 64, 1, 103]

# Formula must be reversed from binary which followed formula ((((val - 90) * 3 )^ 0b110) - 8) seen in cpp file.
def goback(inp):
    return (((inp + 8) ^ 0b110) // 3) + 90

decoded = []
for c in out_number_list:
    res = goback(c)
    decoded.append(chr(res))

print("".join(decoded))
