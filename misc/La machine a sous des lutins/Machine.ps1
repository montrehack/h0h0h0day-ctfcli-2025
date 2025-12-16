## [1.0.0] - 2016-08-01 - Initial script
## [1.0.1] - 2017-01-20 - Changed the format of the write-host messages
## [1.0.2] - 2020-03-31 - Add a message about social distancing
## [1.0.3] - 2023-09-17 - Removed the message about social distancing
## [1.1.0] - 2024-11-05 - Integrate Rudolph-AI chatbot
## [1.1.1] - 2024-11-10 - Removed AI chatbot
## [HO.Ho.oh] - NOEL - Magie de NOEL
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#    ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
# HO Ho ho 

#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#    ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
Set-Alias -Name Preparer-la-luge -Value Set-Alias -Scope Script
Preparer-la-luge -Name Faire-une-maison-en-pain-d-epice -Value Get-FileHash -Scope Script
Preparer-la-luge -Name Placer-les-cadeaux-dans-le-sac -Value Get-Random -Scope Script
Preparer-la-luge -Name Emballer-le-cadeau -Value select-object -Scope Script
Preparer-la-luge -Name Manger-les-biscuits -Value write-host -Scope Script
Preparer-la-luge -Name Faire-une-farce -Value Where-Object -Scope Script
Preparer-la-luge -Name Remplir-le-bas-de-noel -Value Start-Sleep -Scope Script
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#    ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。
# HO Ho ho 

function Accrocher-Lumieres {
    Param ([PSCustomObject[]]$Ampoules)
$Lumieres = "* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *"
$NezRouge = 0
#⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
foreach ($Lumiere in $Lumieres.ToCharArray()) {
    $color = $Ampoules[$NezRouge % $Ampoules.Count]
    Manger-les-biscuits $Lumiere -ForegroundColor $color -NoNewline
    $NezRouge++
}
Manger-les-biscuits ""
#    ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。    
}



$Jouets = @("FLAG-482","FLAG-917","FLAG-356","FLAG-204","FLAG-839","FLAG-123","FLAG-128","FLAG-593","FLAG-745","FLAG-310","FLAG-926","FLAG-457","FLAG-682","FLAG-374","FLAG-815","FLAG-233","FLAG-509","FLAG-798","FLAG-642","FLAG-190")
$Emballages = @("aZ9kL3pQwX7dT1mR8sBvC","G4hYt2nKjQ8rW5xLzM0eD","pR7vC1bXyT9mL4qZsF8gH","W3xLzM0eDgH9kJ2rT7vC1","tN6jQ8rW5xLzM0eDgH9kJ","9de641f2be074675e9156","mL4qZsF8gH9kJ2rT7vC1b","X7dT1mR8sBvC4hYt2nKjQ","rT7vC1bXyT9mL4qZsF8gH","kJ2rT7vC1bXyT9mL4qZsF","LzM0eDgH9kJ2rT7vC1bXy","F8gH9kJ2rT7vC1bXyT9mL","Q8rW5xLzM0eDgH9kJ2rT7","dT1mR8sBvC4hYt2nKjQ8r","sBvC4hYt2nKjQ8rW5xLzM","H9kJ2rT7vC1bXyT9mL4qZ","gH9kJ2rT7vC1bXyT9mL4q","nKjQ8rW5xLzM0eDgH9kJ2","xLzM0eDgH9kJ2rT7vC1bX","vC4hYt2nKjQ8rW5xLzM0e","T9mL4qZsF8gH9kJ2rT7vC")
$Noeuds = @("apple","rivermusic","light","dream","hello","grass","ocean","field","mountain","bridge","flower","shadow","wind","path","tree")

$Noeuds = $Noeuds | Faire-une-farce { $_ -ne "hello" }
# HO Ho ho ;)
write-host $Noeuds

$PoleNord = "E9120818B780A5846795D1572AF45D806AE79B401B40555083BD03408E354DBF"

$Biscuit = $Jouets | Placer-les-cadeaux-dans-le-sac
Remplir-le-bas-de-noel -Seconds 3
Accrocher-Lumieres -Ampoules @("Red","Red","White")
Manger-les-biscuits "Premier symbole:  $($Biscuit)" -ForegroundColor "Green"
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
$Lait = $Emballages | Placer-les-cadeaux-dans-le-sac
Remplir-le-bas-de-noel -Seconds 3
Accrocher-Lumieres -Ampoules @("Red","Red","White")
Manger-les-biscuits "Deuxieme symbole:  $($Lait)" -ForegroundColor "Green"
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
$Cheminee = $Noeuds | Placer-les-cadeaux-dans-le-sac
Remplir-le-bas-de-noel -Seconds 3
Accrocher-Lumieres -Ampoules @("Red","Red","White")
Manger-les-biscuits "Troisieme symbole:  $($Cheminee)" -ForegroundColor "Green"
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
$Sapin= $Biscuit+$Lait+$Cheminee
Remplir-le-bas-de-noel -Seconds 3
Accrocher-Lumieres -Ampoules @("Red","Red","White")
Manger-les-biscuits "Resultat:  $($Sapin)" -ForegroundColor "Yellow"
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
$Etoile = [System.Text.Encoding]::UTF8.GetBytes($Sapin)
$stringAsStream = [System.IO.MemoryStream]::new($Etoile)

Remplir-le-bas-de-noel -Seconds 3
Manger-les-biscuits "Roulement de tambour" -ForegroundColor "Yellow"
Remplir-le-bas-de-noel -Seconds 3

Accrocher-Lumieres -Ampoules @("Red","Green","White")
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
$hash = Faire-une-maison-en-pain-d-epice -Algorithm "SHA256" -InputStream $stringAsStream | Emballer-le-cadeau Hash
#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
if ($hash.hash.ToUpper() -eq $PoleNord.ToUpper()) {
    #Write-host "Bravo! Vous avez gagne, Prenez votre billet d'or et allez chercher votre recompense" -ForegroundColor "Green"
    #Print-award -color "Gold"
    # HO Ho ho ;) 
}else {
    #Write-host "Desole, meilleur chance la prochaine fois!" -ForegroundColor "Red"
    # HO Ho ho ;)
    Manger-les-biscuits "Perdu! Retournez au travail faineants!! HO Ho ho" -ForegroundColor "Red"
}
Accrocher-Lumieres -Ampoules @("Red","Green","White")


#  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。 
#    ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。  ⋆⁺₊❅。
# HO Ho ho