# SendGrid-with-Node.js

### Sending Email with SendGrid and Node.js
##### ก่อนอื่นเลย เราต้องไป sign up ที่เว็บ app.sendgrid.com เพื่อที่เราจะได้เข้าไปสร้าง API Key 
### Create API Key 
##### log in to app.sendgrid.com > go to Settings > API Keys > Create API Key 
##### ตั้งชื่อ API Key กดเลือก Restricted Access และเลื่อนลงไปตรงเมนู Mail Send กดลูกศรลง แล้วกดตรงวงกลมขวาสุดของ Mail Send  เสร็จแล้วกด Create & View  
##### หลังจากนั้นจะมี API Key ขึ้นมา ให้เราก็อป key นั้นไว้ !!อย่าลืมก็อป key ไว้!!
### Create Folder 
##### open terminal >  
```
  mkdir FolderName  
  cd FolderName 
```
  ### Install npm 
```
 npm init --yes 
```
 ### Install library
 ```
  npm install @sendgrid/mail
 ```
 ### วิธีเก็บ API Key ใน environment
```
 echo "export SENDGRID_API_KEY='API Key ที่เราเพิ่งก็อปมา'" > sendgrid.env
 echo "sendgrid.env" >> .gitignore
 source ./sendgrid.env
```
 ##### สร้างไฟล์ในโฟลเดอร์ที่เราเพิ่งสร้างขึ้น ในที่นี้เราจะให้ไฟล์ชื่อว่า sendGrid.js
 ```
 const Mail = require("@sendgrid/mail");
 Mail.setApiKey(process.env.SENDGRID_API_KEY);
 const message = {
     to: "receiver@example.com",
     from: "sender@example.com",
     subject: "Sending mail with SendGrid",
     text: "This email is created with Node.js",
     html: "<strong>This email is created with Node.js</strong>",
 };
 Mail.send(message);
 ```
 ##### ซึ่งเราสามารถใส่ style หรือ แก้ไขเนื้อหาภายใน mail โดยใช้ HTML tag 
 
 ##### เมื่อเสร็จแล้วก็ save file และกลับไปที่ terminal เพื่อ run โปรแกรม
 ```
 node sendGrid.js
 ```
 ##### ตอนนี้เมลจะถูกส่งไปแล้วให้เราลองไปเช็คใน inbox ของผู้รับ หรือไปเช็คใน account ของเราที่ app.sendgrid.com ตรง Activity 
 
 >ศึกษาเพิ่มเติมได้ในวิดีโอนี้ https://www.youtube.com/watch?v=s2bzUzHeSVw
 >หรือใน Setup Guide > Integrate using our Web API or SMTP relay > Web API > Node.js
