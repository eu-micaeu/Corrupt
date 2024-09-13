import { getCookie } from './getCookie.mjs';

function loadGoals() {

    fetch('/listGoals', {

        method: 'GET',

        headers: {

            'Content-Type': 'application/json',

            'Authorization': getCookie('token')

        }
    })
        .then(response => response.json())
        
        .then(data => {

            const goals = data.goals;
            
            const goalList = document.getElementById('goalList');

            goalList.innerHTML = ''; 

            goals.forEach(goal => {

                const goalTitle = document.createElement('h3');

                goalTitle.textContent = goal.title;

                const goalEnd = document.createElement('p');

                const dueDate = new Date(goal.due_date);

                function updateCountdown() {

                    const now = new Date();

                    let timeDifference = dueDate - now;

                    if (timeDifference < 0) {

                        timeDifference = 0;

                        goalEnd.textContent = "Time's up!";

                        return;

                    }

                    const days = Math.floor(timeDifference / (1000 * 60 * 60 * 24));

                    const hours = Math.floor((timeDifference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));

                    const minutes = Math.floor((timeDifference % (1000 * 60 * 60)) / (1000 * 60));

                    const seconds = Math.floor((timeDifference % (1000 * 60)) / 1000);

                    goalEnd.textContent = `Remaining time: ${days}:${hours}:${minutes}:${seconds}`;

                }

                updateCountdown();

                setInterval(updateCountdown, 1000);

                const goalDueDate = document.createElement('p');

                goalDueDate.textContent = `Due date: ${dueDate.toDateString()}`;

                const now = new Date();

                const timeDifference = dueDate - now;

                const days = document.createElement('p');

                days.textContent = Math.floor(timeDifference / (1000 * 60 * 60 * 24)) + ' days left';
                
                const goalDiv = document.createElement('div');        
                
                goalDiv.addEventListener('click', function () {

                    window.location.href = '/viewGoal/' + goal.goal_id;

                });

                goalDiv.appendChild(goalTitle);

                goalDiv.appendChild(goalEnd);

                goalDiv.appendChild(goalDueDate);

                goalDiv.appendChild(days);

                goalDiv.className = 'goalDiv';

                const btApagar = document.createElement('button');

                btApagar.textContent = 'Delete';

                btApagar.className = 'btApagar';

                btApagar.addEventListener('click', function (event) {

                    event.stopPropagation();
                
                    fetch('/deleteGoal/' + goal.goal_id, {

                        method: 'DELETE',

                        headers: {

                            'Content-Type': 'application/json',

                            'Authorization': token

                        },

                    })

                    .then(response => {

                        if (response.status === 200) {

                            loadGoals();

                        }

                    })

                    .catch((error) => {

                        console.error('Error:', error);

                    });

                });
                
                goalDiv.appendChild(btApagar);

                goalList.appendChild(goalDiv);

            });

        })

        .catch((error) => {

            console.error('Error:', error);

        });

}

document.addEventListener('DOMContentLoaded', function () {

    loadGoals(); 

});

document.getElementById('btSaveGoal').addEventListener('click', function () {

    const token = getCookie('token');

    const title = document.getElementById('title').value;

    const description = document.getElementById('description').value;

    const dueDate = document.getElementById('due_date').value;

    fetch('/createGoal', {

        method: 'POST',

        headers: {

            'Content-Type': 'application/json',

            'Authorization': token

        },

        body: JSON.stringify({

            title: title,

            description: description,

            due_date: dueDate + ':00Z'

        })

    })

        .then(response => {

            if (response.status === 200) {

                loadGoals(); 

                document.getElementById('overlay').style.display = 'none';

                document.getElementById('popUpAddGoal').style.display = 'none';

                document.getElementById('title').value = '';

                document.getElementById('description').value = '';

                document.getElementById('due_date').value = '';

            }

            return response.json();
            
        })

        .catch((error) => {

            console.error('Error:', error);

        });

});

document.getElementById('btAddGoal').addEventListener('click', function () {

    const overlay = document.getElementById('overlay');

    overlay.style.display = 'block';

    const popUpAddGoal = document.getElementById('popUpAddGoal');

    popUpAddGoal.style.display = 'flex';

});

document.getElementById('overlay').addEventListener('click', function () {

    this.style.display = 'none';

    const popUpAddGoal = document.getElementById('popUpAddGoal');

    popUpAddGoal.style.display = 'none';

});