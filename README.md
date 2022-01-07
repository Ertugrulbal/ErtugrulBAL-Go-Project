# ErtugrulBAL-Go-Project
-This Project about the Link Converting Processing

-This Project aims to convert to links each other which Web Urls to Deeplinks and Deeplinks to Web Urls 

-We have two types Link.


Examples of Link types and equality theirself:
1. First style of Link Converting Type ( "Produc"t )
    - Web URL : "https://www.ecommerce.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064"
    - Deep Link : "ec://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"
2. Second style of Link Converting Type ("Search")
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

    ![image](https://user-images.githubusercontent.com/92356291/148546326-66d3a167-cbfa-4865-a52e-54728531dce5.png)


    
