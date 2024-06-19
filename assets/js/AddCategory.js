document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("addCategory").addEventListener("click", addCategory);
});

function addCategory() {
    const categoryName = document.getElementById("cat").value.trim();
    const subjectName  = document.getElementById("sub").value.trim();

    if (categoryName === "" || subjectName === "") {
        alert("Il faut remplir les champs");
        return;
    }

    const donnees = new URLSearchParams();
    donnees.append("cat", categoryName);
    donnees.append("sub", subjectName);

    fetch("/category", {
        method: "POST",
        body: donnees
    }).then(response => {
        if (!response.ok) {
            throw new Error("Erreur HTTP : " + response.status);
        }
        return response.json();
    }).then(donnees => {
        console.log("Catégorie ajoutée avec succès", donnees);
        addCategoryToDOM(categoryName, subjectName);
    }).catch(error => {
        console.error("Erreur :", error);
        alert("Erreur lors de l'ajout de la catégorie");
    });

    document.getElementById("cat").value = "";
    document.getElementById("sub").value = "";
}

function addCategoryToDOM(categoryName, subjectName) {
    const newCategory = document.createElement("div");
    newCategory.classList.add("category-item");

    const categoryTitle = document.createElement("div");
    categoryTitle.classList.add("category-title");
    categoryTitle.innerHTML = '<a href="#">' + categoryName + '</a>';

    const categoryContainer = document.createElement("div");
    categoryContainer.classList.add("category-container");

    const subjectList = document.createElement("ul");
    subjectList.classList.add("subject-container");

    const subjectItem = document.createElement("li");
    const subjectLink = document.createElement("a");
    subjectLink.setAttribute("href", "#");
    subjectLink.textContent = subjectName;
    subjectItem.appendChild(subjectLink);
    subjectList.appendChild(subjectItem);

    const newInputSubject = document.createElement("input");
    newInputSubject.setAttribute("type", "text");
    newInputSubject.classList.add("sub-btn");

    const newButtonAdd = document.createElement("button");
    newButtonAdd.type = "button";
    newButtonAdd.textContent = "Ajouter";
    newButtonAdd.classList.add("add-btn");

    newButtonAdd.addEventListener("click", function() {
        const newSubjectName = newInputSubject.value.trim();
        if (newSubjectName === "") {
            alert("Il faut remplir les champs");
            return;
        }

        const newSubject = document.createElement("li");
        const newSubjectLink = document.createElement("a");
        newSubjectLink.setAttribute("href", "#");
        newSubjectLink.textContent = newSubjectName;
        newSubject.appendChild(newSubjectLink);
        subjectList.appendChild(newSubject);
        newInputSubject.value = "";
    });

    categoryContainer.appendChild(subjectList);
    categoryContainer.appendChild(newInputSubject);
    categoryContainer.appendChild(newButtonAdd);
    newCategory.appendChild(categoryTitle);
    newCategory.appendChild(categoryContainer);

    const categoryList = document.querySelector(".container");
    categoryList.appendChild(newCategory);
}
