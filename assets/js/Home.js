document.addEventListener('DOMContentLoaded', () => {
    // Récupérer le lien "Log in"
    const logInLink = document.querySelector('a[href="#2"]');
    
    // Ajouter un écouteur d'événement pour le clic sur le lien
    logInLink.addEventListener('click', (event) => {
        // Empêcher le comportement par défaut du lien (navigation)
        event.preventDefault();
        
        // Rediriger vers la page LoginPage.html
        window.location.href = 'templates/LoginPage.html';
    });
});
