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

        return response.json();

    }).then(function(data) {

        if (data.token) {

            localStorage.setItem("token", data.token);

            window.location.href = "/home";

        } else {
                
            alert("Login failed");

        }

    });

    }
)