# vdock-hugo

Note: 
There are still couple of things to be fixed in this repo. 

To run: 
```
# Run container
$ docker run -td -p 8288:8288-p 8286:8286 rchari/vdock-hugo

# Once you start the container; run below cmd to post "github information" to the service
# edit data with your github details.
$ curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "userId": "rchari66",                
    "blogRepo": "https://github.com/rchari66/tt-blog.git",
    "siteRepo": "https://github.com/rchari66/rchari66.github.io.git",
    "email" : "raghav0966@gmail.com",
    "password":"***ValidPassword**"
   }' 
   http://localhost:8288/api/login

# I've used postman to send post request(content-type:raw) to "http://localhost:8288/api/v1/conf/updateConfig".


# After the above curl cmd wait for it to checkout repositories from github.


# you have to manually run "hugo server"(it will be fixed soon :))
$ api_hops_start_hugo 

```


