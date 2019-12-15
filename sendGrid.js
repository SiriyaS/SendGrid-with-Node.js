// Import Library
const Mail = require("@sendgrid/mail");

Mail.setApiKey(process.env.SENDGRID_API_KEY);

// Message part
const message = {
    to : "siriya013@gmail.com",
    from : "pingsiriya@gmail.com",
    subject: "Sending mail with SendGrid",
     text: "This email is created with Node.js",
     html: "<strong>This email is created with Node.js</strong>",
};

// Sending mail
Mail.send(message);