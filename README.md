## Food Search

Finds a meal that meets your macro goals via tournament selection and mutation.

Defaults are set in `target` in the main function

All values are per 100g. Data is taken from from the AUSNUT 2011-13 food nutrient database:
https://www.foodstandards.gov.au/science/monitoringnutrients/ausnut/ausnutdatafiles/Pages/foodnutrient.aspx


Columns of `foods.json` are the following:

```json
["Survey ID",
  "Food Name",
  "Survey flag",
  "Energy, with dietary fibre (kJ)", 
  "Energy, without dietary fibre (kJ)", 
  "Moisture (g)", 
  "Protein (g)", 
  "Total fat (g)",
 "Available carbohydrates, with sugar alcohols (g)",
 "Available carbohydrates, without sugar alcohol (g)",
 "Starch (g)",
 "Total sugars (g)",
 "Added sugars (g)",
 "Free sugars (g)",
 "Dietary fibre (g)",
 "Alcohol (g)",
 "Ash (g)",
 "Preformed vitamin A (retinol) (g)",
 "Beta-carotene (g)",
 "Provitamin A (b-carotene equivalents) (g)",
 "Vitamin A retinol equivalents (g)",
 "Thiamin (B1) (mg)",
 "Riboflavin (B2) (mg)",
 "Niacin (B3) (mg)",
 "Niacin derived equivalents (mg)",
 "Folate, natural  (g)",
 "Folic acid  (g)",
 "Total Folates  (g)",
 "Dietary folate equivalents  (g)",
 "Vitamin B6 (mg)",
 "Vitamin B12  (g)",
 "Vitamin C (mg)",
 "Alpha-tocopherol (mg)",
 "Vitamin E (mg)",
 "Calcium (Ca) (mg)",
 "Iodine (I) (g)",
 "Iron (Fe) (mg)",
 "Magnesium (Mg) (mg)",
 "Phosphorus (P) (mg)",
 "Potassium (K) (mg)",
 "Selenium (Se) (g)",
 "Sodium (Na) (mg)",
 "Zinc (Zn) (mg)",
 "Caffeine (mg)",
 "Cholesterol (mg)",
 "Tryptophan (mg)",
 "Total saturated fat (g)",
 "Total monounsaturated fat (g)",
 "Total polyunsaturated fat (g)",
 "Linoleic acid (g)",
 "Alpha-linolenic acid (g)",
 "C20:5w3 Eicosapentaenoic (mg)",
 "C22:5w3 Docosapentaenoic (mg)",
 "C22:6w3 Docosahexaenoic (mg)",
 "Total long chain omega 3 fatty acids (mg)",
 "Total trans fatty acids (mg)",
 "category"
]
```

