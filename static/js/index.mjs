import { toastGreen, toastRed } from "./toast.mjs";

document.getElementById("btEnter").addEventListener("click", function() {

    var username = document.getElementById("username").value;

    var password = document.getElementById("password").value;

    var data = {

        username: username,

        password: password

    };

    fetch("/login", {

        method: "POST",

        body: JSON.stringify(data),

    }).then(function(response) {

        if(response.status == 200) {

            toastGreen("Login successful, redirecting...");

            setInterval(() => {

                window.location.href = "/home";

            }, 2000);

        } else {

            toastRed("Invalid username or password");

        }

        return response.json();

    })

    }
)