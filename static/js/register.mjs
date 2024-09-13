import { toastGreen, toastRed } from "./toast.mjs";

document.getElementById("btRegister").addEventListener("click", function() {

    var full_name = document.getElementById("full_name").value;

    var email = document.getElementById("email").value;

    var username = document.getElementById("username").value;

    var password = document.getElementById("password").value;

    // Verificação email

    if (email.indexOf("@") == -1 || email.indexOf(".") == -1) {

        toastRed("Invalid email!");

        return;

    }

    // Verificação senha

    if (password.length < 8) {

        toastRed("Password must have at least 8 characters!");

        return;

    }

    var data = {

        full_name: full_name,

        email: email,

        username: username,

        password: password,

    };

    fetch("/register", {

        method: "POST",

        headers: {

            "Content-Type": "application/json"

        },

        body: JSON.stringify(data)

    })


    .then(function(response) {

        if (response.status == 200) {

            toastGreen("User registered successfully! Redirecting to login page...");

            setInterval(function() {

                window.location.href = "/";

            }, 3000);

        } else {

            toastRed("Error registering user!");

        }

    })

    .catch(function(error) {

        console.log(error);
        
    });

});