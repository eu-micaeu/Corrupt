import { getCookie } from "./getCookie.mjs";

document.getElementById("btUpdateGoal").addEventListener("click", function() {

    let goalTitle = document.getElementById("goalTitle").value;

    let goalDescription = document.getElementById("goalDescription").value;

    let goalID = this.getAttribute("data-goal-id");

    let goal = {

        Title: goalTitle,

        Description: goalDescription,

    };

    fetch("/updateGoal/" + goalID, {

        method: "PUT",

        headers: {

            "Content-Type": "application/json",

            "Authorization": getCookie("token")

        },

        body: JSON.stringify(goal)

    })

        .then(response => {

            if (response.status === 200) {

                window.location.href = "/home";

            }

        })

        .catch(error => {

            console.error("Error:", error);

        });

});