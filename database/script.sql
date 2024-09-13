CREATE TABLE goals (

    goal_id SERIAL PRIMARY KEY,           
    title VARCHAR(255) NOT NULL,         
    description TEXT,                  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    due_date TIMESTAMP,                  
    user_id INT NOT NULL,              
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
    
);
