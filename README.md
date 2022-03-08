# ErtugrulBAL-Go-Project
-This Project about the Link Converting Processing

-This Project aims to convert to links each other which Web Urls to Deeplinks and Deeplinks to Web Urls 

-We have two types of links.


Examples of Link types and equality theirself:
1. First style of Link Converting Type ( "Product ) If Url has "-p-" text we decided a this link "Product" type.
    - Web URL : "https://www.ecommerce.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064"
    - Deep Link : "ec://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"
2. Second style of Link Converting Type ("Search"). If Url has "sr?q=" text we decided a this link "Search" type.
    - Web URL : "https://www.ecommerce.com/sr?q=%C3%BCt%C3%BC"
    - Deep Link : "ec://?Page=Search&Query=%C3%BCt%C3%BC"
3. Third style of Link Converting Type ( "Home") (İf request not equal to "Product" or "Search" we call a "Home")
    - Web URL : "https://www.ecommerce.com/Hesabim/#/Siparisleri
    - Deep Link : "ec://?Page=Home"

##Setup Project Structure

1. First  DB Connection Established.

    ![image](https://user-images.githubusercontent.com/92356291/148545485-7b2ee9d9-f3e0-4ac3-b6fa-4a76575ca86f.png)
2. Secondly Web Server Initialized.

    ![image](https://user-images.githubusercontent.com/92356291/148546223-6383749b-6e31-473f-a861-81d2aeed9f1c.png)
## Then We passed to establishing to Structure.
![image](https://user-images.githubusercontent.com/92356291/148573437-6388dc3d-6f31-47d2-af56-ef9779e4f744.png)
    


    - Controller Package has a ConvertingController.go file. This file provide us When the request come from server, directing to converting from deeplink to Web URL or convert from WebUrl to deeplink.
    
    - Converters meet the requests from the Controller and perform the converting operations.
    
    - Driver file established for DB connection as we mentioned before.
    
    - Object package have a WebUrl and Deeplink files. These files perform parceling and assigning Web Url and Deeplink requests to models.
    
    - Repository was created to store inbound and fetched links.
    
    - Utils file was created for management of errors.
    
    -.env file has stored secret DB informations. 
### V2- This code were written Gin Framework

- All of the above principle valid for this style. 
- Using Frameworks more easy to arrive result.

# We imported a these packages

![image](https://user-images.githubusercontent.com/92356291/148694552-8ec47657-740f-4f51-a3ed-1a3e4703c081.png)

-In this version we showed a simple coding style and arrive a fastly to result.
-The code has a approximately 200 raw 
### TEST RESULTS
- We tested Web URLs and Deeplinks with POSTMAN

1. Web Url convert to Deeplink

    -( Product Style)
    ![image](https://user-images.githubusercontent.com/92356291/148695697-7a73919b-39bf-434c-beaf-1641f3fbdef2.png)
    -(Search Style)
    ![image](https://user-images.githubusercontent.com/92356291/148695743-ff0c5cef-58ad-4158-a190-79fb9428c55a.png)
    -( Home Style )
    ![image](https://user-images.githubusercontent.com/92356291/148695768-eb5a71db-7508-46cc-afd9-e249fcf17848.png)

2. Deeplink convert to Web URL

    -( Product Style )
    ![image](https://user-images.githubusercontent.com/92356291/148695795-c8e69999-ee4e-4a90-9179-c9b99a614e96.png)
    -(Search Style)
    ![image](https://user-images.githubusercontent.com/92356291/148695816-b8f0b8c5-10bb-43e0-8964-aa16cfdb6e65.png)
    -(Other Style )
    ![image](https://user-images.githubusercontent.com/92356291/148695853-f642d496-1956-4552-8b60-29b32f29b4bf.png)

    
