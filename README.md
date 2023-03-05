# 🦾 VallHalla-api 🥃

<i><b>Valhalla-api</b></i> provides fast access to the database of drink recipes from the game <b>"VA-11 Hall-A: Cyberpunk Bartender Action"</b>

### Overview

- [Overview](#overview)
- [Request](#request)
- [Structures](#structures)
- [Types](#types)
- [Shortcuts](#shortcuts)


### Request
  
``` Elixir
https://vall-halla-api.vercel.app/api/search
``` 

Parameter       | Value type | Description   
----------------|------------|-------------------------------------
name            |   string   | search by names ('%VALUE%')
price           |    int     | search by prices (==)
alcoholic       |   string   |
ice             |   string   |
flavour         |   string   | search by flavours (==)
type            |   string   | search by types ('%VALUE%')
shortcut        |   string   | search by shortcuts ('%VALUE%')
recipe          |   string   | search by recipes ('%VALUE%')
description     |   string   | search by descriptions ('%VALUE%')

<p>For example, </br>"https://vall-halla-api.vercel.app/api/search?alcoholic=no&flavour=spicy" - all non-alcoholic spicy drinks,</br>
"https://vall-halla-api.vercel.app/api/search?name=piano" - "Piano Man" and "Piano Woman" recieps,</br>
"https://vall-halla-api.vercel.app/api/search?recipe=aged" - all aged drinks,</br>
"https://vall-halla-api.vercel.app/api/search?type=promo&shortcut=3xT" - all promo drinks with 3 Karmotrine</p>   

### Structures 

#### Response

Field                       |       Type         | Description
----------------------------|--------------------|------------
success                     |       bool         | response status
error                       |      string        | 
result                      |     []Drink        | slice of recipes


#### Drink

Field                       |       Type         | Description
----------------------------|--------------------|------------
name                        |      string        |
price                       |       int          |
alcoholic                   |      string        | "Yes", "No" or "Optional"
ice                         |      string        | "Yes" or "No"
flavour                     |      string        | may be "N/A"
primary_type                |      string        |
secondary_type              |      string        |
shortcut                    |      string        | also may be "N/A"


### Types

Type                     |                                    Values       
-------------------------|------------------------------------------------------------------------
primary_type             | "Bottled", "Classic", "Classy", "Girly", "Manly" or "Promo"
secondary_Type           | "Bland", "Burning", "Happy", "N/A", "Sobering", "Soft", "Strong" or "Vintage"


### Shortcuts

Shortcuts   |    Action   
------------|-------------------- 
Q           | Add one Adelhyde
W           | Add one Bronson Extract
E           | Add one Powdered Delta
R           | Add one Flanergide
T           | Add one Karmotrine (alcoholic component)
A           | Toggle ice
S           | Toggle aging

<p> For example, <i>"2xQ, 3xW, 5xE, 5xR, 3xT, A, all mixed."</i> means <i>"2 Adelhyde, 3 Bronson Extract, 5 Powdered Delta, 5 Flanergide and 3 Karmotrine, all on the rocks and mixed."</i></p>

You can also use route <b>"/info"</b> instead of <b>"/search"</b> to get more information about drink (with description and full recipe). Parameters are the same.


<img src="https://wakatime.com/badge/user/ee2709af-fc5f-498b-aaa1-3ea47bf12a00/project/ca6a9f63-8582-4243-905e-900ec35cede8.svg?style=for-the-badge">

[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
