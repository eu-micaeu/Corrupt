import { toastGreen } from "./toast.mjs";

document.getElementById("btRegister").addEventListener("click", function() {

    var full_name = document.getElementById("full_name").value;

    var email = document.getElementById("email").value;

    var username = document.getElementById("username").value;

    var password = document.getElementById("password").value;

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

            alert("Error: " + response.status);

        }

    })

    .catch(function(error) {

        console.log(error);
        
    });

});