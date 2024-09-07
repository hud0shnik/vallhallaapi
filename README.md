# 🦾 VallHalla-api-vercel-branch 🥃

Branch for deploying on Vercel

<i><b>Valhalla-api</b></i> provides fast access to the database of drink recipes from the game <b>"VA-11 Hall-A: Cyberpunk Bartender Action"</b>

## Overview

- [Overview](#overview)
- [Request](#request)
- [Structures](#structures)
- [Types](#types)
- [Shortcuts](#shortcuts)
- [Samples](#samples)

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

<p>For example,</br>"https://vall-halla-api.vercel.app/api/search?alcoholic=no&flavour=spicy" - all non-alcoholic spicy drinks,</br>
"https://vall-halla-api.vercel.app/api/search?name=piano" - "Piano Man" and "Piano Woman" recipes,</br>
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


### Samples

#### Request
``` Elixir
https://vall-halla-api.vercel.app/api/search?type=manly&alcoholic=optional
``` 

#### Response

``` Json
{
  "success": true,
  "error": "",
  "result": [
    {
      "name": "Crevice Spike",
      "price": 140,
      "alcoholic": "Optional",
      "ice": "No",
      "flavour": "Sour",
      "primary_type": "Manly",
      "secondary_type": "Sobering",
      "shortcut": "2xE, 4xR, optional T, all blended."
    },
    {
      "name": "Gut Punch",
      "price": 80,
      "alcoholic": "Optional",
      "ice": "No",
      "flavour": "Bitter",
      "primary_type": "Manly",
      "secondary_type": "Strong",
      "shortcut": "5xW, 1xR, optional T, S, all mixed."
    }
  ]
}
```

#### Request

``` Elixir
https://vall-halla-api.vercel.app/api/info?flavour=spicy&recipe=bronson%20extract
``` 

#### Response

``` Json
{
  "success": true,
  "error": "",
  "result": [
    {
      "name": "Bleeding Jane",
      "price": 200,
      "alcoholic": "No",
      "ice": "No",
      "flavour": "Spicy",
      "primary_type": "Classic",
      "secondary_type": "Sobering",
      "recipe": "1 Bronson Extract, 3 Powdered Delta and 3 Flanergide. All blended.",
      "shortcut": "1xW, 3xE, 3xR, all blended.",
      "description": "Say the name of this drink three times in front of a mirror and you'll look like a fool."
    },
    {
      "name": "Marsblast",
      "price": 170,
      "alcoholic": "Yes",
      "ice": "No",
      "flavour": "Spicy",
      "primary_type": "Manly",
      "secondary_type": "Strong",
      "recipe": "6 Bronson Extract, 1 Powdered Delta, 4 Flanergide and 2 Karmotrine. All blended.",
      "shortcut": "6xW, 1xE, 4xR, 2xT, all blended.",
      "description": "One of these is enough to leave your face red like the actual planet."
    }
  ]
}
```

<img src="https://wakatime.com/badge/user/ee2709af-fc5f-498b-aaa1-3ea47bf12a00/project/ca6a9f63-8582-4243-905e-900ec35cede8.svg?style=for-the-badge">

[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)
