// // Show comment input
// function showCommentInput(event) {
//     event.preventDefault();

//     var username = getUsername();

//     var commentInputDiv = getOrCreateCommentInput();

//     // Inserting the comment input after the comments
//     var postPrivate = event.target.closest('.PostPrivate');
//     var commentsDiv = postPrivate.querySelector('.comments');
//     postPrivate.insertBefore(commentInputDiv, commentsDiv.nextSibling);
// }

// // Function to show reply input
// function showReplyInput(event, username) {
//     event.preventDefault();

//     var commentInputDiv = getOrCreateCommentInput(); 

//     // Update the input for the replys
//     var input = commentInputDiv.querySelector('input');
//     input.dataset.replyingTo = 'true';
//     input.placeholder = 'Votre réponse...';

//     // Update the button for the reply
//     var button = commentInputDiv.querySelector('button');
//     button.textContent = 'Envoyer la réponse';

//     // Insert the comment input after the comment
//     var commentElement = event.target.closest('.comment');
//     var repliesDiv = commentElement.querySelector('.replies');
//     if (!repliesDiv) {
//         repliesDiv = document.createElement('div');
//         repliesDiv.classList.add('replies');
//         commentElement.appendChild(repliesDiv);
//     }

//     repliesDiv.appendChild(commentInputDiv);
// }

// // Function to get the username
// function getUsername() {
//     return document.querySelector('.links a:last-child').textContent.trim();
// }

// // Function to get or create the comment input
// function getOrCreateCommentInput() {
//     var commentInputDiv = document.querySelector('.comment-input');
//     if (!commentInputDiv) {
//         commentInputDiv = createCommentInput();
//     }
//     return commentInputDiv;
// }

// // Function to create the comment input
// function createCommentInput() {
//     var commentInputDiv = document.createElement('div');
//     commentInputDiv.classList.add('comment-input');

//     var input = document.createElement('input');
//     input.type = 'text';
//     input.placeholder = 'Votre commentaire...';
//     input.style.marginLeft = '25%';
//     input.style.width = '50%';
//     input.style.border = '1px solid #ccc';
//     input.style.marginBottom = '20px';

//     var button = document.createElement('button');
//     button.textContent = 'Envoyer';
//     button.onclick = function() {
//         var input = commentInputDiv.querySelector('input');
//         var username = getUsername();
//         if (input.dataset.replyingTo) {
//             submitReply(input, username);
//         } else {
//             submitComment(input, username);
//         }
//     };

//     commentInputDiv.appendChild(input);
//     commentInputDiv.appendChild(button);

//     return commentInputDiv;
// }

// // Function to render the comments
// function createCommentElement(username, commentText) {
//     var commentElement = document.createElement('div');
//     commentElement.classList.add('comment');
//     commentElement.style.width = '40%';
//     commentElement.style.marginBottom = '20px';
//     commentElement.style.marginLeft = '30%';
//     commentElement.textContent = username + ': ' + commentText;

//     var replyLink = document.createElement('a');
//     replyLink.href = '#';
//     replyLink.textContent = 'Répondre';
//     replyLink.onclick = function(event) {
//         event.preventDefault();
//         showReplyInput(event, username);
//     };

//     commentElement.appendChild(document.createElement('br'));
//     commentElement.appendChild(replyLink);

//     var repliesDiv = document.createElement('div');
//     repliesDiv.classList.add('replies');
//     commentElement.appendChild(repliesDiv);

//     return commentElement;
// }

// // Function to create a reply element
// function createReplyElement(username, replyText) {
//     var replyElement = document.createElement('div');
//     replyElement.classList.add('reply');
//     replyElement.style.marginLeft = '35%';
//     replyElement.style.width = '35%';
//     replyElement.textContent = username + ': ' + replyText;

//     return replyElement;
// }

// // Function to render the comments
// function submitComment(input, username) {
//     var commentText = input.value.trim();

//     if (commentText !== '') {
//         var commentList = input.closest('.PostPrivate').querySelector('.comments');
//         var commentElement = createCommentElement(username, commentText);
//         commentList.appendChild(commentElement);
//         insertCommentInDatabase(username, commentText);
//         input.value = '';
//     }
// }

// // Function to submit a reply
// function submitReply(input, username) {
//     var replyText = input.value.trim();

//     if (replyText !== '') {
//         var commentElement = input.closest('.comment');
//         var commentID = commentElement.dataset.commentId;
//         var replyElement = createReplyElement(username, replyText);
//         var replyList = commentElement.querySelector('.replies');
//         replyList.appendChild(replyElement);
//         insertReplyInDatabase(commentID, username, replyText);
//         input.value = '';
//         delete input.dataset.replyingTo;
//         input.placeholder = 'Votre commentaire...';
//         var button = input.nextSibling;
//         button.textContent = 'Envoyer';
//         var postPrivate = commentElement.closest('.PostPrivate');
//         var commentsDiv = postPrivate.querySelector('.comments');
//         postPrivate.insertBefore(document.querySelector('.comment-input'), commentsDiv.nextSibling);
//     }
// }

// // Call the function to render the comments
// window.addEventListener('load', function() {
//     renderComments();
// });

// function renderComments() {
//     // Get the comments from the database
//     fetch('/getComments')
//     .then(response => {
//         if (!response.ok) {
//             throw new Error('Network response was not ok');
//         }
//         return response.json();
//     })
//     .then(comments => {
//         var postPrivates = document.querySelectorAll('.PostPrivate');
//         postPrivates.forEach(postPrivate => {
//             var commentsDiv = postPrivate.querySelector('.comments');
//             commentsDiv.innerHTML = '';
//             comments.forEach(comment => {
//                 if (comment.postID === postPrivate.dataset.postId) {
//                     var commentElement = createCommentElement(comment.username, comment.commentText);
//                     commentElement.dataset.commentId = comment._id;
//                     commentsDiv.appendChild(commentElement);
//                     comment.replies.forEach(reply => {
//                         var replyElement = createReplyElement(reply.username, reply.replyText);
//                         commentElement.querySelector('.replies').appendChild(replyElement);
//                     });
//                 }
//             });
//         });
//     })
//     .catch(error => {
//         console.error('Error getting comments:', error);
//     });
// }

// Function to show comment input
function showCommentInput(event) {
    event.preventDefault();

    var username = getUsername();
    var commentInputDiv = getOrCreateCommentInput();

    // Inserting the comment input after the comments
    var postPrivate = event.target.closest('.PostPrivate');
    var commentsDiv = postPrivate.querySelector('.comments');
    commentsDiv.appendChild(commentInputDiv);
}

// Function to show reply input
function showReplyInput(event, username) {
    event.preventDefault();

    var commentInputDiv = getOrCreateCommentInput();

    // Update the input for the reply
    var input = commentInputDiv.querySelector('input');
    input.dataset.replyingTo = 'true';
    input.placeholder = 'Votre réponse...';

    // Update the button for the reply
    var button = commentInputDiv.querySelector('button');
    button.textContent = 'Envoyer la réponse';

    // Insert the comment input after the comment
    var commentElement = event.target.closest('.comment');
    var repliesDiv = commentElement.querySelector('.replies');
    if (!repliesDiv) {
        repliesDiv = document.createElement('div');
        repliesDiv.classList.add('replies');
        commentElement.appendChild(repliesDiv);
    }

    repliesDiv.appendChild(commentInputDiv);
}

// Function to get the username
function getUsername() {
    return document.querySelector('.links a:last-child').textContent.trim();
}

// Function to get or create the comment input
function getOrCreateCommentInput() {
    var commentInputDiv = document.querySelector('.comment-input');
    if (!commentInputDiv) {
        commentInputDiv = createCommentInput();
    }
    return commentInputDiv;
}

// Function to create the comment input
function createCommentInput() {
    var commentInputDiv = document.createElement('div');
    commentInputDiv.classList.add('comment-input');

    var input = document.createElement('input');
    input.type = 'text';
    input.placeholder = 'Votre commentaire...';
    input.style.marginLeft = '25%';
    input.style.width = '50%';
    input.style.border = '1px solid #ccc';
    input.style.marginBottom = '20px';

    var button = document.createElement('button');
    button.textContent = 'Envoyer';
    button.onclick = function() {
        var input = commentInputDiv.querySelector('input');
        var username = getUsername();
        if (input.dataset.replyingTo) {
            submitReply(input, username);
        } else {
            submitComment(input, username);
        }
    };

    commentInputDiv.appendChild(input);
    commentInputDiv.appendChild(button);

    return commentInputDiv;
}

// Function to render a comment element
function createCommentElement(username, commentText) {
    var commentElement = document.createElement('div');
    commentElement.classList.add('comment');
    commentElement.style.width = '40%';
    commentElement.style.marginBottom = '20px';
    commentElement.style.marginLeft = '30%';

    var commentUsername = document.createElement('div');
    commentUsername.classList.add('comment-username');
    commentUsername.textContent = username;
    commentElement.appendChild(commentUsername);

    var commentTextElement = document.createElement('div');
    commentTextElement.textContent = commentText;
    commentElement.appendChild(commentTextElement);

    var replyLink = document.createElement('a');
    replyLink.href = '#';
    replyLink.textContent = 'Répondre';
    replyLink.onclick = function(event) {
        showReplyInput(event, username);
    };
    commentElement.appendChild(document.createElement('br'));
    commentElement.appendChild(replyLink);

    var repliesDiv = document.createElement('div');
    repliesDiv.classList.add('replies');
    commentElement.appendChild(repliesDiv);

    return commentElement;
}

// Function to create a reply element
function createReplyElement(username, replyText) {
    var replyElement = document.createElement('div');
    replyElement.classList.add('reply');
    replyElement.style.marginLeft = '35%';
    replyElement.style.width = '35%';

    var replyUsername = document.createElement('div');
    replyUsername.classList.add('reply-username');
    replyUsername.textContent = username;
    replyElement.appendChild(replyUsername);

    var replyTextElement = document.createElement('div');
    replyTextElement.textContent = replyText;
    replyElement.appendChild(replyTextElement);

    return replyElement;
}

// Function to submit a comment to the server
function submitComment(input, username) {
    var commentText = input.value.trim();

    if (commentText !== '') {
        var commentList = input.closest('.PostPrivate').querySelector('.comments');
        var commentElement = createCommentElement(username, commentText);
        commentList.appendChild(commentElement);
        insertCommentInDatabase(username, commentText); // Call API to insert into database
        input.value = '';
    }
}

// Function to submit a reply to the server
function submitReply(input, username) {
    var replyText = input.value.trim();

    if (replyText !== '') {
        var commentElement = input.closest('.comment');
        var replyElement = createReplyElement(username, replyText);
        var repliesDiv = commentElement.querySelector('.replies');
        repliesDiv.appendChild(replyElement);
        insertReplyInDatabase(commentElement.dataset.commentId, username, replyText); // Call API to insert into database
        input.value = '';
        delete input.dataset.replyingTo;
        input.placeholder = 'Votre commentaire...';
        var button = input.nextSibling;
        button.textContent = 'Envoyer';
        var postPrivate = commentElement.closest('.PostPrivate');
        var commentsDiv = postPrivate.querySelector('.comments');
        postPrivate.insertBefore(document.querySelector('.comment-input'), commentsDiv.nextSibling);
    }
}



// Call renderComments() when the page is loaded
window.addEventListener('load', function() {
    renderComments();
});

