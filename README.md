- Use mailtrap to receive email; link : [mailtrap inboxs](https://mailtrap.io/inboxes)
- Change SMTP_HOST; SMTP_USER; SMTP_PASS in .env
- Starting the application : 
```shell 
docker-compose up 
```
- Example :
    - http://localhost:3000/signup 
        {
            "email":"username@gmail.com",
            "username":"username",
            "password":"aaaa"
        }
    - http://localhost:3000/login 
        {
             "username":"username",
             "password":"aaaa"
        }
    - http://localhost:3000/login/verify
        {
             "email":"username@gmail.com",
             "otpcode":""
        }
    - http://localhost:3000/refresh
        - JWT Bearer
    -http://localhost:3000/me
        - JWT Bearer