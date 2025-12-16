#include <iostream>
#include <cstdlib>
#include <string.h>

#include "miss_claus_secret_recipe.h"
using namespace std;

string ingredients[] = {
    "g de farine",
    "g de sucre",
    "oeufs",
    "mL de lait",
    "cuilleres de vanille",
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
        cout << "C'est le temps de cuisiner un petit dessert." << endl;
        unsigned long cookies = bakeFakeCookies(argv[1]);
        cout << "Fini les " << cookies << " biscuits!" << endl;
        cout << "Hmmm, etait-ce la bonne recette?" << endl;
        cookies = bakeCookies(argv[1]);        
        cout << "Fini les " << cookies << " biscuits!" << endl;
        cout << "Ah, ca sent bon maintenant." << endl;
    }

    else {
        cout << "Ho ho! Cette fois-ci il y en a trop." << endl;
        sayHiToRudolph();
    }

    return 0;
}

unsigned long bakeFakeCookies(char* secretQuantities) {
    cout << "Commencons..." << endl;
    char* res;

    unsigned long val = 0;
    for (unsigned int i = 0; i < strlen(secretQuantities); i++)
    {
        val = 20+i;
        val = (((val )^ 0b1010));
        cout << "* " << val << " " << ingredients[i % ingredients_len] << "." << endl;
    }
    
    return 24;
}

unsigned long bakeCookies(char* secretQuantities) {
    char* res;

    unsigned long val = 0;
    for (unsigned int i = 0; i < strlen(secretQuantities); i++)
    {
        val = (unsigned long) secretQuantities[i];
        val = ((((val - 90) * 3 )^ 0b110) - 8);
        cout << "* " << val << " " << ingredients[i % ingredients_len] << "." << endl;
    }
    
    return 42;
}

void sayHiToRudolph() {
    cout << "Coucou Rudolphe!" << endl << "- Mere Noel." << endl;
}
