#include <iostream>
#include <cstdlib>
#include <string.h>

#include "miss_claus_secret_recipe_2.h"
using namespace std;

string ingredients[] = {
    "g de farine",
    "g de sucre",
    "oeufs",
    "mL de lait",
    "cuilleres de vanille",
    "mL d'extrait d'arc-en-ciel",
    "g de cacao",
    "mL d'huile vegetale",
    "mg de sel",
    "mg de bananes ecrasees",
    "mL de jus de pomme",
    "mg d'arachides (sauf en cas d'allergie)",
    "pepites au chocolat",
    "mL de creme fraiche",
    "mg de romarin seche",
    "mL de jus de citron",
    "flocons de neige",
    "mg de gentillesse",
    "mL de sirop d'erable",
};
size_t ingredients_len = sizeof(ingredients)/sizeof(ingredients[0]);

int main(int argc, char *argv[]){
    if(argc < 1) {
        cout << "Oh mon Noël, comment sommes-nous arrivés ici?" << endl;
        sayHiToRudolph();
    }
        
    else if(argc < 2) {
        cout << "Ho ho ho! Il manque quelque chose, non?" << endl;
        sayHiToRudolph();
    }


    else if(argc == 2) {
        char* res;
        cout << "C'est le temps de cuisiner un petit dessert. J'espere que personne ne regarde." << endl;
        unsigned long cookies = bakeFakeCookies(argv[1], 0);
        cout << "Fini les " << cookies << " \"biscuits\"! ;)" << endl;
        cout << "Hah! personne ne saura que cette recette n'est pas ce dont elle a l'air." << endl;
        cout << "Reprenons." << endl;
        cookies = bakeCookies(argv[1], cookies);
        cout << "Fini les " << cookies << " biscuits!" << endl;
        cout << "Ah, ca sent specialement bon maintenant!!" << endl;
    }

    else {
        cout << "Ho ho! Cette fois-ci il y en a trop." << endl;
        sayHiToRudolph();
    }

    return 0;
}

unsigned long bakeFakeCookies(char* secretQuantities, unsigned long iv) {
    cout << "Commencons..." << endl;
    char* res;

    unsigned long val = 0;
    for (unsigned int i = 0; i < strlen(secretQuantities); i++)
    {
        val = 20+i;
        val = (((val )^ 0b1010));
        cout << "* " << val << " " << ingredients[i % ingredients_len] << "." << endl;
    }
    
    return val;
}

unsigned long bakeCookies(char* secretQuantities, unsigned long iv) {
    char* res;

    unsigned long val = iv;
    for (unsigned int i = 0; i < strlen(secretQuantities); i++)
    {
        val = (unsigned long) secretQuantities[i] ^ val;
        val = ((((val + 2) ) ^ 0b110) + 1);
        cout << "* " << val << " " << ingredients[i % ingredients_len] << "." << endl;
    }
    
    return 42;
}

void sayHiToRudolph() {
    cout << "Coucou Rudolphe!" << endl << "- Mere Noel." << endl;
}
