// Inserting comments and replies in the database
function insertCommentInDatabase(username, commentText) {
    fetch('/postComment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username: username, commentText: commentText })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Comment inserted successfully:', data);
    })
    .catch(error => {
        console.error('Error inserting comment:', error);
    });
}

// Function to insert a reply in the database
function insertReplyInDatabase(commentID, username, replyText) {
    fetch('/postReply', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ commentID: commentID, username: username, replyText: replyText })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Reply inserted successfully:', data);
    })
    .catch(error => {
        console.error('Error inserting reply:', error);
    });
}

