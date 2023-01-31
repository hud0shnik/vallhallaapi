# 🦾 VallHalla-api 🥃

<i>Valhalla-api</i> provides fast access to the database of drink recipes from the game <b>"VA-11 Hall-A: Cyberpunk Bartender Action"</b>

<h2>Structures</h2>

<h4>Response</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
success                     |       bool         | response status
error                       |      string        | 
result                      |     []Drink        | slice of recipes


<h4>Drink</h4>

Field                       |       Type         | Description
----------------------------|--------------------|------------
name                        |      string        |
price                       |       int          |
alcoholic                   |      string        | "Yes", "No" or "Optional"
ice                         |      string        | "Yes" or "No"
flavour                     |      string        |
primary_type                |      string        |
secondary_type              |      string        |  may be ""
recipe                      |      string        |  also may be ""


<h2>Types</h2>

Type                     |                                    Values       
-------------------------|------------------------------------------------------------------------
primary_type             |"Bottled", "Classic", "Classy", "Girly", "Manly" or "Promo"
secondary_Type           | "", "Bland", "Burning", "Happy", "Sobering", "Soft", "Strong" or "Vintage"


<h2>Recipe</h2>

Shortcuts   |    Add one   
------------|-------------------- 
Q           | Adelhyde
W           | Bronson Extract
E           | Powdered Delta
R           | Flanergide
T           | Karmotrine (alcoholic component)

<p>For example, "2xQ, 3xW, 5xE, 5xR, 3xT" means "2 Adelhyde, 3 Bronson Extract, 5 Powdered Delta, 5 Flanergide and 3 Karmotrine" </p> </br>


