// Fonction pour afficher la zone de commentaire au clic sur "Comment"
function showCommentInput(event) {
    event.preventDefault();

    // Créer la zone de commentaire (si elle n'existe pas)
    var commentInputDiv = document.querySelector('.comment-input');
    if (!commentInputDiv) {
        commentInputDiv = document.createElement('div');
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
            submitComment(input, getUsername()); // Utilisation de getUsername pour récupérer l'username
        };

        commentInputDiv.appendChild(input);
        commentInputDiv.appendChild(button);
    }

    // Insérer la zone de commentaire avant les commentaires
    var postPrivate = event.target.closest('.PostPrivate');
    var commentsDiv = postPrivate.querySelector('.comments');
    postPrivate.insertBefore(commentInputDiv, commentsDiv.nextSibling);
}

// Fonction pour soumettre un commentaire
function submitComment(input, username) {
    var commentText = input.value.trim();

    if (commentText !== '') {
        var commentList = input.closest('.PostPrivate').querySelector('.comments');

        // Créer un élément de commentaire avec l'username
        var commentElement = createCommentElement(username, commentText);

        // Ajouter le commentaire à la liste des commentaires
        commentList.appendChild(commentElement);

        // Réinitialiser le champ de commentaire
        input.value = '';
    }
}

// Fonction pour afficher la zone de réponse au clic sur "Répondre"
function showReplyInput(event, username) {
    event.preventDefault();

    // Marquer l'input actuel comme étant une réponse
    var commentInputDiv = document.querySelector('.comment-input');
    var input = commentInputDiv.querySelector('input');
    input.dataset.replyingTo = 'true';
    input.placeholder = 'Votre réponse...';

    // Mettre à jour le bouton pour correspondre au contexte de réponse
    var button = commentInputDiv.querySelector('button');
    button.textContent = 'Envoyer la réponse';

    // Insérer la zone de réponse après le commentaire auquel on répond
    var commentElement = event.target.closest('.comment');
    var repliesDiv = commentElement.querySelector('.replies');
    if (!repliesDiv) {
        repliesDiv = document.createElement('div');
        repliesDiv.classList.add('replies');
        commentElement.appendChild(repliesDiv);
    }

    repliesDiv.appendChild(commentInputDiv); // Déplacer l'input et le bouton à l'endroit approprié
}

// Fonction pour soumettre une réponse à un commentaire
function submitReply(input, username) {
    var replyText = input.value.trim();

    if (replyText !== '') {
        var commentElement = input.closest('.comment');
        var replyList = commentElement.querySelector('.replies');

        // Créer un élément de réponse avec l'username
        var replyElement = createReplyElement(username, replyText);

        // Ajouter la réponse à la liste des réponses
        replyList.appendChild(replyElement);

        // Réinitialiser le champ de réponse
        input.value = '';

        // Réinitialiser l'état de réponse de l'input
        delete input.dataset.replyingTo;
        input.placeholder = 'Votre commentaire...';

        // Mettre à jour le bouton après l'envoi de la réponse
        var button = input.nextSibling;
        button.textContent = 'Envoyer';

        // Remettre la zone de commentaire à son emplacement initial
        var postPrivate = commentElement.closest('.PostPrivate');
        var commentsDiv = postPrivate.querySelector('.comments');
        postPrivate.insertBefore(document.querySelector('.comment-input'), commentsDiv.nextSibling);
    }
}

// Fonction utilitaire pour créer un élément de commentaire
function createCommentElement(username, commentText) {
    var commentElement = document.createElement('div');
    commentElement.classList.add('comment');
    commentElement.style.width = '40%';
    commentElement.style.marginBottom = '20px';
    commentElement.style.marginLeft = '30%';
    commentElement.textContent = username + ': ' + commentText;

    // Créer un lien "Répondre"
    var replyLink = document.createElement('a');
    replyLink.href = '#';
    replyLink.textContent = 'Répondre';
    replyLink.onclick = function(event) {
        event.preventDefault();
        showReplyInput(event, username); // Appeler la fonction pour ajouter une réponse avec l'username
    };

    // Ajouter le lien "Répondre" en bas du commentaire
    commentElement.appendChild(document.createElement('br')); // Ajouter un saut de ligne
    commentElement.appendChild(replyLink);

    // Créer un div pour les réponses
    var repliesDiv = document.createElement('div');
    repliesDiv.classList.add('replies');
    commentElement.appendChild(repliesDiv);

    return commentElement;
}

// Fonction utilitaire pour créer un élément de réponse
function createReplyElement(username, replyText) {
    var replyElement = document.createElement('div');
    replyElement.classList.add('reply');
    replyElement.style.marginLeft = '35%';
    replyElement.style.width = '35%';
    replyElement.textContent = username + ': ' + replyText;

    return replyElement;
}

// Fonction utilitaire pour récupérer l'username
function getUsername() {
    var usernameElement = document.querySelector('.links a:last-child'); // Adapté selon votre structure HTML
    return usernameElement.textContent.trim();
}
